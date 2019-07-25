### --------------------------------------------------------------------
### Dockerfile
### arb-avm-cpp
### Note: `##DEV_` commands are run in dev-mode. They are not comments.
### --------------------------------------------------------------------

FROM alpine:3.9

# Alpine dependencies
RUN apk add --no-cache boost-dev cmake g++ gcc make musl-dev python3 python3-dev && \
    pip3 install conan && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"

COPY --chown=user conanfile.txt ./
RUN mkdir -p build && cd build && \
    conan profile new default --detect && \
    conan profile update settings.compiler.libcxx=libstdc++11 default && \
    conan remote add nonstd-lite https://api.bintray.com/conan/martinmoene/nonstd-lite && \
    conan install ..

# Copy source code
COPY --chown=user . ./

# Build cache
##DEV_COPY --from=arb-avm-cpp --chown=user /build build/

RUN cd build && conan install .. && \
    cmake .. -DCMAKE_BUILD_TYPE=Release && \
    cmake --build . && \
    cp ../cavm/cmachine.h lib/* ../cmachine

# Export library binary and header
FROM alpine:3.9
COPY --from=0 /home/user/go.mod /home/user/go.sum arb-avm-cpp/
COPY --from=0 /home/user/cmachine arb-avm-cpp/cmachine/
COPY --from=0 /home/user/build build/
