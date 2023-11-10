#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Avalanche root directory
ROOT_DIR_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Load the versions
source "$ROOT_DIR_PATH"/scripts/versions.sh

# Load the constants
source "$ROOT_DIR_PATH"/scripts/constants.sh

BUILD_IMAGE_ID=${BUILD_IMAGE_ID:-"$AVALANCHE_VERSION-$PRECOMPILE_COMMIT_ID"}

echo "Building Docker Image: $DOCKERHUB_REPO:$BUILD_IMAGE_ID based of $AVALANCHE_VERSION"
docker build -t "$DOCKERHUB_REPO:$BUILD_IMAGE_ID" "$ROOT_DIR_PATH" -f "$ROOT_DIR_PATH/Dockerfile" \
  --build-arg AVALANCHE_VERSION="$AVALANCHE_VERSION" \
  --build-arg PRECOMPILE_COMMIT="$PRECOMPILE_COMMIT" \
  --build-arg CURRENT_BRANCH="$CURRENT_BRANCH"
