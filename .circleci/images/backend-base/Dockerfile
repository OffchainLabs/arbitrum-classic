# FROM debian:buster-slim

# RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
#     apt-get install -y curl git golang librocksdb-dev sudo && \
#     rm -rf /var/lib/apt/lists/* && \
#     curl https://get.parity.io -L | bash /dev/stdin -r stable && \
#     useradd -ms /bin/bash user
# USER user
# WORKDIR /home/user/
# ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
# RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
#     chmod +x /home/user/bin/codecovbash && \
#     go get gotest.tools/gotestsum

FROM ethereum/client-go:stable

RUN apk add --no-cache bash curl findutils git go mercurial musl-dev sudo && \
    apk add rocksdb --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ --allow-untrusted && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/bin:/home/user/.local/bin:/home/user/.npm-global/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    go get gotest.tools/gotestsum
ENTRYPOINT ["/bin/ash"]
