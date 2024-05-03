#!/usr/bin/env bash

# Set the PATHS
GOPATH="$(go env GOPATH)"

# Static compilation
STATIC_LD_FLAGS=''
if [ "${STATIC_COMPILATION:-}" = 1 ]; then
    export CC=musl-gcc
    command -v $CC || (echo $CC must be available for static compilation && exit 1)
    STATIC_LD_FLAGS=' -extldflags "-static" -linkmode external '
fi

# Set the CGO flags to use the portable version of BLST
#
# We use "export" here instead of just setting a bash variable because we need
# to pass this flag to all child processes spawned by the shell.
export CGO_CFLAGS="-O2 -D__BLST_PORTABLE__"
