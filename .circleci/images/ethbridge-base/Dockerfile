FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && \
    apt-get install -y curl gnupg && \
    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
    apt-get update && \
    apt-get install -y git sudo yarn && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://get.parity.io -L | bash /dev/stdin -r stable && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="~/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    yarn global add truffle
