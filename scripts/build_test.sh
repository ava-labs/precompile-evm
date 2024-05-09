#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail
# TODO(marun) Ensure the working directory is the repository root or a non-canonical set of tests may be executed

# Root directory
ROOT_DIR_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

# Load the versions
source "$ROOT_DIR_PATH"/scripts/versions.sh

# Load the constants
source "$ROOT_DIR_PATH"/scripts/constants.sh

# We pass in the arguments to this script directly to enable easily passing parameters such as enabling race detection,
# parallelism, and test coverage.
# DO NOT RUN "tests/precompile" or "tests/load" since it's run by ginkgo
go test -shuffle=on -race -timeout="${TIMEOUT:-600s}" -coverprofile=coverage.out -covermode=atomic "$@" $(go list ./... | grep -v tests/precompile | grep -v tests/load)
