### --------------------------------------------------------------------
### Dockerfile
### demo-app
### --------------------------------------------------------------------

FROM node:carbon-alpine

# Alpine dependencies
RUN apk add --no-cache g++ git make python

# Non-root user
USER node
ENV HOME="/home/node" PATH="/home/node/.npm-global/bin:${PATH}"
WORKDIR "${HOME}"
RUN mkdir "${HOME}/.npm-global" && \
    npm config set prefix "${HOME}/.npm-global" && npm install -g yarn

# Install project dependencies
COPY package.json ./
RUN yarn && yarn cache clean

# Copy source code
COPY . ./

# Run server
CMD yarn start
EXPOSE 8080
