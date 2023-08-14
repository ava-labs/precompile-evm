# Use a base image that includes both Go and Node.js, such as the official Go image
FROM golang:latest

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin
ENV AVALANCHEGO_PLUGIN_PATH=$GOPATH/src/github.com/ava-labs/avalanchego/build/plugins

# Download AvalancheGo Binary
WORKDIR /
COPY versions.sh /
RUN bash -c 'source /versions.sh && curl -LJ -o avalanchego.tar.gz "https://github.com/ava-labs/avalanchego/releases/download/{$AVALANCHEGO_VERSION}/avalanchego-linux-amd64-{$AVALANCHEGO_VERSION}.tar.gz"'
RUN tar -xzf avalanchego.tar.gz --wildcards '*/avalanchego' --strip-components=1 -C /avalanchego
RUN rm avalanchego.tar.gz
ENV AVALANCHEGO_EXEC_PATH=/avalanchego

# Install Node.js and npm using the official Node.js image
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && apt-get install -y nodejs

# Install Avalanche Network Runner
RUN curl -sSfL https://raw.githubusercontent.com/ava-labs/avalanche-network-runner/main/scripts/install.sh | sh -s

ENV PATH ~/bin:$PATH
