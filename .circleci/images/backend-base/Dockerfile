FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl git golang librocksdb-dev libgmp-dev sudo netcat-openbsd && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://get.parity.io -L | bash /dev/stdin -r stable && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    curl -O https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.9.9-01744997.tar.gz && \
    tar xvzf geth-linux-amd64-1.9.9-01744997.tar.gz && \
    mkdir -p bin && \
    mv geth-linux-amd64-1.9.9-01744997/geth bin && \
    rm -rf geth-linux-amd64-1.9.9-01744997* && \
    go get gotest.tools/gotestsum
