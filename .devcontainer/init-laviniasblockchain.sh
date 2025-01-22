#!/usr/bin/env bash
set -euo pipefail

echo "===================================================================================="
echo "BLS Signature Verification Precompile Test Suite"
echo "Testing the Go-implemented precompile contract for BLS signature verification"
echo "===================================================================================="

echo "Deploying blockchain with BLS verification precompile..."
avalanche blockchain deploy laviniasblockchain --local

RPC_URL="http://127.0.0.1:9652/ext/bc/22vm7RDESbGV44B4sYaX2yMKQSsaHcLoTCs65QRQBReKo3EQoo/rpc"
CONTRACT_ADDRESS="0x0100000000000000000000000000000000000003"

# Validation test cases with known BLS signatures
# Each test case provides: BLS Signature, Public Key, Message, and Expected Verification Result
TEST_CASES=(
  "0x86a3ab4c45cfe31cae34c1d06f212434ac71b1be6cfe046c80c162e057614a94a5bc9f1ded1a7029deb0ba4ca7c9b71411e293438691be79c2dbf19d1ca7c3eadb9c756246fc5de5b7b89511c7d7302ae051d9e03d7991138299b5ed6a570a98 0x8f95423f7142d00a48e1014a3de8d28907d420dc33b3052a6dee03a3f2941a393c2351e354704ca66a3fc29870282e15 0xabcdef false"
  "0x806047e7989551abc504b71a6d448e94af8cc6dbade0c27f734e38064dd1c88ef98b054d81ebd2573f14a0eb92967fa016f71d4de760245a100c731f126c728fe44b8109ebaa132150496245334a2c0e2ba1dee626b14864bc6ab4aa2d5e3d45 0xb580dd251f9faf4b8c66ab8ffe20846fa7fa8c239e31749663e0f311a61af43c6e9d19f1448cc09993c5b388a68e0f3e 0x4d7920696d706f7274616e74206d65737361676520746f207369676e true"
)

echo "Phase 1: Verifying precompile against known test vectors"
echo "------------------------------------------------------"

for i in "${!TEST_CASES[@]}"; do
  TEST="${TEST_CASES[$i]}"
  EXPECTED_RESULT=$(echo "$TEST" | awk '{print $4}')
  PARAMETERS=$(echo "$TEST" | awk '{print $1, $2, $3}')
  
  echo
  echo "Test Vector $((i + 1))"
  echo "Calling precompile with BLS components..."
  RESPONSE=$(cast call \
    --rpc-url "$RPC_URL" \
    "$CONTRACT_ADDRESS" \
    "verifySignature(bytes,bytes,bytes)(bool)" \
    $PARAMETERS)
  
  echo "Expected: $EXPECTED_RESULT"
  echo "Result: $RESPONSE"
  
  if [ "$RESPONSE" == "$EXPECTED_RESULT" ]; then
    echo "✓ Verification correct"
  else
    echo "✗ Verification failed"
  fi
done

echo
echo "Phase 2: Live BLS Signature Generation and Verification"
echo "----------------------------------------------------"
echo "Generating fresh BLS signature components..."

# Generate new BLS signature components
GO_OUTPUT=$(go run generate_bls.go)

# Extract the BLS components
LIVE_PUBLIC_KEY=$(echo "$GO_OUTPUT" | grep "^Public Key: " | head -n 1 | sed 's/^Public Key: //')
LIVE_MESSAGE=$(echo "$GO_OUTPUT" | grep "^Message in hex: " | head -n 1 | sed 's/^Message in hex: //')
LIVE_SIGNATURE=$(echo "$GO_OUTPUT" | grep "^Signature: " | head -n 1 | sed 's/^Signature: //')
ORIGINAL_MESSAGE=$(echo "$GO_OUTPUT" | grep "^Original Message: " | head -n 1 | sed 's/^Original Message: //')

echo "Message: \"$ORIGINAL_MESSAGE\""
echo
echo "Generated BLS Components:"
echo "------------------------"
echo "Public Key: $LIVE_PUBLIC_KEY"
echo "Message (hex): $LIVE_MESSAGE"
echo "Signature: $LIVE_SIGNATURE"

# Validate component extraction
if [[ -z "$LIVE_PUBLIC_KEY" || -z "$LIVE_MESSAGE" || -z "$LIVE_SIGNATURE" ]]; then
    echo "Error: BLS component generation failed"
    echo "Debug output:"
    echo "$GO_OUTPUT"
    exit 1
fi

echo
echo "Verifying fresh signature using precompile..."
LIVE_RESPONSE=$(cast call \
  --rpc-url "$RPC_URL" \
  "$CONTRACT_ADDRESS" \
  "verifySignature(bytes,bytes,bytes)(bool)" \
  "$LIVE_SIGNATURE" \
  "$LIVE_PUBLIC_KEY" \
  "$LIVE_MESSAGE")

echo "Verification Result: $LIVE_RESPONSE"
if [ "$LIVE_RESPONSE" == "true" ]; then
    echo "✓ Signature verified successfully"
else
    echo "✗ Signature verification failed"
fi