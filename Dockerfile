### --------------------------------------------------------------------
### Dockerfile
### arb-avm-cpp
### Note: `##DEV_` commands are run in dev-mode. They are not comments.
### --------------------------------------------------------------------

FROM alpine:3.9

# Alpine dependencies
RUN apk add --no-cache musl-dev boost-dev gcc g++ cmake python3 python3-dev make

RUN pip3 install conan

# Non-root user
RUN addgroup -g 1000 -S user && \
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

RUN mkdir -p build && cd build && \
	conan install .. && \
	cmake .. -DCMAKE_BUILD_TYPE=Release && \
	cmake --build .

RUN cp cavm/cmachine.h build/lib/* cmachine

# Export library binary and header
FROM alpine:3.9
COPY --from=0 /home/user/go.mod arb-avm-cpp/
COPY --from=0 /home/user/go.sum arb-avm-cpp/
COPY --from=0 /home/user/cmachine arb-avm-cpp/cmachine/
COPY --from=0 /home/user/build build/
CMD /bin/true
