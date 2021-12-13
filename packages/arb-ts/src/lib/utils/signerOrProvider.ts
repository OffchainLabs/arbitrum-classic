import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'

/**
 * Utiliy functions for signer/provider union types
 */
export class SignerOrProvider {
  public static isSigner(
    signerOrProvider: Provider | Signer
  ): signerOrProvider is Signer {
    return (signerOrProvider as Signer).sendTransaction !== undefined
  }

  public static getProvider(signerOrProvider: Provider | Signer) {
    return this.isSigner(signerOrProvider)
      ? signerOrProvider.provider
      : signerOrProvider
  }
}
