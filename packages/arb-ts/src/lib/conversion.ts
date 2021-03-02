import { BigNumber, BigNumberish } from 'ethers'

// TODO async generator that pulls constants from contracts
// TODO are these values up to date?
export class ArbConversion {
  constructor(
    readonly ticksPerBlock: BigNumber = BigNumber.from(1000),
    readonly secondsPerBlock: BigNumber = BigNumber.from(13),
    readonly gasPerSecond: number = 10 ** 8,
    readonly gasPerStep: number = 5
  ) {}

  blocksToSeconds(blocks: BigNumberish): BigNumber {
    return this.secondsPerBlock.mul(blocks)
  }

  blocksToTicks(blocks: BigNumberish): BigNumber {
    return this.ticksPerBlock.mul(blocks)
  }

  ticksToBlocks(ticks: BigNumber): BigNumber {
    return ticks.div(this.ticksPerBlock)
  }

  ticksToSeconds(ticks: BigNumber): BigNumber {
    return this.blocksToSeconds(this.ticksToBlocks(ticks))
  }

  secondsToBlocks(seconds: BigNumberish): BigNumber {
    return BigNumber.from(seconds).div(this.secondsPerBlock)
  }

  secondsToTicks(seconds: BigNumberish): BigNumber {
    return this.blocksToTicks(this.secondsToBlocks(seconds))
  }

  cpuFactorToSpeedLimitSecs(factor: number): number {
    return factor * this.gasPerSecond
  }

  speedLimitSecsToCpuFactor(seconds: number): number {
    return seconds / this.gasPerSecond
  }

  assertionTimeToSteps(seconds: number, speedLimitSeconds: number): number {
    return (seconds * speedLimitSeconds) / this.gasPerStep
  }

  stepsToAssertionTime(steps: number, speedLimitSeconds: number): number {
    return (steps * this.gasPerStep) / speedLimitSeconds
  }
}
