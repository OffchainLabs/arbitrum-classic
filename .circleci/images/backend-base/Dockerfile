FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl git librocksdb-dev libgmp-dev libssl-dev sudo netcat-openbsd gcc g++ libjemalloc-dev && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz --output go1.15.2.linux-amd64.tar.gz && \
    tar -xf go1.15.2.linux-amd64.tar.gz && \
    mv go /usr/local && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/usr/local/go/bin:/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN export GOROOT=/usr/local/go && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    go get gotest.tools/gotestsum
