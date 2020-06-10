import * as Hashing from '../src/lib/hashing'
import * as ethers from 'ethers'

//helper:
const randHexString = () => '0x' + Math.random().toString().substring(2, 15)

describe('calculateTransactionHash tests', function () {
  test('works with precalucated hash', function () {
    expect(
      Hashing.calculateTransactionHash(
        '0x1',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        '0x755449b9901f91deC52DB39AF8c655206C63eD8e',
        ethers.utils.bigNumberify(5),
        ethers.utils.bigNumberify(10),
        '0xabc'
      )
    ).toBe('0xffbd8a9fedd19f018642a8bb711a91a1b49b1774bf8af75132fba1af7d94b7c4')
  })

  test('throws with non-hex string data', function () {
    expect(() => {
      Hashing.calculateTransactionHash(
        '1',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        '0x755449b9901f91deC52DB39AF8c655206C63eD8e',
        ethers.utils.bigNumberify(5),
        ethers.utils.bigNumberify(10),
        '0xabc'
      )
    }).toThrow()

    expect(() => {
      Hashing.calculateTransactionHash(
        '0x1',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        '0x755449b9901f91deC52DB39AF8c655206C63eD8e',
        ethers.utils.bigNumberify(5),
        ethers.utils.bigNumberify(10),
        'abc'
      )
    }).toThrow()
  })

  test('random inputs output 32 bytes of data', function () {
    expect(
      ethers.utils.arrayify(
        Hashing.calculateTransactionHash(
          randHexString(),
          '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
          '0x755449b9901f91deC52DB39AF8c655206C63eD8e',
          ethers.utils.bigNumberify(Math.round(Math.random() * 10)),
          ethers.utils.bigNumberify(Math.round(Math.random() * 10)),
          randHexString()
        )
      ).length
    ).toBe(32)
  })
})

describe('calculateBatchTransactionHash tests', function () {
  test('works with precalucated hash', function () {
    expect(
      Hashing.calculateBatchTransactionHash(
        '0x1',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        ethers.utils.bigNumberify(5),
        ethers.utils.bigNumberify(10),
        '0xabc'
      )
    ).toBe('0x935e9ab43661b8911d589b4f7bec0f7638d15c798ff6b6346bab20b7549eb9b8')
  })
  test('throws with non-hex string data', function () {
    expect(() => {
      Hashing.calculateBatchTransactionHash(
        '1',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        ethers.utils.bigNumberify(5),
        ethers.utils.bigNumberify(10),
        '0xabc'
      )
    }).toThrow()

    expect(() => {
      Hashing.calculateBatchTransactionHash(
        '0x1',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        ethers.utils.bigNumberify(5),
        ethers.utils.bigNumberify(10),
        'abc'
      )
    }).toThrow()
  })

  test('random inputs output 32 bytes of data', function () {
    expect(
      ethers.utils.arrayify(
        Hashing.calculateBatchTransactionHash(
          randHexString(),
          '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
          ethers.utils.bigNumberify(Math.round(Math.random() * 10)),
          ethers.utils.bigNumberify(Math.round(Math.random() * 10)),
          randHexString()
        )
      ).length
    ).toBe(32)
  })
})
