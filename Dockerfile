### --------------------------------------------------------------------
### Dockerfile
### arb-avm-cpp
### Note: `##DEV_` commands are run in dev-mode. They are not comments.
### --------------------------------------------------------------------

FROM alpine:3.9

# Alpine dependencies
RUN apk add --no-cache build-base boost-dev clang cmake ninja

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"

# Build cache
##DEV_COPY --from=arb-avm-cpp --chown=user /build build/

# Build
COPY --chown=user . ./
RUN mkdir -p build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -GNinja && cmake --build .

# Export library binary and header
FROM scratch
COPY --from=0 \
    /home/user/avm/include/avm/cmachine.h \
    /home/user/build build/
