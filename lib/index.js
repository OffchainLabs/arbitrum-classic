const ganache = require("ganache-core");
const path = require("path");
const fs = require("fs");
const callsite = require('callsite');
const spawnSync = require("child_process").spawnSync;

function provider(outputFolder, buildLocation, options) {
  var outputLocation = path.resolve(outputFolder, 'compiled.json')
  var outputAOLocation = path.resolve(outputFolder, 'contract.ao')
  var stack = callsite()
  var rootPath = path.dirname(stack[1].getFileName())
  if (!buildLocation) {
    buildLocation = path.resolve(rootPath, 'build/contracts');
  }
  options["network_id"] = 123456789
  let arbProvider = ganache.provider(options)
  var storage = {}
  var netID = arbProvider.options.network_id;
  arbProvider.engine.manager.waitForInitialization(function(err, state) {
    state.blockchain.vm.on("step", function(info) {
      var address_string = '0x' + info.address.toString('hex')
      if (!(address_string in storage)) {
        storage[address_string] = {}
      }
      if (info.opcode.name == "SSTORE") {
        var args = info.stack
          .slice(-info.opcode.in)
          .map((arg) => '0x' + arg.toString('hex'))
        
        storage[address_string][args[1]] = args[0]
      }
    });
  })
  process.on('exit', (code) => {
    var contracts = []
    var files = fs.readdirSync(buildLocation, {})
    for (var filePath of files) {
      var contract = JSON.parse(fs.readFileSync(path.resolve(buildLocation, filePath)))
      var networkInfo = contract.networks[netID]
      if (networkInfo) {
        var address = networkInfo.address
        contracts.push({
          name: contract.contractName,
          address: address,
          code: contract.deployedBytecode,
          storage: storage[address.toLowerCase()], 
          abi: contract.abi
        })
      }
    }
    fs.writeFileSync(outputLocation, JSON.stringify(contracts, null, 2))
    compile = spawnSync("arbc-truffle-compile", [outputLocation, outputAOLocation]);
  })
  return arbProvider;
}
module.exports = {
  provider: provider
}
