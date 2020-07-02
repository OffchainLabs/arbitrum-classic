### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:edge as arb-avm-cpp
# Alpine dependencies
RUN apk update && apk add --no-cache autoconf automake boost-dev cmake file g++ \
    git gmp-dev inotify-tools libtool make musl-dev openssl-dev && \
    apk add py-pip --no-cache && \
    apk add rocksdb-dev --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"
# Copy source code
COPY --chown=user arb-avm-cpp/ ./
# Copy build cache
COPY --from=arb-validator --chown=user /cpp-build build/
# Build arb-avm-cpp
RUN mkdir -p build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release && \
    cmake --build . -j $(nproc) && \
    cp lib/lib* ../cmachine


FROM alpine:edge as arb-validator-builder
# Alpine dependencies
RUN apk add --no-cache build-base git go \
    libc-dev linux-headers && \
    apk add gmp-dev rocksdb-dev --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ && \
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
RUN go mod download
# Copy source code
COPY --from=arb-avm-cpp /home/user/go.mod /home/user/go.sum /home/user/arb-avm-cpp/
COPY --from=arb-avm-cpp /home/user/cavm/*.h /home/user/arb-avm-cpp/cavm/
COPY --from=arb-avm-cpp /home/user/cmachine /home/user/arb-avm-cpp/cmachine/
COPY --chown=user arb-util/ /home/user/arb-util/
COPY --chown=user arb-validator/ /home/user/arb-validator/
COPY --chown=user arb-validator-core/ /home/user/arb-validator-core/
# Copy build cache
COPY --from=arb-validator --chown=user /build /home/user/.cache/go-build
# Build arb-validator
RUN go install -v ./cmd/arb-validator


FROM alpine:3.9 as arb-validator
# Export binary
RUN apk add --no-cache libstdc++=8.3.0-r0 libgcc=8.3.0-r0 && \
    apk add rocksdb --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/"
COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin
COPY --chown=user arb-validator/server.crt arb-validator/server.key ./

# Build cache
COPY --chown=user --from=arb-validator-builder /home/user/.cache/go-build /build
COPY --from=arb-avm-cpp /home/user/build /cpp-build

ENTRYPOINT ["/home/user/go/bin/arb-validator"]
EXPOSE 1235 1236
