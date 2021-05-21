### --------------------------------------------------------------------
### Dockerfile
### cpp-base
### --------------------------------------------------------------------

FROM debian:bullseye-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y cmake curl git g++ libboost-dev libboost-filesystem-dev \
    lcov librocksdb-dev make libgmp-dev libssl-dev \
    autotools-dev dh-autoreconf && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash
