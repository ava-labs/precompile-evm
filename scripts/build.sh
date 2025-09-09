#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Root directory
ROOT_DIR_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

# Load the versions
source "$ROOT_DIR_PATH"/scripts/versions.sh

# Load the constants
source "$ROOT_DIR_PATH"/scripts/constants.sh

if [[ $# -eq 1 ]]; then
    BINARY_PATH=$1
elif [[ $# -eq 0 ]]; then
    BINARY_PATH="$DEFAULT_PLUGIN_DIR/$DEFAULT_VM_ID"
else
    echo "Invalid arguments to build precompile-evm. Requires zero (default location) or one argument to specify binary location."
    exit 1
fi

# Build Subnet EVM, which is run as a subprocess
echo "Building Precompile-EVM with Subnet-EVM version: $SUBNET_EVM_VERSION at $BINARY_PATH"
go build -ldflags "-X github.com/ava-labs/subnet-evm/plugin/evm.Version=$SUBNET_EVM_VERSION $STATIC_LD_FLAGS" -o "$BINARY_PATH" "plugin/"*.go
