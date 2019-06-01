const ethers = require('ethers');
var jaysonBrowserClient = require('jayson/lib/client/browser');
var fetch = require('node-fetch');
var promisePoller = require('promise-poller').default;
const vmTrackerJson = require('./VMTracker.json');
var ArbValue = require("./arb-value")

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
      // console.log("perform", method, params)
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
                  resolve(result["assertionCount"])
                }
              }
            )
          });
        case "getTransactionReceipt":
          return new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.GetMessageResult',
              [{
                "txHash": params.transactionHash
              }],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  if (result["found"]) {
                    let status = 0
                    if (result["success"]) {
                      status = 1;
                    }
                    let receipt = {
                      "to": result["address"],
                      "from": result["sender"],
                      "transactionIndex": 0,
                      "gasUsed": 1,
                      "blockHash": params.transactionHash,
                      "transactionHash": params.transactionHash,
                      "logs": [],
                      "blockNumber": result["assertionNum"],
                      "confirmations": 1000,
                      "cumulativeGasUsed": 1,
                      "status": status
                    }
                    resolve(receipt);
                  } else {
                    resolve(null)
                  }
                }
              }
            )
          });

        case "getTransaction":
          var getMessageRequest = () => new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.GetMessageResult',
              [{
                "txHash": params.transactionHash
              }],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  if (result["found"]) {
                    let receipt = {
                      "hash": params.transactionHash,
                      "blockHash": params.transactionHash,
                      "blockNumber": result["assertionNum"],
                      "transactionIndex": 0,
                      "confirmations": 1000,
                      "from": result["sender"],
                      "gasPrice": 1,
                      "gasLimit": 1,
                      "to": result["address"],
                      "cumulativeGasUsed": 1,
                      "value": result["amount"],
                      "nonce": 0,
                      "data": result["data"],
                      "status": result["success"]
                    }
                    console.log(receipt)
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
        case "getLogs":
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
                  resolve(result["logs"])
                }
              }
            )
          });
      }
      let forwardResponse = self.provider.perform(method, params);
      console.log("Forwarding query to provider", method, forwardResponse);
      return forwardResponse;
    }

    async call(transaction) {
      let dest = await transaction.to
      let self = this;
      let contractData = self.contracts[dest.toLowerCase()]
      if (contractData) {
        let data = {
          "address": await transaction.to,
          "data": transaction.data,
          // Call doesn't have sender specified
          "sender": await this.provider.getSigner(0).getAddress()
        }
        return new Promise(function(resolve, reject) {
          self.client.request(
            'Validator.CallMessage',
            [data],
            function(err, error, result) {
              if (err) {
                reject(err)
              } else if (error) {
                reject(error)
              } else {
                if (result["Success"]) {
                  resolve(result["ReturnVal"])
                } else {
                  reject("Call failed")
                }
                
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

        let self = this
        provider.provider.getBlockNumber().then(height => {
          var seq = ethers.utils.bigNumberify(height)
          for (var i = 0; i < 128; i++) {
            seq = seq.mul(2);
          }
          var timeStamp = Math.floor(Date.now());
          seq = seq.add(timeStamp)
          seq = seq.mul(2);
          self.seq = seq;

        })
    }

    getAddress() {
      return this.signer.getAddress()
    }

    async sendTransaction(transaction) {
      let self = this
      self.seq = self.seq.add(2)

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
          let seqHex = self.seq.toHexString()
          let encodedData = ArbValue.hexToSizedByteRange(transaction.data)
          let arbMsg = new ArbValue.TupleValue([
            encodedData,
            new ArbValue.IntValue(dest),
            new ArbValue.IntValue(self.seq)
          ]);
          let messageHash = ethers.utils.solidityKeccak256(
            [ 'bytes32', 'bytes32', 'uint256', 'bytes21'],
            [ vmId, arbMsg.hash(), "0x00", ethers.utils.hexZeroPad("0x00", 21) ]
          );

          let messageHashBytes = ethers.utils.arrayify(messageHash)
          let sig = await self.signer.signMessage(messageHashBytes)
          let recovered = ethers.utils.recoverAddress(messageHashBytes, sig);
          
          let data = {
            "data": ArbValue.marshal(arbMsg),
            "signature": sig
          }
          let txHash = await new Promise(function(resolve, reject) {
            self.client.request(
              'Validator.SendMessage',
              [data],
              function(err, error, result) {
                if (err) {
                  reject(error)
                } else {
                  resolve(result["hash"])
                }
              });
          })
          let tx = await self.provider.getTransaction(txHash)
          let ret = self.provider._wrapTransaction(tx, txHash)
          return ret
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