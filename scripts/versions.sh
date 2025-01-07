#!/usr/bin/env bash

# Set up the versions to be used - populate ENV variables only if they are not already populated
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-'v0.7.0'}
# Don't export them as they're used in the context of other calls
AVALANCHE_VERSION=${AVALANCHE_VERSION:-'v1.12.0-config-pebble-sync'}
GINKGO_VERSION=${GINKGO_VERSION:-'v2.2.0'}
