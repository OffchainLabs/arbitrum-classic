/* eslint-env node, jest */
import { assert, expect } from 'chai'
import { ArbConversion } from '../src/lib/conversion'
import { bigNumberify } from 'ethers/utils'

const defaultVals = {
  ticksPerBlock: bigNumberify(1000),
  secondsPerBlock: bigNumberify(13),
  gasPerSecond: 10 ** 8,
  gasPerStep: 5,
}

let defaultConverter: ArbConversion
before(() => {
  defaultConverter = new ArbConversion()
})

describe('ArbConversion', () => {
  it('uses expected default values with no constructor args', () => {
    assert(defaultConverter.ticksPerBlock.eq(defaultVals.ticksPerBlock))
    assert(defaultConverter.secondsPerBlock.eq(defaultVals.secondsPerBlock))
    expect(defaultConverter.gasPerSecond).to.equal(defaultVals.gasPerSecond)
    expect(defaultConverter.gasPerStep).to.equal(defaultVals.gasPerStep)
  })

  it('uses passed in args to execute conversion', () => {
    const tpb = bigNumberify(500)
    const spb = bigNumberify(6)
    const gpsp = 10 ** 6
    const gpst = 3

    const converter = new ArbConversion(tpb, spb, gpsp, gpst)
    expect(converter.ticksPerBlock).to.equal(tpb)
    expect(converter.secondsPerBlock).to.equal(spb)
    expect(converter.gasPerSecond).to.equal(gpsp)
    expect(converter.gasPerStep).to.equal(gpst)
  })

  it('converts correctly', () => {
    const rand = Math.floor(Math.random() * 1000)
    const randSpeed = Math.floor(Math.random() * 20)
    const randBN = bigNumberify(rand)
    assert(
      defaultConverter
        .blocksToSeconds(rand)
        .eq(defaultVals.secondsPerBlock.mul(rand))
    )
    assert(
      defaultConverter
        .blocksToTicks(rand)
        .eq(defaultVals.ticksPerBlock.mul(rand))
    )
    assert(
      defaultConverter
        .ticksToBlocks(randBN)
        .eq(randBN.div(defaultVals.ticksPerBlock))
    )
    assert(
      defaultConverter
        .ticksToSeconds(randBN)
        .eq(
          randBN.div(defaultVals.ticksPerBlock).mul(defaultVals.secondsPerBlock)
        )
    )
    assert(
      defaultConverter
        .secondsToBlocks(randBN)
        .eq(randBN.div(defaultVals.secondsPerBlock))
    )
    assert(
      defaultConverter
        .secondsToTicks(randBN)
        .eq(
          randBN.div(defaultVals.secondsPerBlock).mul(defaultVals.ticksPerBlock)
        )
    )
    expect(defaultConverter.cpuFactorToSpeedLimitSecs(rand)).to.equal(
      rand * defaultVals.gasPerSecond
    )
    expect(defaultConverter.speedLimitSecsToCpuFactor(rand)).to.equal(
      rand / defaultVals.gasPerSecond
    )
    expect(defaultConverter.assertionTimeToSteps(rand, randSpeed)).to.equal(
      (rand * randSpeed) / defaultVals.gasPerStep
    )
    expect(defaultConverter.stepsToAssertionTime(rand, randSpeed)).to.equal(
      (rand * defaultVals.gasPerStep) / randSpeed
    )
  })
})
