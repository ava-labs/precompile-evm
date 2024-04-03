FROM avaplatform/avalanchego:v1.11.3 as avalanchego
FROM mcr.microsoft.com/devcontainers/base

COPY --from=avalanchego /avalanchego/build /go/src/github.com/ava-labs/avalanchego/build

RUN curl -sSfL https://raw.githubusercontent.com/ava-labs/avalanche-network-runner/main/scripts/install.sh | sh -s
