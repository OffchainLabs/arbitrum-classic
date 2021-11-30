### --------------------------------------------------------------------
### Dockerfile
### integration-base
### --------------------------------------------------------------------

FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl cmake git gcc g++ libboost-dev libboost-filesystem-dev \
    lcov make libgmp-dev libssl-dev libusb-dev sudo netcat-openbsd nodejs \
    autotools-dev dh-autoreconf pkg-config \
    google-perftools libgoogle-perftools-dev \
    libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev libzstd-dev && \
    rm -rf /var/lib/apt/lists/* && \
    curl -L https://github.com/facebook/rocksdb/archive/refs/tags/v6.20.3.tar.gz --output rocksdb-6.20.3.tar.gz && \
    tar xf rocksdb-6.20.3.tar.gz && \
    cd rocksdb-6.20.3 && \
    PREFIX=/usr make shared_lib install-shared && \
    curl -fsSL https://deb.nodesource.com/setup_14.x | bash - && \
    curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor | sudo tee /usr/share/keyrings/yarnkey.gpg >/dev/null && \
    echo "deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main" | sudo tee /etc/apt/sources.list.d/yarn.list && \
    apt-get update && sudo apt-get install yarn && \
    curl https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz --output go1.15.2.linux-amd64.tar.gz && \
    tar -xf go1.15.2.linux-amd64.tar.gz && \
    mv go /usr/local && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}"
RUN export GOROOT=/usr/local/go && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    go get gotest.tools/gotestsum
