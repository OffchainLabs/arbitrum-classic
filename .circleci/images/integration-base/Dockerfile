### --------------------------------------------------------------------
### Dockerfile
### integration-base
### --------------------------------------------------------------------

FROM offchainlabs/backend-base:0.6.1

USER root
WORKDIR /

RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash - && \
    curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor > /usr/share/keyrings/yarnkey.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main" > /etc/apt/sources.list.d/yarn.list && \
    apt-get update && apt-get install yarn
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:/usr/local/cargo/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}" \
    GOROOT=/usr/local/go
