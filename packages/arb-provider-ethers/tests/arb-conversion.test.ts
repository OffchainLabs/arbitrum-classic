/* eslint-env node, jest */
import { ArbConversion } from '../src/lib/conversion'
import { bigNumberify } from 'ethers/utils'

const defaultVals = {
  ticksPerBlock: bigNumberify(1000),
  secondsPerBlock: bigNumberify(13),
  gasPerSecond: 10 ** 8,
  gasPerStep: 5,
}

let defaultConverter: ArbConversion
beforeAll(() => {
  defaultConverter = new ArbConversion()
})

describe('ArbConversion', () => {
  it('uses expected default values with no constructor args', () => {
    expect(defaultConverter.ticksPerBlock).toEqual(defaultVals.ticksPerBlock)
    expect(defaultConverter.secondsPerBlock).toEqual(
      defaultVals.secondsPerBlock
    )
    expect(defaultConverter.gasPerSecond).toEqual(defaultVals.gasPerSecond)
    expect(defaultConverter.gasPerStep).toEqual(defaultVals.gasPerStep)
  })

  it('uses passed in args to execute conversion', () => {
    const tpb = bigNumberify(500)
    const spb = bigNumberify(6)
    const gpsp = 10 ** 6
    const gpst = 3

    const converter = new ArbConversion(tpb, spb, gpsp, gpst)
    expect(converter.ticksPerBlock).toEqual(tpb)
    expect(converter.secondsPerBlock).toEqual(spb)
    expect(converter.gasPerSecond).toEqual(gpsp)
    expect(converter.gasPerStep).toEqual(gpst)
  })

  it('converts correctly', () => {
    const rand = Math.floor(Math.random() * 1000)
    const randSpeed = Math.floor(Math.random() * 20)
    const randBN = bigNumberify(rand)
    expect(defaultConverter.blocksToSeconds(rand)).toEqual(
      defaultVals.secondsPerBlock.mul(rand)
    )
    expect(defaultConverter.blocksToTicks(rand)).toEqual(
      defaultVals.ticksPerBlock.mul(rand)
    )
    expect(defaultConverter.ticksToBlocks(randBN)).toEqual(
      randBN.div(defaultVals.ticksPerBlock)
    )
    expect(defaultConverter.ticksToSeconds(randBN)).toEqual(
      randBN.div(defaultVals.ticksPerBlock).mul(defaultVals.secondsPerBlock)
    )
    expect(defaultConverter.secondsToBlocks(randBN)).toEqual(
      randBN.div(defaultVals.secondsPerBlock)
    )
    expect(defaultConverter.secondsToTicks(randBN)).toEqual(
      randBN.div(defaultVals.secondsPerBlock).mul(defaultVals.ticksPerBlock)
    )
    expect(defaultConverter.cpuFactorToSpeedLimitSecs(rand)).toEqual(
      rand * defaultVals.gasPerSecond
    )
    expect(defaultConverter.speedLimitSecsToCpuFactor(rand)).toEqual(
      rand / defaultVals.gasPerSecond
    )
    expect(defaultConverter.assertionTimeToSteps(rand, randSpeed)).toEqual(
      (rand * randSpeed) / defaultVals.gasPerStep
    )
    expect(defaultConverter.stepsToAssertionTime(rand, randSpeed)).toEqual(
      (rand * defaultVals.gasPerStep) / randSpeed
    )
  })
})
