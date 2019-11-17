FROM alpine:3.10
# Alpine dependencies
RUN apk update && apk add --no-cache git go && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"
