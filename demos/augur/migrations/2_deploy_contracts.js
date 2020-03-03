var Augur = artifacts.require("./Augur.sol");
var AugurTrading = artifacts.require("./AugurTrading.sol");
var WarpSync = artifacts.require("./WarpSync.sol");
var EthExchange = artifacts.require("./EthExchange.sol");

var GnosisSafe = artifacts.require("./GnosisSafe.sol");
var ProxyFactory = artifacts.require("./ProxyFactory.sol");

var GnosisSafeRegistry = artifacts.require("./GnosisSafeRegistry.sol");
var RepExchange = artifacts.require("./RepExchange.sol");
var Universe = artifacts.require("./reporting/Universe.sol");
var TestNetReputationToken = artifacts.require("./TestNetReputationToken.sol");

var CreateOrder = artifacts.require("./trading/CreateOrder.sol");
var Trade = artifacts.require("./trading/Trade.sol");
var Order = artifacts.require("./trading/Order.sol");
var ProfitLoss = artifacts.require("./trading/ProfitLoss.sol");
var SimulateTrade = artifacts.require("./trading/SimulateTrade.sol");
var ZeroXTrade = artifacts.require("./trading/ZeroXTrade.sol");
var LegacyReputationToken = artifacts.require(
  "./trading/LegacyReputationToken.sol"
);

var BuyParticipationTokens = artifacts.require(
  "./utility/BuyParticipationTokens.sol"
);
var Formulas = artifacts.require("./utility/Formulas.sol");
var HotLoading = artifacts.require("./utility/HotLoading.sol");
var RedeemStake = artifacts.require("./utility/RedeemStake.sol");
var RepSymbol = artifacts.require("./utility/RepSymbol.sol");

var AffiliateValidator = artifacts.require(
  "./reporting/AffiliateValidator.sol"
);
var Affiliates = artifacts.require("./reporting/Affiliates.sol");
var DisputeCrowdsourcer = artifacts.require(
  "./reporting/DisputeCrowdsourcer.sol"
);
var InitialReporter = artifacts.require("./reporting/InitialReporter.sol");
var Market = artifacts.require("./reporting/Market.sol");
var OICash = artifacts.require("./reporting/OICash.sol");
var ReputationToken = artifacts.require("./reporting/ReputationToken.sol");
var ShareToken = artifacts.require("./reporting/ShareToken.sol");
var DisputeWindow = artifacts.require("./reporting/DisputeWindow.sol");

var DisputeCrowdsourcerFactory = artifacts.require(
  "./factories/DisputeCrowdsourcerFactory.sol"
);
var DisputeWindowFactory = artifacts.require(
  "./factories/DisputeWindowFactory.sol"
);
var InitialReporterFactory = artifacts.require(
  "./factories/InitialReporterFactory.sol"
);
var MarketFactory = artifacts.require("./factories/MarketFactory.sol");
var OICashFactory = artifacts.require("./factories/OICashFactory.sol");
var RepExchangeFactory = artifacts.require(
  "./factories/RepExchangeFactory.sol"
);
var ReputationTokenFactory = artifacts.require(
  "./factories/ReputationTokenFactory.sol"
);
var TestNetReputationTokenFactory = artifacts.require(
  "./factories/TestNetReputationTokenFactory.sol"
);
var UniverseFactory = artifacts.require("./factories/UniverseFactory.sol");

//--optimize-off

module.exports = async function(deployer) {
  let auagur = await deployer.deploy(Augur);
  await deployer.deploy(AugurTrading, Augur.address);
  await deployer.deploy(GnosisSafe);
  // await deployer.deploy(ProxyFactory);
  await deployer.deploy(GnosisSafeRegistry);
  await deployer.deploy(WarpSync);
  await deployer.deploy(CreateOrder);
  await deployer.deploy(Order);
  await deployer.deploy(ProfitLoss);
  await deployer.deploy(SimulateTrade);
  await deployer.deploy(LegacyReputationToken);
  await deployer.deploy(BuyParticipationTokens);
  await deployer.deploy(Formulas);
  await deployer.deploy(HotLoading);
  await deployer.deploy(RedeemStake);
  await deployer.deploy(RepSymbol);
  await deployer.deploy(ShareToken);

  await deployer.deploy(AffiliateValidator);
  let afiliates = await deployer.deploy(Affiliates);
  afiliates.initializeValidatorFactory(AffiliateValidator.address);

  await deployer.deploy(InitialReporter);
  let initialReporterFactory = await deployer.deploy(InitialReporterFactory);
  initialReporterFactory.initializeFactory(InitialReporter.address);

  await deployer.deploy(Market);
  let marketFactory = await deployer.deploy(MarketFactory);
  marketFactory.initializeFactory(Market.address);

  await deployer.deploy(OICash);
  let _OICashFactory = await deployer.deploy(OICashFactory);
  _OICashFactory.initializeFactory(OICash.address);

  await deployer.deploy(RepExchange);
  let _repExchangeFactory = await deployer.deploy(RepExchangeFactory);
  _repExchangeFactory.initializeFactory(RepExchange.address);

  await deployer.deploy(ReputationToken);
  let _reputationTokenFactory = await deployer.deploy(ReputationTokenFactory);
  _reputationTokenFactory.initializeFactory(ReputationToken.address);

  await deployer.deploy(TestNetReputationToken);
  let _testNetReputationTokenFactory = await deployer.deploy(
    TestNetReputationTokenFactory
  );
  _testNetReputationTokenFactory.initializeFactory(
    TestNetReputationToken.address
  );

  await deployer.deploy(Universe);
  let _universeFactory = await deployer.deploy(UniverseFactory);
  _universeFactory.initializeFactory(Universe.address);

  await deployer.deploy(DisputeWindow);
  let disputeWindowFactory = await deployer.deploy(DisputeWindowFactory);
  disputeWindowFactory.initializeFactory(DisputeWindow.address);

  await deployer.deploy(DisputeCrowdsourcer);
  let disputeCSFactory = await deployer.deploy(DisputeCrowdsourcerFactory);
  disputeCSFactory.initializeFactory(DisputeCrowdsourcer.address);

  // await deployer.deploy(ZeroXTrade);
  // Exception: Can't resolve AVMLabel(jumpdest_784030224795475933405737832577560929931042096197_13374), got AVMLabel(jumpdest_784030224795475933405737832577560929931042096197_13374)
  // await deployer.deploy(Trade); //
  //   Traceback (most recent call last):
  //   File "/usr/local/bin/arbc-truffle", line 7, in <module>
  //     exec(compile(f.read(), __file__, 'exec'))
  //   File "/Users/minhtruong/Dev/arbitrum-pro/packages/arb-compiler-evm/bin/arbc-truffle", line 56, in <module>
  //     main()
  //   File "/Users/minhtruong/Dev/arbitrum-pro/packages/arb-compiler-evm/bin/arbc-truffle", line 41, in main
  //     vm = create_evm_vm(contracts, not args.optimize_off)
  //   File "/Users/minhtruong/Dev/arbitrum-pro/packages/arb-compiler-evm/arbitrum/evm/contract.py", line 63, in create_evm_vm
  //     initial_block, code = generate_evm_code(code, storage)
  //   File "/Users/minhtruong/Dev/arbitrum-pro/packages/arb-compiler-evm/arbitrum/evm/compile.py", line 179, in generate_evm_code
  //     contract,
  //   File "/Users/minhtruong/Dev/arbitrum-pro/packages/arb-compiler-evm/arbitrum/evm/compile.py", line 391, in generate_contract_code
  //     code = replace_self_balance(code)
  //   File "/Users/minhtruong/Dev/arbitrum-pro/packages/arb-compiler-evm/arbitrum/evm/compile.py", line 46, in replace_self_balance
  //     and instrs[i + 1].name == "PUSH20"
  // IndexError: list index out of range
  // await deployer.deploy(EthExchange);
  // Exception: Can't resolve AVMLabel(jumpdest_784030224795475933405737832577560929931042096197_6304), got AVMLabel(jumpdest_784030224795475933405737832577560929931042096197_6304)
};
