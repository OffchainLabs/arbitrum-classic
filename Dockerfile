### --------------------------------------------------------------------
### Dockerfile
### Launch Demo
### --------------------------------------------------------------------

FROM node:carbon as demo

# Dependencies
RUN npm -g config set user root && npm install -g yarn
COPY package.json ./
RUN yarn

# Source code
COPY . ./

# Run server
CMD npm start
EXPOSE 8080
