FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

WORKDIR /omexchain
# Add source files
COPY . .

# build OMExChain
RUN make build-linux

# Final image
FROM golang:1.14 as final

WORKDIR /omexchaind
# Copy over binaries from the build-env
COPY --from=build-env /omexchain/build/omexchaind /usr/bin/omexchaind
COPY --from=build-env /omexchain/build/omexchaincli /usr/bin/omexchaincli
COPY --from=build-env /omexchain/networks/local/node/wrapper.sh /usr/bin/wrapper.sh

EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

