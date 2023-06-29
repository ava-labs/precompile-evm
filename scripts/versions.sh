#!/usr/bin/env bash

# Set up the versions to be used - populate ENV variables only if they are not already populated
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-'v0.5.2'}
# Don't export them as they're used in the context of other calls
AVALANCHEGO_VERSION=${AVALANCHE_VERSION:-'v1.10.2'}
GINKGO_VERSION=${GINKGO_VERSION:-'v2.2.0'}
