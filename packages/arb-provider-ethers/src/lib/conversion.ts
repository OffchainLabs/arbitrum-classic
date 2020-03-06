import { BigNumber, bigNumberify, BigNumberish } from 'ethers/utils';

// TODO async generator that pulls constants from contracts
export class ArbConversion {
    constructor(
        private readonly ticksPerBlock: BigNumber = bigNumberify(1000),
        private readonly secondsPerBlock: BigNumber = bigNumberify(13),
        private readonly gasPerSecond: number = 10 ** 8,
        private readonly gasPerStep: number = 5,
    ) {}

    blocksToSeconds(blocks: BigNumberish): BigNumber {
        return this.secondsPerBlock.mul(blocks);
    }

    blocksToTicks(blocks: BigNumberish): BigNumber {
        return this.ticksPerBlock.mul(blocks);
    }

    ticksToBlocks(ticks: BigNumber): BigNumber {
        return ticks.div(this.ticksPerBlock);
    }

    ticksToSeconds(ticks: BigNumber): BigNumber {
        return this.blocksToSeconds(this.ticksToBlocks(ticks));
    }

    secondsToBlocks(seconds: BigNumberish): BigNumber {
        return bigNumberify(seconds).div(this.secondsPerBlock);
    }

    secondsToTicks(seconds: BigNumberish): BigNumber {
        return this.blocksToTicks(this.secondsToBlocks(seconds));
    }

    cpuFactorToSpeedLimit(factor: number): number {
        return factor * this.gasPerSecond;
    }

    assertionTimeToSteps(seconds: number, speedLimitSeconds: number): number {
        return (seconds * speedLimitSeconds) / this.gasPerStep;
    }
}
