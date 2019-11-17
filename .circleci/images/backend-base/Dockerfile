FROM alpine:3.10
# Alpine dependencies
RUN apk add --no-cache bash curl git go libc-dev libstdc++ linux-headers && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user && \
    wget https://codecov.io/bash -O codecovbash && chmod +x codecovbash && \
    mv codecovbash /usr/local/bin/
USER user
WORKDIR "/home/user"
ENV PATH="/home/user/go/bin:${PATH}"
RUN go get gotest.tools/gotestsum
