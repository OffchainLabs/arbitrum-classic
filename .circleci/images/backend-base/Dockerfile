FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl git librocksdb-dev libgmp-dev sudo netcat-openbsd gcc g++ && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz --output go1.13.3.linux-amd64.tar.gz && \
    tar -xf go1.13.3.linux-amd64.tar.gz && \
    mv go /usr/local && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/usr/local/go/bin:/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN export GOROOT=/usr/local/go && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    curl -O https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.9.9-01744997.tar.gz && \
    tar xvzf geth-linux-amd64-1.9.9-01744997.tar.gz && \
    mkdir -p bin && \
    mv geth-linux-amd64-1.9.9-01744997/geth bin && \
    rm -rf geth-linux-amd64-1.9.9-01744997* && \
    go get gotest.tools/gotestsum
