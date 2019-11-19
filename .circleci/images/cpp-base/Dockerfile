### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:edge
# Alpine dependencies
RUN apk update && apk add --no-cache bash boost-dev cmake curl findutils g++ git \
    make musl-dev perl python3 python3-dev && \
    apk add rocksdb-dev --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ && \
    pip3 install --upgrade pip && \
    pip3 install conan && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user && \
    mkdir -p /usr/local/etc && \
    chown -R user:user /usr/local/bin /usr/local/share /usr/local/etc
USER user
WORKDIR "/home/user/"
RUN wget https://codecov.io/bash && chmod +x bash && \
    mv bash /usr/local/bin/codecovbash && \
    conan profile new default --detect && \
    conan profile update settings.compiler.libcxx=libstdc++11 default && \
    conan remote add nonstd-lite https://api.bintray.com/conan/martinmoene/nonstd-lite && \
    wget http://downloads.sourceforge.net/ltp/lcov-1.14.tar.gz && \
    tar -xvf lcov-1.14.tar.gz && cd lcov-1.14 && make install
