var Augur = artifacts.require("./Augur.sol");
var AugurTrading = artifacts.require("./AugurTrading.sol");
var WarpSync = artifacts.require("./WarpSync.sol");
// var EthExchange = artifacts.require("./EthExchange.sol");
var GnosisSafeRegistry = artifacts.require("./GnosisSafeRegistry.sol");
var Time = artifacts.require("./Time.sol");

var CreateOrder = artifacts.require("./trading/CreateOrder.sol");
// var Trade = artifacts.require("./trading/Trade.sol");
var Order = artifacts.require("./trading/Order.sol");
var ProfitLoss = artifacts.require("./trading/ProfitLoss.sol");
var SimulateTrade = artifacts.require("./trading/SimulateTrade.sol");
// var ZeroXTrade = artifacts.require("./trading/ZeroXTrade.sol");
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

//--optimize-off

module.exports = async function(deployer) {
  await deployer.deploy(Augur);
  await deployer.deploy(AugurTrading, Augur.address);
  await deployer.deploy(GnosisSafeRegistry);
  await deployer.deploy(WarpSync);
  // await deployer.deploy(EthExchange);

  await deployer.deploy(CreateOrder);
  // await deployer.deploy(Trade);
  await deployer.deploy(Order);
  await deployer.deploy(ProfitLoss);
  await deployer.deploy(SimulateTrade);
  // await deployer.deploy(ZeroXTrade);
  await deployer.deploy(LegacyReputationToken);

  await deployer.deploy(BuyParticipationTokens);
  await deployer.deploy(Formulas);
  await deployer.deploy(HotLoading);
  await deployer.deploy(RedeemStake);
  await deployer.deploy(RepSymbol);

  await deployer.deploy(AffiliateValidator);
  let afiliates = await deployer.deploy(Affiliates);
  afiliates.initializeValidatorFactory(AffiliateValidator.address);

  await deployer.deploy(InitialReporter);
  await deployer.deploy(OICash);
  // await deployer.deploy(ReputationToken); input params
  await deployer.deploy(DisputeCrowdsourcer);
  await deployer.deploy(Market);
  // await deployer.deploy(ShareToken); difficulty

  // await deployer.deploy(DisputeCrowdsourcerFactory);
  // await deployer.deploy(DisputeWindowFactory);
  // await deployer.deploy(InitialReporterFactory);
  // await deployer.deploy(MarketFactory);
  // await deployer.deploy(OICashFactory);
  // await deployer.deploy(RepExchangeFactory);
  // await deployer.deploy(ReputationTokenFactory);
  // await deployer.deploy(TestNetReputationTokenFactory);
  // await deployer.deploy(UniverseFactory);
};
