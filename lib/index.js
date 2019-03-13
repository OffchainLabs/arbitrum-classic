const ethers = require('ethers');
var jaysonBrowserClient = require('jayson/lib/client/browser');
var fetch = require('node-fetch');
var promisePoller = require('promise-poller').default;

function _arbClient(managerAddress) {
  var callServer = function(request, callback) {
    var options = {
      method: 'POST',
      body: request, // request is a string
      headers: {
        'Content-Type': 'application/json',
      }
    };

    fetch(managerAddress, options)
      .then(function(res) {
        return res.text(); 
      })
      .then(function(text) {
        callback(null, text); 
      })
      .catch(function(err) { 
        callback(err); 
      });
  };

  return jaysonBrowserClient(callServer, {});
}

class ArbProvider extends ethers.providers.BaseProvider {
    constructor(managerUrl, contracts, provider) {
        super(provider.getNetwork());
        this.provider = provider
        this.client = _arbClient(managerUrl)
        this.contracts = {}
        for (var contract of contracts) {
          this.contracts[contract.address.toLowerCase()] = contract
        }
    }

    getSigner(index) {
      return new ArbWallet(this.client, this.contracts, this.provider.getSigner(index), this);
    }

    // This should return a Promise (and may throw errors)
    // method is the method name (e.g. getBalance) and params is an
    // object with normalized values passed in, depending on the method
    perform(method, params) {
      var _this = this
      switch(method) {
        case "getCode":
          if (this.contracts[params.address.toLowerCase()]) {
            return new Promise((resolve, reject) => {
              resolve(this.contracts[params.address.toLowerCase()].code)
            })
          }
          break
        case "getTransaction":
          let data = {
            "txHash": params.transactionHash
          }
          var getMessageRequest = () => new Promise(function(resolve, reject) {
            _this.client.request(
              'Validator.GetMessageResult',
              [data],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  if (result["found"]) {
                    let receipt = {
                      "hash": data.txHash,
                      "transactionIndex": 0,
                      "blockHash": data.txHash,
                      "confirmations": 1000,
                      "blockNumber": 0,
                      "from": result["sender"],
                      "to": result["address"],
                      "cumulativeGasUsed": 1,
                      "gasPrice": 1,
                      "gasLimit": 1,
                      "value": result["amount"],
                      "contractAddress": null,
                      "logs": null,
                      "logsBloom": null,
                      "status": result["success"],
                      "data": result["data"],
                      "nonce": 0
                    }
                    resolve(receipt);
                  } else {
                    resolve(null)
                  }
                }
              }
            )
          });
          return promisePoller({
            taskFn: getMessageRequest,
            interval: 100,
            shouldContinue: (reason, value) => {
              if (reason) {
                return true
              } else if (value) {
                return false
              } else {
                return true
              }
            }
          });
          return poller
      }
      return this.provider.perform(method, params)
    }

    async call(transaction) {
      let dest = await transaction.to
      let _this = this;
      let contractData = _this.contracts[dest.toLowerCase()]
      if (contractData) {
        let data = {
          "address": await transaction.to,
          "data": transaction.data,
          "sender": await transaction.from
        }
        return new Promise(function(resolve, reject) {
          _this.client.request(
            'Validator.CallMessage',
            [data],
            function(err, error, result) {
              if (err) {
                reject(error)
              } else {
                resolve(result["ReturnVal"])
              }
            }
          )
        });
      } else {
        return this.provider.call(transaction)
      }
      
    }
}

class ArbWallet extends ethers.Signer {
    constructor(client, contracts, signer, provider) {
        super()
        this.contracts = contracts
        this.signer = signer
        this.provider = provider
        this.client = client
    }

    getAddress() {
      return this.signer.getAddress()
    }

    async sendTransaction(transaction) {
      var _this = this
      let dest = await transaction.to
      if (_this.contracts[dest.toLowerCase()]) {
        let data = {
          "address": dest,
          "data": transaction.data,
          "amount": "0x00",
          "tokenType": "0x00",
          "sender": await _this.signer.getAddress(),
          "destination": dest
        }
        return new Promise(function(resolve, reject) {
          _this.client.request(
            'Validator.SendMessage',
            [data],
            function(err, error, result) {
              if (err) {
                reject(error)
              } else {
                data.hash = result["hash"]
                _this.provider.getTransaction(result["hash"]).then(tx => {
                  let ret = _this.provider._wrapTransaction(tx, result["hash"])
                  resolve(ret)
                })
              }
            }
          )
        });
      } else {
        return this.signer.sendTransaction(transaction)
      }
    };
}

module.exports = ArbProvider