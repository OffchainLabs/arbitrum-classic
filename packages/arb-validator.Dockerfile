### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:edge as arb-avm-cpp
# Alpine dependencies
RUN apk update && apk add --no-cache autoconf automake boost-dev cmake file g++ libstdc++=9.3.0-r4 libgcc=9.3.0-r4 \
    git gmp-dev inotify-tools libtool make musl-dev openssl-dev && \
    apk add py-pip --no-cache && \
    apk add rocksdb-dev --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/testing
    # addgroup -g 1000 -S user && \
    # adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
# USER user
# WORKDIR "/home/user/"
# Copy source code
COPY arb-avm-cpp/ ./
# Copy build cache
COPY --from=arb-validator /cpp-build build/
# Build arb-avm-cpp
RUN echo "nameserver 8.8.8.8" > /etc/resolv.conf && \
    mkdir -p build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc) && \
    cd ../ && \
    ./scripts/install-cmachine-build


FROM alpine:edge as arb-validator-builder
# Alpine dependencies
RUN apk add --no-cache build-base git go libstdc++=9.3.0-r4 libgcc=9.3.0-r4 \
    libc-dev linux-headers && \
    apk add gmp-dev rocksdb-dev --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/testing && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user && \
    mkdir /home/user/arb-validator && \
    chown user:user /home/user/arb-validator
USER user
WORKDIR "/home/user/arb-validator"
# Build dependencies
COPY --chown=user arb-avm-cpp/go.* /home/user/arb-avm-cpp/
COPY --chown=user arb-util/go.* /home/user/arb-util/
COPY --chown=user arb-validator/go.* /home/user/arb-validator/
COPY --chown=user arb-validator-core/go.* /home/user/arb-validator-core/
COPY --chown=user arb-provider-go/go.* /home/user/arb-provider-go/
COPY --chown=user arb-checkpointer/go.* /home/user/arb-checkpointer/
COPY --chown=user arb-evm/go.* /home/user/arb-evm/
COPY --chown=user arb-tx-aggregator/go.* /home/user/arb-tx-aggregator/
RUN go mod download
# Copy source code
COPY --chown=user --from=arb-avm-cpp /go.mod /go.sum /home/user/arb-avm-cpp/
COPY --chown=user --from=arb-avm-cpp /cavm/*.h /home/user/arb-avm-cpp/cavm/
COPY --chown=user --from=arb-avm-cpp /cmachine /home/user/arb-avm-cpp/cmachine/

COPY --chown=user arb-util/ /home/user/arb-util/
COPY --chown=user arb-avm-cpp/ /home/user/arb-avm-cpp/
COPY --chown=user arb-validator/ /home/user/arb-validator/
COPY --chown=user arb-validator-core/ /home/user/arb-validator-core/
COPY --chown=user arb-provider-go/ /home/user/arb-provider-go/
COPY --chown=user arb-checkpointer/ /home/user/arb-checkpointer/
COPY --chown=user arb-evm/ /home/user/arb-evm/
COPY --chown=user arb-tx-aggregator/ /home/user/arb-tx-aggregator/
# Copy build cache
COPY --from=arb-validator --chown=user /build /home/user/.cache/go-build
# Build arb-validator
RUN go install -v ./cmd/arb-validator
WORKDIR "/home/user/arb-tx-aggregator"
RUN go install -v ./cmd/arb-tx-aggregator

FROM alpine:edge as arb-validator
# Export binary
RUN apk add --no-cache libstdc++=9.3.0-r4 libgcc=9.3.0-r4 && \
    apk add rocksdb --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/testing && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/"
COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin

# Build cache
COPY --chown=user --from=arb-validator-builder /home/user/.cache/go-build /build
COPY --chown=user --from=arb-avm-cpp build /cpp-build

ENTRYPOINT ["/home/user/go/bin/arb-validator"]
EXPOSE 1235 1236
