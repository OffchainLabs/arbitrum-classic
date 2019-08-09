FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y cmake curl g++ git golang lcov libboost-dev \
        libffi-dev npm python3-dev python3-pip && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/.local/bin:/home/user/.npm-global/bin:${PATH}"
RUN pip3 install coverage virtualenv && \
    virtualenv arbitrum/vconan && \
    . arbitrum/vconan/bin/activate && \
    pip install --upgrade conan && \
    conan remote add nonstd-lite https://api.bintray.com/conan/martinmoene/nonstd-lite && \
    deactivate && \
    mkdir /home/user/.npm-global && npm config set prefix '~/.npm-global' && \
    npm install -g ganache-cli@6.5.0 truffle@5.0.30 yarn@1.17.3
WORKDIR /home/user/arbitrum
