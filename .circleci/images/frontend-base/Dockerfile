FROM node:14.17.4-alpine3.12

RUN apk add --no-cache bash curl findutils git mercurial psmisc python2 sudo libusb-dev linux-headers make pkgconfig eudev-dev g++
USER node
WORKDIR /home/node/
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/node/bin/codecovbash
ENTRYPOINT ["/bin/ash"]
