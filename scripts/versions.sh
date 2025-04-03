#!/usr/bin/env bash

# Ignore warnings about variables appearing unused since this file is not the consumer of the variables it defines.
# shellcheck disable=SC2034

set_module_version_var() {
  local env_var_name="$1"
  local module_path="$2"

  if [[ -z "$env_var_name" || -z "$module_path" ]]; then
    echo "Usage: set_module_version_var <ENV_VAR_NAME> <GO_MODULE_PATH>"
    return 1
  fi

  # Check if the environment variable is already set
  if [[ -z "${!env_var_name:-}" ]]; then
    # Get module details from go.mod
    local module_details
    module_details="$(go list -m "$module_path" 2>/dev/null)" || return 1

    # Extract version from module details
    local version
    version="$(echo "$module_details" | awk '{print $2}')"

    # If version is a pseudo-version (timestamp-hash format), extract the short hash
    if [[ "$version" =~ ^v.*[0-9]{14}-[0-9a-f]{12}$ ]]; then
      local module_hash
      module_hash="$(echo "$version" | grep -Eo '[0-9a-f]{12}$')"
      version="${module_hash::8}"
    fi

    # Don't export them as they're used in the context of other calls
    eval "$env_var_name=$version"
  fi
}

set_module_version_var "SUBNET_EVM_VERSION" "github.com/ava-labs/subnet-evm"
set_module_version_var "AVALANCHE_VERSION" "github.com/ava-labs/avalanchego"
