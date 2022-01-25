import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { ArbTsError, MissingProviderArbTsError } from '../dataEntities/errors'

export type SignerOrProvider = Signer | Provider

/**
 * Utiliy functions for signer/provider union types
 */
export class SignerProviderUtils {
  public static isSigner(
    signerOrProvider: SignerOrProvider
  ): signerOrProvider is Signer {
    return (signerOrProvider as Signer).signMessage !== undefined
  }

  /**
   * If signerOrProvider is a provider then return itself.
   * If signerOrProvider is a signer then return signer.provider
   * @param signerOrProvider
   * @returns
   */
  public static getProvider(
    signerOrProvider: SignerOrProvider
  ): Provider | undefined {
    return this.isSigner(signerOrProvider)
      ? signerOrProvider.provider
      : signerOrProvider
  }

  public static getProviderOrThrow(
    signerOrProvider: SignerOrProvider
  ): Provider {
    const maybeProvider = this.getProvider(signerOrProvider)
    if (!maybeProvider) throw new MissingProviderArbTsError('signerOrProvider')
    return maybeProvider
  }

  /**
   * Check if the signer has a connected provider
   * @param signer
   */
  public static signerHasProvider(
    signer: Signer
  ): signer is Signer & { provider: Provider } {
    return !!signer.provider
  }

  /**
   * Checks that the signer/provider that's provider matches the chain id
   * Throws if not.
   * @param signerOrProvider
   * @param chainId
   */
  public static async checkNetworkMatches(
    signerOrProvider: SignerOrProvider,
    chainId: number
  ): Promise<void> {
    const provider = this.getProvider(signerOrProvider)
    if (!provider) throw new MissingProviderArbTsError('signerOrProvider')

    const providerChainId = (await provider.getNetwork()).chainId
    if (providerChainId !== chainId) {
      throw new ArbTsError(
        `Signer/provider chain id: ${providerChainId} doesn't match provided chain id: ${chainId}.`
      )
    }
  }
}
