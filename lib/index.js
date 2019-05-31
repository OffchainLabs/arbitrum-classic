const ethers = require('ethers');
var jaysonBrowserClient = require('jayson/lib/client/browser');
var fetch = require('node-fetch');
var promisePoller = require('promise-poller').default;
const vmTrackerJson = require('./VMTracker.json');

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
        super(123456789)
        this.chainId = 123456789;
        this.provider = provider
        this.client = _arbClient(managerUrl)
        let contractAddress = "0x5EBF59dBff8dCDa41610738634b396DfCB24A7c7";
        this.vmTracker = new ethers.Contract(contractAddress, vmTrackerJson["abi"], provider);
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
      console.log("perform", method, params)
      var self = this
      switch(method) {
        case "getCode":
          if (self.contracts[params.address.toLowerCase()]) {
            return new Promise((resolve, reject) => {
              resolve(self.contracts[params.address.toLowerCase()].code)
            })
          }
          break
        case "getBlockNumber":
          return new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.GetAssertionCount',
              [],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  console.log("assertionCount", result["assertionCount"])
                  resolve(result["assertionCount"])
                }
              }
            )
          });
        case "getTransaction":
          let data = {
            "txHash": params.transactionHash
          }
          var getMessageRequest = () => new Promise(function(resolve, reject) {
            self.client.request(
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
                      "blockNumber": result["assertionNum"],
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
        case "getLogs":
        console.log(params.filter)
          return new Promise(function(resolve, reject) {
            let data = {
                "fromHeight": params.filter.fromBlock,
                "toHeight": params.filter.toBlock,
                "address": params.filter.address,
                "topics": params.filter.topics
            }
            self.client.request(
              'Validator.FindLogs',
              [data],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  console.log("FindLogs", result["logs"])
                  resolve(result["logs"])
                }
              }
            )
          });
      }
      let forwardResponse = self.provider.perform(method, params);
      console.log("forwardResponse", forwardResponse);
      return forwardResponse;
    }

    async call(transaction) {
      console.log("CALL", transaction);
      let dest = await transaction.to
      let self = this;
      let contractData = self.contracts[dest.toLowerCase()]
      if (contractData) {
        let data = {
          "address": await transaction.to,
          "data": transaction.data,
          "sender": await transaction.from
        }
        return new Promise(function(resolve, reject) {
          self.client.request(
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
        return self.provider.call(transaction)
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
        this.vmTracker = provider.vmTracker.connect(signer);
    }

    getAddress() {
      return this.signer.getAddress()
    }

    async sendTransaction(transaction) {
      console.log("sendTransaction", transaction)
      let self = this

      let dest = await transaction.to
      if (self.contracts[dest.toLowerCase()]) {
        let vmId = await new Promise(function(resolve, reject) {
          self.client.request(
            'Validator.GetVMInfo',
            [],
            function(err, error, result) {
              if (err) {
                reject(error)
              } else {
                resolve(result["vmId"])
              }
            }
          )
        });
        if (!transaction.value || transaction.value == 0) {
          console.log("Sending message")
          let seq = ethers.utils.bigNumberify(10);
          let seqHex = seq.toHexString()
          console.log("seq", seq, seqHex)
          let data = {
            "address": dest,
            "data": transaction.data,
            "sequenceNum": seqHex
          }
          let dataHash = await new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.TranslateToValueHash',
              [data],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  resolve(result)
                }
              }
            )
          })
          console.log("Got data hash", dataHash)
          console.log("Hashing message", [ vmId, dataHash, "0x00", ethers.utils.hexZeroPad("0x00", 21) ])
          let messageHash = ethers.utils.solidityKeccak256(
            [ 'bytes32', 'bytes32', 'uint256', 'bytes21'],
            [ vmId, dataHash, "0x00", ethers.utils.hexZeroPad("0x00", 21) ]
          );

          console.log("Got messageHash", messageHash)

          let messageHashBytes = ethers.utils.arrayify(messageHash)
          let sig = await self.signer.signMessage(messageHashBytes)
          let recovered = ethers.utils.recoverAddress(messageHashBytes, sig);
          

          console.log("sig:", sig)
          data["signature"] = sig
          return new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.SendMessage',
              [data],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  console.log("Sent message with hash", result["hash"])
                  data.hash = result["hash"]
                  self.provider.getTransaction(result["hash"]).then(tx => {
                    let ret = self.provider._wrapTransaction(tx, result["hash"])
                    resolve(ret)
                  })
                }
              }
            )
          });
        } else {
          let data = await new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.TranslateToValue',
              [transaction.data],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  resolve(result)
                }
              }
            )
          })
          return self.vmTracker.sendEthMessage(vmId, data, {
            value: transaction.value
          });
          
        }
      } else {
        return self.signer.sendTransaction(transaction)
      }
    };
}

module.exports = ArbProvider