### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:3.10
# Alpine dependencies
RUN apk update && apk add --no-cache g++ git make python2 yarn && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"
RUN yarn global add node-gyp
