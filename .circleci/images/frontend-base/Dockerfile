FROM alpine:3.10
# Alpine dependencies
RUN apk update && apk add --no-cache git python3 yarn && \
    pip3 install --upgrade pip && \
    pip3 install coverage && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"
ENV PATH="~/bin:/home/user/.local/bin:/home/user/.yarn/bin:${PATH}"
RUN mkdir bin && curl -s https://codecov.io/bash > ~/bin/codecovbash && \
    chmod +x /home/user/bin/codecovbash && \
    yarn global add truffle
