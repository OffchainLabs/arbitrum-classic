### --------------------------------------------------------------------
### Dockerfile
### dist-base
### --------------------------------------------------------------------

FROM debian:bullseye-slim as debian-base

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl cmake git gcc g++ libboost-dev libboost-filesystem-dev \
    lcov make libgmp-dev libssl-dev libusb-dev sudo netcat-openbsd nodejs \
    autotools-dev dh-autoreconf pkg-config libjemalloc-dev \
    libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev libzstd-dev && \
    curl -L https://github.com/facebook/rocksdb/archive/refs/tags/v6.20.3.tar.gz --output rocksdb-6.20.3.tar.gz && \
    tar xf rocksdb-6.20.3.tar.gz && \
    cd rocksdb-6.20.3 && \
    PREFIX=/usr make shared_lib install-shared

FROM debian:bullseye-slim as debian

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl \
    node-ws vim-tiny libatomic1 python3 \
    libgmp10 libssl1.1 \
    libjemalloc2 \
    libgflags2.2 libsnappy1v5 libzstd1  && \
    rm -rf /var/lib/apt/lists/* /usr/share/doc/* && \
    useradd -ms /bin/bash user

COPY --from=0 /usr/lib/librocksdb.so.6.20.3 /usr/lib/librocksdb.so.6.20.3

RUN ln -s librocksdb.so.6.20.3 /usr/lib/librocksdb.so.6.20 && \
    ln -s librocksdb.so.6.20.3 /usr/lib/librocksdb.so.6 && \
    ln -s librocksdb.so.6.20.3 /usr/lib/librocksdb.so

USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}"
RUN export GOROOT=/usr/local/go && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash
