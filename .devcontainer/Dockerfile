# Use a base image that includes both Go and Node.js, such as the official Go image
FROM golang:latest

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin
ENV AVALANCHEGO_PLUGIN_PATH=$GOPATH/src/github.com/ava-labs/avalanchego/build/plugins

# Download AvalancheGo Binary
WORKDIR /
COPY ./scripts/versions.sh /
COPY ./.devcontainer/install_avalanchego.sh /

ENV AVALANCHEGO_EXEC_PATH=/avalanchego

RUN bash -c "source ./install_avalanchego.sh"

# Install Avalanche Network Runner
RUN curl -sSfL https://raw.githubusercontent.com/ava-labs/avalanche-network-runner/main/scripts/install.sh | sh -s

ENV PATH ~/bin:$PATH
