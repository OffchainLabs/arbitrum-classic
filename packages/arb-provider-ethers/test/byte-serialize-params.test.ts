import { serializeParams } from '../src/lib/byte_serialize_params'
import { expect } from 'chai'

describe('serializeParams', () => {
  it('returns bytes array of expected length', async () => {
    const res = await serializeParams([
      true,
      '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
      [
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
      ],
      '1',
      '2',
    ])
    expect(res.length).to.equal(128)
  })

  it('returns bytes array of expected length when addresses are indexed', async () => {
    const res = await serializeParams(
      [
        true,
        '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        [
          '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
          '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
        ],
        '1',
        '2',
      ],
      () => new Promise(exec => exec(42))
    )
    expect(res.length).to.equal(80)
  })

  it('returns bytes array of expected length when only 1 address is indexed', async () => {
    const res = await serializeParams(
      [
        [
          '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
          '0x7363EB6D7ebFB0EcbF80F9d15688CfBf8D7EF191',
        ],
      ],
      str => {
        return new Promise(exec =>
          exec(str === '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F' ? 1 : -1)
        )
      }
    )
    expect(res.join('')).to.equal(
      '2056411571161611052301412442181332511819825334361062211591159923510912619117623619112824920986136207191141126241145'
    )

    expect(res.length).to.equal(42)
  })
})
