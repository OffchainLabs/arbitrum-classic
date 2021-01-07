import { ethers } from 'hardhat'
import { Signer } from 'ethers'

export async function initializeAccounts(): Promise<Signer[]> {
  const [account0] = await ethers.getSigners()
  const provider = account0.provider!

  const accounts: Signer[] = [account0]
  for (let i = 0; i < 9; i++) {
    const account = ethers.Wallet.createRandom().connect(provider)
    accounts.push(account)
    const tx = await account0.sendTransaction({
      value: ethers.utils.parseEther('1.0'),
      to: await account.getAddress(),
    })
    await tx.wait()
  }
  return accounts
}
