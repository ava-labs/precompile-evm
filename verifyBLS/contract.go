// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package verifyBLS

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/vmerrs"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"

	bls "github.com/ava-labs/avalanchego/utils/crypto/bls"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	AggregatePublicKeysGasCost uint64 = 1 /* SET A GAS COST HERE */
	AggregateSignaturesGasCost uint64 = 1 /* SET A GAS COST HERE */
	VerifySignatureBLSGasCost  uint64 = 1 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
	_ = vmerrs.ErrOutOfGas
	_ = common.Big0
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// VerifyBLSRawABI contains the raw ABI of VerifyBLS contract.
	//go:embed contract.abi
	VerifyBLSRawABI string

	VerifyBLSABI = contract.ParseABI(VerifyBLSRawABI)

	VerifyBLSPrecompile = createVerifyBLSPrecompile()
)

type VerifySignatureBLSInput struct {
	Message   string
	Signature []byte
	PublicKey []byte
}

// UnpackAggregatePublicKeysInput attempts to unpack [input] into the [][]byte type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackAggregatePublicKeysInput(input []byte) ([][]byte, error) {
	res, err := VerifyBLSABI.UnpackInput("aggregatePublicKeys", input, false)
	if err != nil {
		return nil, err
	}
	unpacked := *abi.ConvertType(res[0], new([][]byte)).(*[][]byte)
	return unpacked, nil
}

// PackAggregatePublicKeys packs [publicKeys] of type [][]byte into the appropriate arguments for aggregatePublicKeys.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackAggregatePublicKeys(publicKeys [][]byte) ([]byte, error) {
	return VerifyBLSABI.Pack("aggregatePublicKeys", publicKeys)
}

// PackAggregatePublicKeysOutput attempts to pack given publicKey of type []byte
// to conform the ABI outputs.
func PackAggregatePublicKeysOutput(publicKey []byte) ([]byte, error) {
	return VerifyBLSABI.PackOutput("aggregatePublicKeys", publicKey)
}

// UnpackAggregatePublicKeysOutput attempts to unpack given [output] into the []byte type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackAggregatePublicKeysOutput(output []byte) ([]byte, error) {
	res, err := VerifyBLSABI.Unpack("aggregatePublicKeys", output)
	if err != nil {
		return []byte{}, err
	}
	unpacked := *abi.ConvertType(res[0], new([]byte)).(*[]byte)
	return unpacked, nil
}

func aggregatePublicKeys(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, AggregatePublicKeysGasCost); err != nil {
		return nil, 0, err
	}
	
	inputStruct, err := UnpackAggregatePublicKeysInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	var pubKeys []*bls.PublicKey
  for _, pkBytes := range inputStruct {
      pk, err := bls.PublicKeyFromCompressedBytes(pkBytes)
      if err != nil {
          return nil, remainingGas, fmt.Errorf("failed to parse public key: %w", err)
      }
      pubKeys = append(pubKeys, pk)
  }
	
	aggPk, err := bls.AggregatePublicKeys(pubKeys)
    if err != nil {
        return nil, remainingGas, fmt.Errorf("failed to aggregate public keys: %w", err)
    }

  // Convert the aggregated public key to compressed bytes.
  output := bls.PublicKeyToCompressedBytes(aggPk)

	packedOutput, err := PackAggregatePublicKeysOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackAggregateSignaturesInput attempts to unpack [input] into the [][]byte type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackAggregateSignaturesInput(input []byte) ([][]byte, error) {
	res, err := VerifyBLSABI.UnpackInput("aggregateSignatures", input, false)
	if err != nil {
		return nil, err
	}
	unpacked := *abi.ConvertType(res[0], new([][]byte)).(*[][]byte)
	return unpacked, nil
}

// PackAggregateSignatures packs [signatures] of type [][]byte into the appropriate arguments for aggregateSignatures.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackAggregateSignatures(signatures [][]byte) ([]byte, error) {
	return VerifyBLSABI.Pack("aggregateSignatures", signatures)
}

// PackAggregateSignaturesOutput attempts to pack given signature of type []byte
// to conform the ABI outputs.
func PackAggregateSignaturesOutput(signature []byte) ([]byte, error) {
	return VerifyBLSABI.PackOutput("aggregateSignatures", signature)
}

// UnpackAggregateSignaturesOutput attempts to unpack given [output] into the []byte type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackAggregateSignaturesOutput(output []byte) ([]byte, error) {
	res, err := VerifyBLSABI.Unpack("aggregateSignatures", output)
	if err != nil {
		return []byte{}, err
	}
	unpacked := *abi.ConvertType(res[0], new([]byte)).(*[]byte)
	return unpacked, nil
}

func aggregateSignatures(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, AggregateSignaturesGasCost); err != nil {
		return nil, 0, err
	}
	

	inputStruct, err := UnpackAggregateSignaturesInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	var sigs []*bls.Signature
  for _, sigBytes := range inputStruct {
      sig, err := bls.SignatureFromBytes(sigBytes)
      if err != nil {
          return nil, remainingGas, fmt.Errorf("failed to parse signature: %w", err)
      }
      sigs = append(sigs, sig)
  }

	aggSig, err := bls.AggregateSignatures(sigs)
  if err != nil {
      return nil, remainingGas, fmt.Errorf("failed to aggregate signatures: %w", err)
  }

  // Convert the aggregated signature to its compressed byte format.
  output := bls.SignatureToBytes(aggSig)

  // Pack the output according to your ABI.
  packedOutput, err := PackAggregateSignaturesOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackVerifySignatureBLSInput attempts to unpack [input] as VerifySignatureBLSInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifySignatureBLSInput(input []byte) (VerifySignatureBLSInput, error) {
	inputStruct := VerifySignatureBLSInput{}
	err := VerifyBLSABI.UnpackInputIntoInterface(&inputStruct, "verifySignatureBLS", input, false)

	return inputStruct, err
}

// PackVerifySignatureBLS packs [inputStruct] of type VerifySignatureBLSInput into the appropriate arguments for verifySignatureBLS.
func PackVerifySignatureBLS(inputStruct VerifySignatureBLSInput) ([]byte, error) {
	return VerifyBLSABI.Pack("verifySignatureBLS", inputStruct.Message, inputStruct.Signature, inputStruct.PublicKey)
}

// PackVerifySignatureBLSOutput attempts to pack given result of type bool
// to conform the ABI outputs.
func PackVerifySignatureBLSOutput(result bool) ([]byte, error) {
	return VerifyBLSABI.PackOutput("verifySignatureBLS", result)
}

// UnpackVerifySignatureBLSOutput attempts to unpack given [output] into the bool type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackVerifySignatureBLSOutput(output []byte) (bool, error) {
	res, err := VerifyBLSABI.Unpack("verifySignatureBLS", output)
	if err != nil {
		return false, err
	}
	unpacked := *abi.ConvertType(res[0], new(bool)).(*bool)
	return unpacked, nil
}

func verifySignatureBLS(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, VerifySignatureBLSGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the VerifySignatureBLSInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackVerifySignatureBLSInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// Convert the message string to bytes.
	messageBytes := []byte(inputStruct.Message)

	// Parse the signature from bytes.
	sig, err := bls.SignatureFromBytes(inputStruct.Signature)
	if err != nil {
		return nil, remainingGas, fmt.Errorf("failed to parse signature: %w", err)
	}

	// Parse the public key from bytes.
	pubKey, err := bls.PublicKeyFromCompressedBytes(inputStruct.PublicKey)
	if err != nil {
		return nil, remainingGas, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Use the BLS package to verify the signature.
	// This returns a boolean indicating whether the signature is valid.
	verified := bls.Verify(pubKey, sig, messageBytes)

	// Pack the boolean output.
	packedOutput, err := PackVerifySignatureBLSOutput(verified)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createVerifyBLSPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createVerifyBLSPrecompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"aggregatePublicKeys": aggregatePublicKeys,
		"aggregateSignatures": aggregateSignatures,
		"verifySignatureBLS":  verifySignatureBLS,
	}

	for name, function := range abiFunctionMap {
		method, ok := VerifyBLSABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}
