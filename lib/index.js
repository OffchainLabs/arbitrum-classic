var eventEmitter = require('eventemitter3');
var BN = require("bn.js");
var jaysonBrowserClient = require('jayson/lib/client/browser');
var fetch = require('node-fetch');

class ArbitrumProvider extends eventEmitter {
  constructor(managerAddress) {
    // Call super for `this` to be defined
    super();

    this.managerAddress = managerAddress;

    // Init storage
    this._isConnected = false;
    this._nextJsonrpcId = 0;
    this._promises = {};

    // Fire the connect
    this._connect();

    // Listen for jsonrpc responses
    // window.addEventListener('message', this._handleJsonrpcMessage.bind(this));
    this.addDefaultMethods();
  }

  addDefaultMethods() {
    this._ethRPCMethods = {};
    this._ethRPCMethods['eth_accounts'] = this._ethAccounts;
    this._ethRPCMethods['eth_blockNumber'] = this._ethBlockNumber;
    this._ethRPCMethods['eth_call'] = this._ethCall;
    this._ethRPCMethods['eth_estimateGas'] = this._ethEstimateGas;
    this._ethRPCMethods['eth_gasPrice'] = this._ethGasPrice;
    this._ethRPCMethods['eth_getBalance'] = this._ethGetBalance;
    this._ethRPCMethods['eth_getBlockByHash'] = this._ethGetBlockByHash;
    this._ethRPCMethods['eth_getBlockByNumber'] = this._ethGetBlockByNumber;
    this._ethRPCMethods['eth_getCode'] = this._ethGetCode;
    this._ethRPCMethods['eth_getFilterChanges'] = this._ethGetFilterChanges;
    this._ethRPCMethods['eth_getLogs'] = this._ethGetLogs;
    this._ethRPCMethods['eth_getTransactionByHash'] = this._ethGetTransactionByHash;
    this._ethRPCMethods['eth_getTransactionReceipt'] = this._ethGetTransactionReceipt;
    this._ethRPCMethods['eth_newBlockFilter'] = this._ethNewBlockFilter;
    this._ethRPCMethods['eth_newFilter'] = this._ethNewFilter;
    this._ethRPCMethods['eth_newPendingTransactionFilter'] = 
      this._ethNewPendingTransactionFilter;
    this._ethRPCMethods['eth_sendTransaction'] = this._ethSendTransaction;
    this._ethRPCMethods['eth_sign'] = this._ethSign;
    this._ethRPCMethods['eth_subscribe'] = this._ethSubscribe;
    this._ethRPCMethods['eth_uninstallFilter'] = this._ethUninstallFilter;
    this._ethRPCMethods['eth_unsubscribe'] = this._ethUnsubscribe;
    this._ethRPCMethods['net_version'] = this._netVersion;
  }

  _ethAccounts() {
    return ["0x81183C9C61bdf79DB7330BBcda47Be30c0a85064"]
  }

  _ethGasPrice() {
    return 0;
  }

  async _ethCall(payload) {
    if (!payload.params[0].to) {
      return new Error('Contract creation is not supported.');
    }

    this.from = payload.params[0].from
    this.to = payload.params[0].to
    let self = this;
    let data = {
      "address": payload.params[0].to,
      "data": payload.params[0]["data"],
      "sender": payload.params[0].from
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
  }

  async _ethSendTransaction(payload) {
    if (!payload.params[0].to) {
      return new Error('Contract creation is not supported.');
    }
    this.from = payload.params[0].from
    this.to = payload.params[0].to
    let data = {
      "address": payload.params[0].to,
      "data": payload.params[0]["data"],
      "amount": "0x00",
      "tokenType": "0x00",
      "sender": payload.params[0].from,
      "destination": payload.params[0].to
    }
    let self = this
    return new Promise(function(resolve, reject) {
      self.client.request(
        'Validator.SendMessage',
        [data],
        function(err, error, result) {
          if (err) {
            reject(error)
          } else {
            resolve(result["hash"])
          }
        }
      )
    });
  }

  _ethGetTransactionReceipt(txhash) {
    let self = this
    let data = {
      "txHash": txhash.params[0]
    }
    return new Promise(function(resolve, reject) {
      self.client.request(
        'Validator.GetMessageResult',
        [data],
        function(err, error, result) {
          if (err) {
            reject(error)
          } else {
            if (result["found"]) {
              var status = 0
              if (result["success"]) {
                status = 1
              }
              let receipt = {
                "transactionHash": txhash.params[0],
                "transactionIndex": 0,
                "blockHash": "0x10000",
                "blockNumber": 0,
                "from": self.from,
                "to": self.to,
                "cumulativeGasUsed": 1,
                "gasUsed": 1,
                "contractAddress": null,
                "logs": null,
                "logsBloom": null,
                "status": status
              }
              resolve(receipt);
            } else {
              resolve(self._ethGetTransactionReceipt(txhash))
            }
          }
        }
      )
    });
  }

  /* Methods */

  async send(method, params = []) {
    if (!method || typeof method !== 'string') {
      return new Error('Method is not a valid string.');
    }

    if (!(params instanceof Array)) {
      return new Error('Params is not a valid array.');
    }

    const id = this._nextJsonrpcId++;
    const jsonrpc = '2.0';
    const payload = { jsonrpc, id, method, params };

    const methodToCall = this._ethRPCMethods[method];
    if (!methodToCall) {
      throw Error(`Method "${payload.method}" not supported on this provider`)
    }

    let result = await methodToCall.bind(this)(payload);
    return result;
  }

  /* Internal methods */

  _handleJsonrpcMessage(event) {
    console.log("_handleJsonrpcMessage", event);
    // Return if no data to parse
    if (!event || !event.data) {
      return;
    }

    let data;
    try {
      data = JSON.parse(event.data);
    } catch (error) {
      // Return if we can't parse a valid object
      return;
    }

    // Return if not a jsonrpc response
    if (!data || !data.message || !data.message.jsonrpc) {
      return;
    }

    const message = data.message;
    const { id, method, error, result } = message;

    if (typeof id !== 'undefined') {
      const promise = this._promises[id];
      if (promise) {
        // Handle pending promise
        if (data.type === 'error') {
          promise.reject(message);
        } else if (message.error) {
          promise.reject(error);
        } else {
          promise.resolve(result);
        }
        delete this._promises[id];
      }
    } else {
      if (method && method.indexOf('_subscription') > -1) {
        // Emit subscription notification
        this._emitNotification(message.params);
      }
    }
  }

  /* Connection handling */

  _connect() {
    // Send to Mist
    console.log("Connecting");
    let self = this;

    var callServer = function(request, callback) {
      var options = {
        method: 'POST',
        body: request, // request is a string
        headers: {
          'Content-Type': 'application/json',
        }
      };

      fetch(self.managerAddress, options)
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

    this.client = jaysonBrowserClient(callServer, {});

    // window.postMessage(
    //   { type: 'mistAPI_ethereum_provider_connect' },
    //   targetOrigin
    // );

    // Reconnect on close
    this.once('close', this._connect.bind(this));
  }

  /* Events */

  _emitNotification(result) {
    this.emit('notification', result);
  }

  _emitConnect() {
    this._isConnected = true;
    this.emit('connect');
  }

  _emitClose(code, reason) {
    this._isConnected = false;
    this.emit('close', code, reason);
  }

  _emitNetworkChanged(networkId) {
    this.emit('networkChanged', networkId);
  }

  _emitAccountsChanged(accounts) {
    this.emit('accountsChanged', accounts);
  }

  /* web3.js Provider Backwards Compatibility */

  sendAsync(payload, callback) {
    return this.send(payload.method, payload.params)
      .then(result => {
        const response = payload;
        response.result = result;
        callback(null, response);
      })
      .catch(error => {
        callback(error, null);
        // eslint-disable-next-line no-console
        console.error(
          `Error from EthereumProvider sendAsync ${payload}: ${error}`
        );
      });
  }

  isConnected() {
    return this._isConnected;
  }
}

module.exports = ArbitrumProvider