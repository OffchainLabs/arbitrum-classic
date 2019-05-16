### --------------------------------------------------------------------
### Dockerfile
### arbc-solidity
### --------------------------------------------------------------------

FROM python:3.7
# Dependencies
COPY requirements.txt ./
RUN pip3 install -r requirements.txt
# Source code
COPY . ./
RUN python3 setup.py install
COPY --from=truffle-deploy-demo compiled.json ./
RUN arbc-truffle-compile compiled.json contract.ao

# Minimize
FROM scratch
COPY --from=0 contract.ao /usr/local/bin/arbc-truffle-compile \
/usr/local/bin/arbc-compile ./
