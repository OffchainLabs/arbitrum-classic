FROM ethereum/client-go:stable

RUN apk add --no-cache bash curl findutils git mercurial nodejs npm psmisc python2 sudo && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/bin:/home/user/.local/bin:/home/user/.npm-global/bin:${PATH}"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g truffle yarn && \
    mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash
ENTRYPOINT ["/bin/ash"]
