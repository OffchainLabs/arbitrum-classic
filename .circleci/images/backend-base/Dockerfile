### --------------------------------------------------------------------
### Dockerfile
### backend-base
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.6.1

USER root
WORKDIR /
RUN export DEBIAN_FRONTEND=noninteractive && \
    curl https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz --output go1.17.3.linux-amd64.tar.gz && \
    tar -xf go1.17.3.linux-amd64.tar.gz && \
    mv go /usr/local && \
    rm *.tar.gz

USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:${PATH}" \
    GOROOT=/usr/local/go
RUN go install gotest.tools/gotestsum@v1.7.0 && \
    go clean --modcache
