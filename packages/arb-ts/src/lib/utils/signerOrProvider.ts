import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'

export type SignerOrProvider = Signer | Provider

/**
 * Utiliy functions for signer/provider union types
 */
export class SignerOrProviderUtils {
  public static isSigner(
    signerOrProvider: SignerOrProvider
  ): signerOrProvider is Signer {
    return (signerOrProvider as Signer).sendTransaction !== undefined
  }

  public static getProvider(signerOrProvider: SignerOrProvider) {
    return this.isSigner(signerOrProvider)
      ? signerOrProvider.provider
      : signerOrProvider
  }
}
