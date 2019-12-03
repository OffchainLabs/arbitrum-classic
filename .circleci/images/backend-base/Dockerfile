FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl git golang librocksdb-dev sudo && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://get.parity.io -L | bash /dev/stdin -r stable && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    go get gotest.tools/gotestsum
