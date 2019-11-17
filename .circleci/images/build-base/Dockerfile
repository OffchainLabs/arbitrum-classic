FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl gnupg && \
    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
    apt-get update && \
    apt-get install -y cmake g++ git golang lcov libboost-dev \
        libboost-filesystem-dev libffi-dev python3-dev python3-pip sudo yarn && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://get.parity.io -L | bash /dev/stdin -r stable && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN pip3 install coverage virtualenv && \
    virtualenv vconan && \
    . vconan/bin/activate && \
    pip install --upgrade conan && \
    conan profile new default --detect && \
    conan profile update settings.compiler.libcxx=libstdc++11 default && \
    conan remote add nonstd-lite https://api.bintray.com/conan/martinmoene/nonstd-lite && \
    deactivate && \
    curl -s https://codecov.io/bash > /home/user/.local/bin/codecovbash && \
    chmod +x /home/user/.local/bin/codecovbash && \
    go get gotest.tools/gotestsum && \
    yarn global add truffle
WORKDIR /home/user/
