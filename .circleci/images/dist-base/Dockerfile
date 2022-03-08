### --------------------------------------------------------------------
### Dockerfile
### dist-base
### --------------------------------------------------------------------

FROM offchainlabs/backend-base:0.6.1

USER root
WORKDIR /
RUN export CARGO_HOME=/usr/local/cargo && \
    curl https://sh.rustup.rs -sSf | bash -s - -y  && \
    /usr/local/cargo/bin/cargo install b3sum

USER user

FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y curl \
    procps jq rsync \
    node-ws vim-tiny libatomic1 python3 \
    libgmp10 libssl1.1 \
    libgoogle-perftools4 \
    libgflags2.2 libsnappy1v5 libzstd1 dnsutils && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /usr/share/doc/* && \
    useradd -ms /bin/bash user

COPY --from=0 /usr/lib/librocksdb.so.6.20.3 /usr/lib/librocksdb.so.6.20.3
COPY --from=0 /usr/local/cargo/bin/b3sum /usr/local/cargo/bin/b3sum

RUN ln -s librocksdb.so.6.20.3 /usr/lib/librocksdb.so.6.20 && \
    ln -s librocksdb.so.6.20.3 /usr/lib/librocksdb.so.6 && \
    ln -s librocksdb.so.6.20.3 /usr/lib/librocksdb.so

USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:/usr/local/cargo/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}" \
    GOROOT=/usr/local/go
