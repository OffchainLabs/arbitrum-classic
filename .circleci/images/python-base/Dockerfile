FROM alpine:3.10
# Alpine dependencies
RUN apk update && apk add --no-cache gcc git musl-dev python3 python3-dev && \
    pip3 install --upgrade pip && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"
