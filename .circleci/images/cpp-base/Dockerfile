### --------------------------------------------------------------------
### Dockerfile
### cpp-base
### --------------------------------------------------------------------

FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl cmake git gcc g++ libboost-dev libboost-filesystem-dev \
    lcov make libgmp-dev libssl-dev libusb-dev sudo netcat-openbsd \
    autotools-dev dh-autoreconf pkg-config \
    google-perftools libgoogle-perftools-dev vim-tiny \
    libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev libzstd-dev && \
    curl -fsSL https://deb.nodesource.com/setup_14.x | bash - && \
    curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor | sudo tee /usr/share/keyrings/yarnkey.gpg >/dev/null && \
    echo "deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main" | sudo tee /etc/apt/sources.list.d/yarn.list && \
    apt-get update && sudo apt-get install -y nodejs yarn && \
    curl -L https://github.com/facebook/rocksdb/archive/refs/tags/v6.20.3.tar.gz --output rocksdb-6.20.3.tar.gz && \
    tar xf rocksdb-6.20.3.tar.gz && \
    cd rocksdb-6.20.3 && \
    PREFIX=/usr PORTABLE=1 make shared_lib install-shared && \
    strip --strip-unneeded /usr/lib/librocksdb.so.6.20.3 && \
    cd .. && \
    rm -rf rocksdb-6.20.3* && \
    rm -rf /usr/share/doc/* /usr/share/fonts/* && \
    useradd -ms /bin/bash user

USER user
WORKDIR /home/user/
ENV PATH="/home/user/bin:/home/user/.local/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash
