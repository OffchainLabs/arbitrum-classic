const ArbBalanceTracker = artifacts.require("ArbBalanceTracker");
const TestToken = artifacts.require("TestToken");
const TestItem = artifacts.require("TestItem");
let ethAddress = "0x0000000000000000000000000000000000000000";

contract("ArbBalanceTracker", accounts => {
  it("should put 1000 wei in the first account", async () => {
    let tracker = await ArbBalanceTracker.new();
    let balance = await tracker.getTokenBalance(ethAddress, accounts[0]);
    assert.equal(balance, 0, "Eth balance should start at 0");
    await tracker.depositEth(accounts[0], { value: 1000 });
    balance = await tracker.getTokenBalance(ethAddress, accounts[0]);
    assert.equal(balance, 1000, "Eth balance wasn't deposited successfully");
  });

  it("should put 1000 ERC20 tokens in the first account", async () => {
    let tracker = await ArbBalanceTracker.new();
    let token = await TestToken.new(10000, { from: accounts[0] });
    await token.approve(tracker.address, 100);
    let balance = await tracker.getTokenBalance(token.address, accounts[0]);
    assert.equal(balance, 0, "ERC20 balance should start at 0");
    await tracker.depositERC20(token.address, 50);
    balance = await tracker.getTokenBalance(token.address, accounts[0]);
    assert.equal(balance, 50, "ERC20 Balance wasn't deposited successfully");
  });

  it("should put ERC721 tokens in the first account", async () => {
    let tracker = await ArbBalanceTracker.new();
    let token = await TestItem.new({ from: accounts[0] });
    await token.mintItem(accounts[0], 1234);
    let hasFunds = await tracker.hasNFT(token.address, accounts[0], 1234);
    assert.isFalse(hasFunds, "ERC721 Balance should start 0");
    await token.safeTransferFrom(accounts[0], tracker.address, 1234, {
      from: accounts[0]
    });
    hasFunds = await tracker.hasNFT(token.address, accounts[0], 1234);
    assert.isTrue(hasFunds, "ERC721 Balance wasn't deposited successfully");
  });

  it("should put correctly calculate hasFunds", async () => {
    let tracker = await ArbBalanceTracker.new();
    let erc20token = await TestToken.new(10000, { from: accounts[0] });
    let erc721token = await TestItem.new({ from: accounts[0] });
    await erc20token.approve(tracker.address, 100);
    await erc721token.mintItem(accounts[0], 1234);

    let tokenTypes = [
      erc20token.address + "00",
      ethAddress + "00",
      erc721token.address + "01"
    ];

    await tracker.depositERC20(erc20token.address, 50);
    await tracker.depositEth(accounts[0], { value: 1000 });
    await erc721token.safeTransferFrom(accounts[0], tracker.address, 1234, {
      from: accounts[0]
    });

    hasFunds = await tracker.hasFunds(accounts[0], [tokenTypes[0]], [20]);
    assert.isTrue(hasFunds, "Account should have sufficient erc20");

    hasFunds = await tracker.hasFunds(accounts[0], [tokenTypes[1]], [40]);
    assert.isTrue(hasFunds, "Account should have sufficient eth");

    hasFunds = await tracker.hasFunds(accounts[0], [tokenTypes[2]], [1234]);
    assert.isTrue(hasFunds, "Account should have correct erc721");

    hasFunds = await tracker.hasFunds(accounts[0], [tokenTypes[2]], [1230]);
    assert.isFalse(hasFunds, "Account should not have incorrect erc721");

    hasFunds = await tracker.hasFunds(accounts[0], tokenTypes, [
      50,
      1000,
      1234
    ]);
    assert.isTrue(hasFunds, "Account should have sufficient balance");

    hasFunds = await tracker.hasFunds(accounts[0], tokenTypes, [10, 10, 1234]);
    assert.isTrue(hasFunds, "Account should have sufficient balance");

    hasFunds = await tracker.hasFunds(accounts[0], tokenTypes, [200, 50, 1234]);
    assert.isFalse(hasFunds, "Account should not have sufficient balance");

    hasFunds = await tracker.hasFunds(accounts[0], tokenTypes, [
      20,
      1100,
      1234
    ]);
    assert.isFalse(hasFunds, "Account should not have sufficient balance");

    hasFunds = await tracker.hasFunds(accounts[0], tokenTypes, [20, 40, 1230]);
    assert.isFalse(hasFunds, "Account should not have sufficient balance");
  });
});
