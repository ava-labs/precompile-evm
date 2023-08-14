# Use a base image that includes both Go and Node.js, such as the official Go image
FROM golang:latest


ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin
ENV AVALANCHEGO_PLUGIN_PATH=$GOPATH/src/github.com/ava-labs/avalanchego/build/plugins

# Install Node.js and npm using the official Node.js image
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && apt-get install -y nodejs

WORKDIR /

# # Build the avalanchego project using the sh script
RUN curl -LJ -o avalanchego.tar.gz https://github.com/ava-labs/avalanchego/releases/download/v1.10.5/avalanchego-linux-amd64-v1.10.5.tar.gz 
RUN mkdir avalanchego
RUN tar -xzf avalanchego.tar.gz --wildcards '*/avalanchego' --strip-components=1 -C /avalanchego
RUN rm avalanchego.tar.gz
ENV AVALANCHEGO_EXEC_PATH=/avalanchego

RUN curl -sSfL https://raw.githubusercontent.com/ava-labs/avalanche-network-runner/main/scripts/install.sh | sh -s

ENV PATH ~/bin:$PATH
