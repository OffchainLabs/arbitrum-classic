import { ethers } from 'hardhat'
import { Signer } from '@ethersproject/abstract-signer'

export async function initializeAccounts(): Promise<Signer[]> {
  const [account0] = await ethers.getSigners()
  const provider = account0.provider!

  const accounts: Signer[] = [account0]
  for (let i = 0; i < 9; i++) {
    const account = ethers.Wallet.createRandom().connect(provider)
    accounts.push(account)
    const tx = await account0.sendTransaction({
      value: ethers.utils.parseEther('10000.0'),
      to: await account.getAddress(),
    })
    await tx.wait()
  }
  return accounts
}

export async function tryAdvanceChain(
  account: Signer,
  blocks: number
): Promise<void> {
  try {
    for (let i = 0; i < blocks; i++) {
      await ethers.provider.send('evm_mine', [])
    }
  } catch (e) {
    // EVM mine failed. Try advancing the chain by sending txes if the node
    // is in dev mode and mints blocks when txes are sent
    for (let i = 0; i < blocks; i++) {
      const tx = await account.sendTransaction({
        value: 0,
        to: await account.getAddress(),
      })
      await tx.wait()
    }
  }
}
