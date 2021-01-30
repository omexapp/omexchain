# Simple usage with a mounted data directory:
# > docker build -t omexchain .
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.omexchaind:/root/.omexchaind -v ~/.omexchaincli:/root/.omexchaincli omexchain omexchaind init mynode
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.omexchaind:/root/.omexchaind -v ~/.omexchaincli:/root/.omexchaincli omexchain omexchaind start
FROM golang:alpine AS build-env

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev

# Set working directory for the build
WORKDIR /go/src/github.com/omexapp/omexchain

# Add source files
COPY . .

# Build OMExChain
RUN GOPROXY=http://goproxy.cn make install

# Final image
FROM alpine:edge

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/omexchaind /usr/bin/omexchaind
COPY --from=build-env /go/bin/omexchaincli /usr/bin/omexchaincli

# Run omexchaind by default, omit entrypoint to ease using container with omexchaincli
CMD ["omexchaind"]
