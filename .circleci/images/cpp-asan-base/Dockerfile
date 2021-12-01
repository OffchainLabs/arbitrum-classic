### --------------------------------------------------------------------
### Dockerfile
### cpp-asan-base
### --------------------------------------------------------------------

FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl cmake git gcc g++ libboost-dev libboost-filesystem-dev \
    lcov make libgmp-dev libssl-dev libusb-dev sudo netcat-openbsd nodejs \
    autotools-dev dh-autoreconf pkg-config \
    google-perftools libgoogle-perftools-dev \
    libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev && \
    rm -rf /var/lib/apt/lists/* && \
    curl -L https://github.com/facebook/zstd/archive/refs/tags/v1.5.0.tar.gz --output zstd-1.5.0.tar.gz && \
    tar xf zstd-1.5.0.tar.gz && \
    cd zstd-1.5.0 && \
    CFLAGS=-fsanitize=address LDFLAGS=-fsanitize=address make install && \
    cd .. && \
    curl -L https://github.com/facebook/rocksdb/archive/refs/tags/v6.20.3.tar.gz --output rocksdb-6.20.3.tar.gz && \
    tar xf rocksdb-6.20.3.tar.gz && \
    cd rocksdb-6.20.3 && \
    COMPILE_WITH_ASAN=1 PREFIX=/usr make shared_lib install-shared && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}"
RUN export GOROOT=/usr/local/go && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash
