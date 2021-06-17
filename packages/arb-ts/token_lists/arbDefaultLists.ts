import { TokenInfo } from './tokenListTypes'

interface ArbTokenLists {
  [chainId: number]: TokenInfo[]
}

const arbTokenLists: ArbTokenLists = {
  1: [
    {
      address: '0x2ba592f78db6436527729929aaf6c908497cb200',
      chainId: 1,
      decimals: 18,
      logoURI: 'https://zapper.fi/images/CREAM-icon.png',
      name: 'Cream',
      symbol: 'CREAM',
    },
    {
      address: '0x6b3595068778dd592e39a122f4f5a5cf09c90fe2',
      chainId: 1,
      decimals: 18,
      logoURI: 'https://zapper.fi/images/SUSHI-icon.png',
      name: 'Sushi',
      symbol: 'SUSHI',
    },
    {
      address: '0x0a5e677a6a24b2f1a2bf4f3bffc443231d2fdec8',
      chainId: 1,
      decimals: 18,
      logoURI: '',
      name: 'dForce USD',
      symbol: 'USX',
    },
    {
      address: '0xb986f3a2d91d3704dc974a24fb735dcc5e3c1e70',
      chainId: 1,
      decimals: 18,
      logoURI: '',
      name: 'dForce EUR',
      symbol: 'EUX',
    },
  ],
}

export default arbTokenLists
