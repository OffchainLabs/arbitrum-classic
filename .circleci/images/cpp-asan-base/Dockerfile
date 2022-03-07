### --------------------------------------------------------------------
### Dockerfile
### cpp-asan-base
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.6.0

USER root
WORKDIR /
RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get remove -y libzstd-dev && \
    rm /usr/lib/librocksdb.* && \
    curl -L https://github.com/facebook/zstd/archive/refs/tags/v1.5.0.tar.gz --output zstd-1.5.0.tar.gz && \
    tar xf zstd-1.5.0.tar.gz && \
    cd zstd-1.5.0 && \
    CFLAGS=-fsanitize=address LDFLAGS=-fsanitize=address make install && \
    cd .. && \
    rm -rf zstd-1.5.0* && \
    curl -L https://github.com/facebook/rocksdb/archive/refs/tags/v6.20.3.tar.gz --output rocksdb-6.20.3.tar.gz && \
    tar xf rocksdb-6.20.3.tar.gz && \
    cd rocksdb-6.20.3 && \
    COMPILE_WITH_ASAN=1 PREFIX=/usr PORTABLE=1 make shared_lib install-shared && \
    strip --strip-unneeded /usr/lib/librocksdb.so.6.20.3 && \
    cd .. && \
    rm -rf rocksdb-6.20.3*

USER user
WORKDIR /home/user/
ENV PATH="/home/user/bin:/home/user/.local/bin:${PATH}"
