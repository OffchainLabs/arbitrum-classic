### --------------------------------------------------------------------
### Dockerfile
### cpp-base
### --------------------------------------------------------------------

FROM debian:bookworm-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl cmake git gcc g++ libboost-dev libboost-filesystem-dev \
    lcov make libgmp-dev libssl-dev libusb-dev sudo netcat-openbsd \
    autotools-dev dh-autoreconf pkg-config xz-utils gnupg \
    google-perftools libgoogle-perftools-dev vim-tiny \
    libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev libzstd-dev && \
    NODE_VERSION=14.21.3 && \
    case "$(dpkg --print-architecture)" in \
      amd64) NODE_ARCH=x64 ;; \
      arm64) NODE_ARCH=arm64 ;; \
      *) echo "unsupported architecture: $(dpkg --print-architecture)" >&2; exit 1 ;; \
    esac && \
    curl -fsSL "https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-${NODE_ARCH}.tar.xz" -o /tmp/node.tar.xz && \
    tar -xJf /tmp/node.tar.xz -C /usr/local --strip-components=1 --no-same-owner && \
    rm /tmp/node.tar.xz && \
    curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor > /usr/share/keyrings/yarnkey.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main" > /etc/apt/sources.list.d/yarn.list && \
    apt-get update && apt-get install -y yarn && \
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
