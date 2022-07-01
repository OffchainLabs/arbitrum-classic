### --------------------------------------------------------------------
### Dockerfile
### frontend-base
### --------------------------------------------------------------------

FROM node:14.17.4-alpine3.12

RUN apk add --no-cache bash curl findutils git mercurial psmisc python2 sudo libusb-dev linux-headers make pkgconfig eudev-dev g++ && \
    deluser --remove-home node && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user

USER user
WORKDIR /home/user/
ENV PATH="~/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash
ENTRYPOINT ["/bin/ash"]
