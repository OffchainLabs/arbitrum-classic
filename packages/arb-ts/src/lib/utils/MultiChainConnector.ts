import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import networks, { Network } from '../networks'

export interface SignersAndProviders {
  l1Provider?: Provider
  l2Provider?: Provider
  l1Signer?: Signer
  l2Signer?: Signer
}

export class MultiChainConnector {
  l1Provider?: Provider
  l2Provider?: Provider
  l1Signer?: Signer
  l2Signer?: Signer
  constructor(signersAndProviders: SignersAndProviders) {
    const { l1Provider, l2Provider, l1Signer, l2Signer } = signersAndProviders
    // TODO: check l1 and l2 networks reasonably match
    if (l1Signer) {
      this.l1Signer = l1Signer
      if (l1Signer.provider) {
        this.l1Provider = l1Signer.provider
      }
    }

    if (l2Signer) {
      this.l2Signer = l2Signer
      if (l2Signer.provider) {
        this.l2Provider = l2Signer.provider
      }
    }

    if (!this.l1Provider && l1Provider) {
      this.l1Provider = l1Provider
    }

    if (!this.l2Provider && l2Provider) {
      this.l2Provider = l2Provider
    }
  }

  protected requireL2Provider() {
    if (!this.l2Provider) throw new Error('Must have l2 provider')
  }
}
