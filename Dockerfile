### --------------------------------------------------------------------
### Dockerfile
### demo-app
### --------------------------------------------------------------------

FROM node:carbon-alpine

# Alpine dependencies
RUN apk add --no-cache g++ git make python

# Non-root user
USER node
ENV HOME="/home/node" PATH="/home/node/.npm-global/bin:${PATH}" \
    DEP="arb-web3-provider"
WORKDIR "${HOME}"
RUN mkdir "${HOME}/.npm-global" && mkdir "${HOME}/${DEP}" && \
    npm config set prefix "${HOME}/.npm-global" && \
# Dependencies \
    npm install -g yarn
COPY package.json ./
RUN yarn && yarn cache clean
COPY arb-web3-provider/package.json ./arb-web3-provider/package.json
RUN cd arb-web3-provider && yarn && yarn cache clean

# Source code
COPY . ./

# Copy compiled.json to src
COPY --from=truffle-deploy-demo compiled.json ./src/

# Run server
CMD yarn start
EXPOSE 8080
