import { ethers } from "hardhat"
import { EthERC20Bridge__factory } from "arb-ts/src/lib/abi/factories/EthERC20Bridge__factory"
import { StandardArbERC20__factory } from "arb-ts/src/lib/abi/factories/StandardArbERC20__factory"

const main = async () => {
    const l1TokenAddress = process.env.L1_TOKEN_ADDR || "0xFe45D1134cA3ACd42E0a25905BD1cf5c295d8B09"
    const tokenBridgeAddr = process.env.L1_TOKEN_BRIDGE_ADDR || "0x6d48782028e460a17Bc0ceD652e7cB3649d28881"
    
    const tokenBridge = EthERC20Bridge__factory.connect(tokenBridgeAddr, ethers.provider)

    console.log(await tokenBridge.l2Buddy())


    const l1Token = StandardArbERC20__factory.connect(l1TokenAddress, ethers.provider)
    console.log(await l1Token.decimals())

    // console.log(`balanceOf("")`)
    // console.log(await l1Token.balanceOf(""))

    /*
        const accounts = await ethers.getSigners()
        console.log("balance")
        console.log(
            await ethers.provider.getBalance(
                accounts[0].address
            )
        )
    */

    const isErc20 = true
    const maxSubmissionCost = 0
    const gasPrice = 0
    const maxGas = 100000000000

    console.log("sending update tx")
    const updateDecimals = await tokenBridge.updateTokenInfo(l1TokenAddress, isErc20, maxSubmissionCost, maxGas, gasPrice)
    console.log(updateDecimals)

    const receipt = await updateDecimals.wait()
    console.log(receipt)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })


