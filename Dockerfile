### --------------------------------------------------------------------
### Dockerfile
### demo-app
### --------------------------------------------------------------------

FROM node:carbon

# Dependencies
RUN npm -g config set user root && npm install -g yarn
COPY package.json ./
ENV DEP=arb-web3-provider
COPY $DEP/package.json ./$DEP/package.json
RUN yarn && cd $DEP && yarn && yarn link && cd .. && yarn link $DEP
# Source code
COPY . ./
# Copy compiled.json to src
COPY --from=truffle-deploy-demo compiled.json ./src/
# Run server
CMD npm start
EXPOSE 8080
