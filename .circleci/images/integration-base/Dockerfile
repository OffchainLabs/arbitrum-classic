### --------------------------------------------------------------------
### Dockerfile
### integration-base
### --------------------------------------------------------------------

FROM offchainlabs/backend-base:0.6.2

USER root
WORKDIR /

# Node and yarn are installed in cpp-base, which this image inherits through
# backend-base. The NodeSource setup_14.x script that used to run here no longer
# publishes a usable signing key, and it was only re-installing what is already
# present.
RUN node --version && yarn --version
USER user
WORKDIR /home/user/
ENV PATH="/home/user/go/bin:/home/user/bin:/home/user/.local/bin:/usr/local/go/bin:/usr/local/cargo/bin:/home/user/.npm-global/bin:/home/user/.yarn/bin:${PATH}" \
    GOROOT=/usr/local/go
