ARG AVALANCHEGO_VERSION
FROM avaplatform/avalanche-cli:latest as avalanche-cli
FROM ghcr.io/foundry-rs/foundry:latest as foundry
FROM avaplatform/avalanchego:$AVALANCHEGO_VERSION as avalanchego
FROM mcr.microsoft.com/devcontainers/base

COPY --from=avalanchego /avalanchego/build /go/src/github.com/ava-labs/avalanchego/build
COPY --from=avalanche-cli /avalanche /usr/local/bin/avalanche

COPY --from=foundry /usr/local/bin/forge /usr/local/bin/forge
COPY --from=foundry /usr/local/bin/cast /usr/local/bin/cast
COPY --from=foundry /usr/local/bin/anvil /usr/local/bin/anvil
COPY --from=foundry /usr/local/bin/chisel /usr/local/bin/chisel
