const ArbValueJS = require("arb-provider-ethers").ArbValue;
const ArbValue = artifacts.require("ArbValue");
const ethers = require("ethers");

let testVal =
  "0x5345325345325345325345325345325345325345325345325345325345325435";

contract("ArbValue", accounts => {
  it("should properly calculate bytestack hash 32 bytes", async () => {
    let arbValue = await ArbValue.deployed();
    let ethVal = await arbValue.bytesToBytestackHash(testVal);
    let jsVal = ArbValueJS.hexToBytestack(testVal).hash();
    assert.equal(ethVal, jsVal);
  });

  it("should properly calculate bytestack hash 64 bytes", async () => {
    let arbValue = await ArbValue.deployed();
    let ethVal = await arbValue.bytesToBytestackHash(
      testVal + testVal.slice(2)
    );
    let jsVal = ArbValueJS.hexToBytestack(testVal + testVal.slice(2)).hash();
    assert.equal(ethVal, jsVal);
  });

  it("should properly calculate bytestack hash 16 bytes", async () => {
    let arbValue = await ArbValue.deployed();
    let ethVal = await arbValue.bytesToBytestackHash(testVal.slice(0, 34));
    let jsVal = ArbValueJS.hexToBytestack(testVal.slice(0, 34)).hash();
    assert.equal(ethVal, jsVal);
  });

  it("should properly calculate bytestack hash 19 bytes", async () => {
    let arbValue = await ArbValue.deployed();
    let ethVal = await arbValue.bytesToBytestackHash(testVal.slice(0, 40));
    let jsVal = ArbValueJS.hexToBytestack(testVal.slice(0, 40)).hash();
    assert.equal(ethVal, jsVal);
  });

  it("should properly convert bytestack to bytes", async () => {
    let bytestack = ArbValueJS.hexToBytestack(testVal.slice(0, 40));
    let bytestackData = ethers.utils.hexlify(ArbValueJS.marshal(bytestack));
    let arbValue = await ArbValue.deployed();
    let ethVal = await arbValue.bytestackToBytes(bytestackData);
    assert.equal(ethVal, testVal.slice(0, 40));
  });
});
