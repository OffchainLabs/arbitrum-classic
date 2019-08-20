/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

"use strict";

// Code taken from https://github.com/ethers-io/ethers-web3-bridge
// with modifications made for compatibility

var providers = require("ethers-providers");
var utils = require("ethers-utils");

var Errors = {
  InternalError: -32603,
  InvalidRequest: -32600,
  ParseError: -32700,
  MethodNotFound: -32601,
  InvalidParams: -32602
};

// Some implementations of things do not play well with leading zeros
function smallHexlify(value) {
  value = utils.hexlify(value);
  while (value.length > 3 && value.substring(0, 3) === "0x0") {
    value = "0x" + value.substring(3);
  }
  return value;
}

// Convert a Web3 Transaction into an ethers.js Transaction
function makeTransaction(tx) {
  var result = {};
  ["data", "from", "gasPrice", "to", "value"].forEach(function(key) {
    if (tx[key] == null) {
      return;
    }
    result[key] = tx[key];
  });
  if (tx.gas != null) {
    result.gasLimit = tx.gas;
  }
  return result;
}

function fillCompact(values, result, keys, keepNull) {
  keys.forEach(function(key) {
    var value = values[key];
    if (value == null) {
      if (!keepNull) {
        return;
      }
      value = null;
    } else {
      value = smallHexlify(value);
    }
    result[key] = value;
  });
}

function fillCopy(values, result, keys, keepNull) {
  keys.forEach(function(key) {
    var value = values[key];
    if (value == null) {
      if (!keepNull) {
        return;
      }
      value = null;
    }
    result[key] = value;
  });
}

// Convert ethers.js Block into Web3 Block
function formatBlock(block) {
  var result = {};

  fillCompact(block, result, [
    "difficulty",
    "gasLimit",
    "gasUsed",
    "number",
    "timestamp"
  ]);

  fillCopy(block, result, ["extraData", "miner", "parentHash"]);

  fillCompact(block, result, ["number"], true);

  fillCopy(block, result, ["hash", "nonce"], true);

  return result;
}

// Convert ethers.js Transaction into Web3 Transaction
function formatTransaction(tx) {
  var result = {};

  if (tx.gasLimit) {
    result.gas = smallHexlify(tx.gasLimit);
  }
  result.input = tx.data || "0x";

  fillCompact(
    tx,
    result,
    ["blockNumber", "gasPrice", "nonce", "transactionIndex", "value"],
    true
  );

  fillCopy(tx, result, ["blockHash", "from", "hash", "to"], true);

  return result;
}

// Convert ethers.js Transaction Receiptinto Web3 Transaction
function formatReceipt(receipt) {
  var result = { logs: [] };

  fillCompact(
    receipt,
    result,
    [
      "blockNumber",
      "cumulativeGasUsed",
      "gasPrice",
      "gasUsed",
      "transactionIndex"
    ],
    true
  );

  fillCopy(
    receipt,
    result,
    [
      "blockHash",
      "contractAddress",
      "from",
      "logsBloom",
      "transactionHash",
      "root",
      "to"
    ],
    true
  );

  (receipt.logs || []).forEach(function(log) {
    result.logs.push(log);

    if (receipt.removed != null) {
      log.removed = receipt.removed;
    }
    if (receipt.topics != null) {
      log.topics = receipt.topics;
    }

    fillCompact(
      receipt,
      log,
      ["blockNumber", "logIndex", "transactionIndex"],
      true
    );

    fillCopy(
      receipt,
      log,
      ["address", "blockHash", "data", "transactionHash"],
      true
    );
  });

  return result;
}
// Convert ethers.js Log into Web3 Log
function formatLog(log) {
  var result = {};
  ["blockNumber", "logIndex", "transactionIndex"].forEach(function(key) {
    if (log[key] == null) {
      return;
    }
    result[key] = smallHexlify(log[key]);
  });
  ["address", "blockHash", "data", "topics", "transactionHash"].forEach(
    function(key) {
      if (log[key] == null) {
        return;
      }
      result[key] = log[key];
    }
  );
  return log;
}

function FilterManager(provider) {
  utils.defineProperty(this, "provider", provider);
  utils.defineProperty(this, "filters", {});

  var nextFilterId = 1;
  utils.defineProperty(this, "_getFilterId", function() {
    return nextFilterId++;
  });
}

FilterManager.prototype.addFilter = function(onblock, getLogs) {
  if (!getLogs) {
    getLogs = function() {
      return Promise.resolve([]);
    };
  }

  var filterId = this._getFilterId();

  var seq = Promise.resolve([]);
  var self = this;

  function emitBlock(blockNumber) {
    seq = seq.then(function(result) {
      return new Promise(function(resolve, reject) {
        function check() {
          self.provider.getBlock(blockNumber).then(
            function(block) {
              onblock(block, result).then(function(result) {
                resolve(result);
              });
            },
            function(error) {
              // Does not exist yet; try again in a second
              setTimeout(check, 1000);
            }
          );
        }
        check();
      });
    });
  }

  this.filters[smallHexlify(filterId)] = {
    getChanges: function() {
      var result = seq;

      // Reset the filter results
      seq = Promise.resolve([]);
      return result;
    },
    getLogs: getLogs,
    lastPoll: Date.now(),
    uninstall: function() {
      self.provider.removeListener("block", emitBlock);
      seq = null;
    }
  };

  self.provider.on("block", emitBlock);

  return smallHexlify(filterId);
};

FilterManager.prototype.removeFilter = function(filterId) {
  var filter = this.filters[smallHexlify(filterId)];
  if (!filter) {
    return false;
  }
  filter.uninstall();
  return true;
};

FilterManager.prototype.getChanges = function(filterId) {
  var filter = this.filters[smallHexlify(filterId)];
  if (!filter) {
    Promise.resolve([]);
  }
  return filter.getChanges();
};

FilterManager.prototype.getLogs = function(filterId) {
  var filter = this.filters[smallHexlify(filterId)];
  if (!filter) {
    return Promise.resolve([]);
  }
  return filter.getLogs();
};

var version = require("./package.json").version;

function ProviderBridge(provider, signer) {
  if (!(this instanceof ProviderBridge)) {
    throw new Error("missing new");
  }
  this._provider = provider || null;
  this._signer = signer || null;

  var self = this;
  setInterval(function() {
    if (!this._signer) {
      this._address = null;
      return;
    }

    this._signer.getAddress().then(
      function(address) {
        this._address = address;
      },
      function(error) {
        this._address = null;
      }
    );
  }, 1000);

  this._queue = [];

  utils.defineProperty(this, "isMetaMask", true);
  utils.defineProperty(this, "isEthers", true);
  utils.defineProperty(this, "isConnected", true);
  utils.defineProperty(this, "ethersVersion", version);
  utils.defineProperty(this, "client", "ethers/" + version);

  utils.defineProperty(this, "filterManager", new FilterManager(provider));
}

utils.defineProperty(ProviderBridge.prototype, "_drainQueue", function() {
  var self = this;
  this._queue.forEach(function(operation) {
    setTimeout(function() {
      self._sendAsync(JSON.parse(operation.payload), operation.callback);
    }, 0);
  });
});

utils.defineProperty(ProviderBridge.prototype, "connectWeb3", function(web3) {
  this._web3 = web3;
  this._drainQueue();
});

utils.defineProperty(ProviderBridge.prototype, "connectEthers", function(
  provider,
  signer
) {
  if (!signer) {
    var missingSigner = function() {
      return Promise.reject("no signer connected");
    };

    signer = {
      getAddress: missingSigner,
      sendTransaction: missingSigner,
      signMessage: missingSigner
    };
  }

  this._provider = provider;
  this._signer = signer;

  this._drainQueue();
});

utils.defineProperty(ProviderBridge.prototype, "sendAsync", function(
  payload,
  callback
) {
  if (!(this._provider || this._web3)) {
    this._queue.push({
      payload: JSON.stringify(payload),
      callback: callback
    });
    return;
  }
  this._sendAsync(payload, callback);
});

utils.defineProperty(ProviderBridge.prototype, "_sendAsync", function(
  payload,
  callback
) {
  if (this._web3) {
    this._web3.sendAsync(payload, callback);
    return;
  }

  var self = this;

  if (Array.isArray(payload)) {
    var promises = [];
    payload.forEach(function(payload) {
      promises.push(
        new Promise(function(resolve, reject) {
          self.sendAsync(payload, function(error, result) {
            resolve(error || result);
          });
        })
      );
    });

    Promise.all(promises).then(function(result) {
      callback(null, result);
    });

    return;
  }

  function respondError(message, code) {
    if (!code) {
      code = Errors.InternalError;
    }

    callback(null, {
      id: payload.id,
      jsonrpc: "2.0",
      error: {
        code: code,
        message: message
      }
    });
  }

  function respond(result) {
    callback(null, {
      id: payload.id,
      jsonrpc: "2.0",
      result: result
    });
  }

  if (
    payload == null ||
    typeof payload.method !== "string" ||
    typeof payload.id !== "number" ||
    !Array.isArray(payload.params)
  ) {
    respondError("invalid sendAsync parameters", Errors.InvalidRequest);
    return;
  }

  var signer = this._signer;
  var provider = this._provider;

  var params = payload.params;
  switch (payload.method) {
    // Account Actions

    case "eth_accounts":
      signer.getAddress().then(
        function(address) {
          respond([address.toLowerCase()]);
        },
        function(error) {
          respond([]);
        }
      );
      break;

    case "eth_sign":
      params = [params[1], params[0]];
    // Fall through

    case "personal_sign":
      signer.getAddress().then(
        function(address) {
          if (utils.getAddress(params[1]) !== address) {
            respondError("invalid from address", Errors.InvalidParams);
            return;
          }

          signer.signMessage(params[0]).then(
            function(signature) {
              respond(signature);
            },
            function(error) {
              respondError("eth_sign error", Errors.InternalError);
            }
          );
        },
        function(error) {
          respondError("no account", Errors.InvalidParams);
        }
      );

      break;

    case "eth_sendTransaction":
      signer.getAddress().then(
        function(address) {
          if (utils.getAddress(params[0].from) !== address) {
            respondError("invalid from address", Errors.InvalidParams);
          }
          signer.sendTransaction(params[0]).then(
            function(tx) {
              respond(tx.hash);
            },
            function(error) {
              console.error("eth_sendTransaction error", error);
              respondError("eth_sendTransaction error", Errors.InternalError);
            }
          );
        },
        function(error) {
          respondError("eth_sendTransaction error", Errors.InternalError);
        }
      );
      break;

    // Client State (mostly just default values we can pull from sync)

    case "eth_coinbase":
    case "eth_getCompilers":
    case "eth_hashrate":
    case "eth_mining":
    case "eth_syncing":
    case "net_listening":
    case "net_peerCount":
    case "net_version":
    case "eth_protocolVersion":
      setTimeout(function() {
        respond(self.send(payload).result);
      }, 0);
      break;

    // Blockchain state

    case "eth_blockNumber":
      provider.getBlockNumber().then(function(blockNumber) {
        respond(smallHexlify(blockNumber));
      });
      break;

    case "eth_gasPrice":
      provider.getGasPrice().then(function(gasPrice) {
        respond(smallHexlify(gasPrice));
      });
      break;

    // Accounts Actions

    case "eth_getBalance":
      provider.getBalance(params[0], params[1]).then(function(balance) {
        respond(smallHexlify(balance));
      });
      break;

    case "eth_getCode":
      provider.getCode(params[0], params[1]).then(function(code) {
        respond(code);
      });
      break;

    case "eth_getTransactionCount":
      provider.getTransactionCount(params[0], params[1]).then(function(nonce) {
        respond(smallHexlify(nonce));
      });
      break;

    // Execution (read-only)

    case "eth_call":
      provider.call(makeTransaction(params[0]), params[1]).then(
        function(data) {
          respond(data);
        },
        function(error) {
          console.error("eth_call error", error);
          respondError("eth_call error", Errors.InternalError);
        }
      );
      break;

    case "eth_estimateGas":
      provider
        .estimateGas(makeTransaction(params[0]), params[1])
        .then(function(data) {
          respond(data);
        });
      break;

    case "eth_getStorageAt":
      provider
        .getStorageAt(params[0], params[1], params[2])
        .then(function(data) {
          respond(data);
        });
      break;

    // Blockchain Queries

    case "eth_getBlockByHash":
    case "eth_getBlockByNumber":
      provider.getBlock(params[0]).then(function(block) {
        var result = formatBlock(block);

        if (params[1]) {
          result.transactions = [];

          var seq = Promise.resolve();

          if (block.transactions) {
            block.transactions.forEach(function(hash) {
              return provider.getTransaction(hash).then(function(tx) {
                result.transactions.push(tx);
              });
            });
          }

          seq.then(function() {
            respond(result);
          });
        } else {
          if (block.transactions) {
            result.transactions = block.transactions;
          }
          respond(result);
        }
      });
      break;

    case "eth_getBlockTransactionCountByHash":
    case "eth_getBlockTransactionCountByNumber":
      provider.getBlock(params[0]).then(function(block) {
        respond(
          smallHexlify(block.transactions ? block.transactions.length : 0)
        );
      });
      break;

    case "eth_getTransactionByHash":
      provider.getTransaction(params[0]).then(function(tx) {
        if (tx != null) {
          tx = formatTransaction(tx);
        }
        respond(tx);
      });
      break;

    case "eth_getTransactionByBlockHashAndIndex":
    case "eth_getTransactionByBlockNumberAndIndex":
      provider.getBlock(params[0]).then(function(block) {
        if (block == null) {
          block = {};
        }
        if (block.transactions == null) {
          block.transactions = [];
        }
        var hash = block.transactions[params[1]];
        if (hash) {
          provider.getTransaction(hash).then(function(tx) {
            if (tx != null) {
              tx = formatTransaction(tx);
            }
            respond(tx);
          });
        } else {
          respond(null);
        }
      });
      break;

    case "eth_getTransactionReceipt":
      provider.getTransactionReceipt(params[0]).then(function(receipt) {
        if (receipt != null) {
          receipt = formatReceipt(receipt);
        }
        respond(receipt);
      });
      break;

    // Blockchain Manipulation

    case "eth_sendRawTransaction":
      provider.sendTransaction(params[0]).then(function(hash) {
        respond(hash);
      });
      break;

    // Unsupported methods
    case "eth_getUncleByBlockHashAndIndex":
    case "eth_getUncleByBlockNumberAndIndex":
    case "eth_getUncleCountByBlockHash":
    case "eth_getUncleCountByBlockNumber":
      respondError("unsupported method", { method: payload.method });
      break;

    // Filters

    case "eth_newFilter":
      (function(filter) {
        function getLogs(filter) {
          return provider.getLogs(filter).then(function(result) {
            for (var i = 0; i < result.length; i++) {
              result[i] = formatLog(result[i]);
            }
            return result;
          });
        }

        respond(
          self.filterManager.addFilter(
            function(block, result) {
              var blockFilter = {
                fromBlock: block.number,
                toBlock: block.number
              };
              if (filter.address) {
                blockFilter.address = filter.address;
              }
              if (filter.topics) {
                blockFilter.topics = filter.topics;
              }
              return provider.getLogs(blockFilter).then(function(logs) {
                logs.forEach(function(log) {
                  log.blockHash = block.hash;
                  result.push(formatLog(log));
                });
                return result;
              });
            },
            function() {
              return provider.getLogs(filter).then(function(logs) {
                var seq = Promise.resolve(logs);
                logs.forEach(function(log) {
                  seq = seq.then(function() {
                    return provider
                      .getBlock(log.blockNumber)
                      .then(function(block) {
                        log.blockHash = block.hash;
                        return logs;
                      });
                  });
                });
                return seq;
              });
            }
          )
        );
      })(params[0]);
      break;

    case "eth_newPendingTransactionFilter":
      respond(
        this.filterManager.addFilter(function(block, result) {
          (block.transactions || []).forEach(function(hash) {
            result.push(hash);
          });
          result.push(block.hash);
          return Promise.resolve(result);
        })
      );
      break;

    case "eth_newBlockFilter":
      respond(
        this.filterManager.addFilter(function(block, result) {
          result.push(block.hash);
          return Promise.resolve(result);
        })
      );
      break;

    case "eth_uninstallFilter":
      respond(this.filterManager.removeFilter(params[0]));
      break;

    case "eth_getFilterChanges":
      this.filterManager.getChanges(params[0]).then(function(result) {
        respond(result);
      });
      break;

    case "eth_getFilterLogs":
      this.filterManager.getLogs(params[0]).then(function(result) {
        respond(result);
      });
      break;

    default:
      respondError("unknown method", Errors.MethodNotFound);
  }
});

utils.defineProperty(ProviderBridge.prototype, "send", function(payload) {
  if (this._web3) {
    return this._web3.send(payload);
  }

  var provider = this._provider;

  var result = null;
  switch (payload.method) {
    case "eth_accounts":
      if (this._address) {
        result = [this._address.toLowerCase()];
      } else {
        result = [];
      }
      break;

    case "eth_coinbase":
      result = null;
      break;

    case "eth_getCompilers":
      result = [];
      break;

    case "eth_hashrate":
    case "net_peerCount":
      result = "0x0";
      break;

    case "eth_mining":
    case "eth_syncing":
    case "net_listening":
      result = false;
      break;

    case "net_version":
      result = String(provider.chainId);
      break;

    // Using Parity/v1.8.0-beta-9882902-20171015/x86_64-macos/rustc1.20.0:
    // /Users/ethers>  curl -H 'Content-Type: application/json' -X POST --data '{"jsonrpc":"2.0","method":"eth_protocolVersion","params":[],"id":67}' http://localhost:8545
    // {"jsonrpc":"2.0","result":"63","id":67}
    case "eth_protocolVersion":
      result = "63";
      break;

    default:
      throw new Error("sync unsupported");
  }

  return {
    id: payload.id,
    jsonrpc: "2.0",
    result: result
  };
});

module.exports = ProviderBridge;
