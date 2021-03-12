export const addressToSymbol = (erc20L1Address: string) => {
  return erc20L1Address.substr(erc20L1Address.length - 3).toUpperCase() + '?'
}
