FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl git librocksdb-dev libgmp-dev sudo netcat-openbsd nodejs npm gcc g++ && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz --output go1.13.3.linux-amd64.tar.gz && \
    tar -xf go1.13.3.linux-amd64.tar.gz && \
    mv go /usr/local && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/usr/local/go/bin:/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g yarn && \
    export GOROOT=/usr/local/go && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    go get gotest.tools/gotestsum
