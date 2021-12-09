import fs from 'fs'
/* eslint-disable  @typescript-eslint/no-var-requires */
const args = require('args-parser')(process.argv)

interface GasMeasure {
  contract: string
  name: string
  min: number | string
  minPercent: number | string
  max: number | string
  maxPercent: number | string
  average: number | string
  averagePercent: number | string
  numberOfCalls: number | string
  numberOfCallsPercent: number | string
  key: string
}

interface Deployment {
  name: string
  bytecode: string
  deployedBytecode: string
  gasData: number[]
}

/**
 * Calculates diffs from hardhat-gas-reporter output json files
 */
class GasDiffReporter {
  private readonly differences: GasMeasure[] = []
  /* eslint-disable  @typescript-eslint/no-explicit-any */
  private readonly gasReport1: any
  /* eslint-disable  @typescript-eslint/no-explicit-any */
  private readonly gasReport2: any

  constructor(
    public readonly gasReportFileLocation1: string,
    public readonly gasReportFileLocation2: string
  ) {
    this.gasReport1 = this.loadJson(gasReportFileLocation1)
    this.gasReport2 = this.loadJson(gasReportFileLocation2)
  }

  private loadJson(fileName: string) {
    const rawdata = fs.readFileSync(fileName)
    return JSON.parse(rawdata.toString())
  }

  private toMinMaxAvg(data: number[]) {
    const min = Math.min(...data)
    const max = Math.max(...data)
    const avg = Math.round(data.reduce((a, b) => a + b, 0) / data.length)

    return {
      min,
      max,
      avg,
    }
  }

  private getPercentOf(a: number, b: number) {
    return Math.round(((a - b) / b) * 10000) / 100
  }

  /**
   * Calculate gas usage diffs of contract function calls
   */
  public calcFunctionDiffs() {
    const gasReport1Methods = this.gasReport1['info']['methods']
    const gasReport2Methods = this.gasReport2['info']['methods']

    const validEntries1 = Object.keys(gasReport1Methods).filter(
      m => gasReport1Methods[m]['gasData'].length > 0
    )
    const validEntries2 = Object.keys(gasReport2Methods).filter(
      m => gasReport2Methods[m]['gasData'].length > 0
    )
    const inBoth = validEntries1.filter(v1 => validEntries2.includes(v1))
    const in1 = validEntries1.filter(v1 => !validEntries2.includes(v1))
    const in2 = validEntries2.filter(v1 => !validEntries1.includes(v1))

    for (const methodKey of inBoth) {
      const method1 = gasReport1Methods[methodKey]
      const method2 = gasReport2Methods[methodKey]

      const gasData1 = this.toMinMaxAvg(method1['gasData'])
      const gasData2 = this.toMinMaxAvg(method2['gasData'])

      const measure: GasMeasure = {
        key: methodKey,
        contract: method1['contract'],
        name: method1['method'],
        numberOfCalls: method2['numberOfCalls'] - method1['numberOfCalls'],
        numberOfCallsPercent: this.getPercentOf(
          method2['numberOfCalls'],
          method1['numberOfCalls']
        ),
        min: gasData2.min - gasData1.min,
        minPercent: this.getPercentOf(gasData2.min, gasData1.min),
        max: gasData2.max - gasData1.max,
        maxPercent: this.getPercentOf(gasData2.max, gasData1.max),
        average: gasData2.avg - gasData1.avg,
        averagePercent: this.getPercentOf(gasData2.avg, gasData1.avg),
      }

      this.differences.push(measure)
    }
    in1.forEach(i => {
      const method = gasReport1Methods[i]
      this.differences.push({
        key: i,
        contract: method['contract'],
        name: method['method'],
        numberOfCalls: 'MISSING_AFTER',
        numberOfCallsPercent: 'MISSING_AFTER',
        min: 'MISSING_AFTER',
        minPercent: 'MISSING_AFTER',
        max: 'MISSING_AFTER',
        maxPercent: 'MISSING_AFTER',
        average: 'MISSING_AFTER',
        averagePercent: 'MISSING_AFTER',
      })
    })
    in2.forEach(i => {
      const method = gasReport2Methods[i]
      this.differences.push({
        key: i,
        contract: method['contract'],
        name: method['method'],
        numberOfCalls: 'MISSING_BEFORE',
        numberOfCallsPercent: 'MISSING_BEFORE',
        min: 'MISSING_BEFORE',
        minPercent: 'MISSING_BEFORE',
        max: 'MISSING_BEFORE',
        maxPercent: 'MISSING_BEFORE',
        average: 'MISSING_BEFORE',
        averagePercent: 'MISSING_BEFORE',
      })
    })

    this.sortDifferences()
  }

  /**
   * Calculate gas usages diffs of contract deployments
   */
  public calcDeploymentDiffs() {
    const gasReport1Deployments = (
      this.gasReport1['info']['deployments'] as Deployment[]
    ).filter(d => d['gasData'].length > 0)
    const gasReport2Deployments = (
      this.gasReport2['info']['deployments'] as Deployment[]
    ).filter(d => d['gasData'].length > 0)
    const inBothDeploys = gasReport1Deployments.filter(
      d1 => gasReport2Deployments.filter(d2 => d2.name == d1.name).length > 0
    )
    const in1Deploys = gasReport1Deployments.filter(
      d1 => gasReport2Deployments.filter(d2 => d2.name == d1.name).length === 0
    )
    const in2Deploys = gasReport2Deployments.filter(
      d1 => gasReport1Deployments.filter(d2 => d2.name == d1.name).length === 0
    )

    for (const deployment1 of inBothDeploys) {
      const deployment2 = gasReport2Deployments.filter(
        a => a.name === deployment1.name
      )[0]

      const gasData1 = this.toMinMaxAvg(deployment1['gasData'])
      const gasData2 = this.toMinMaxAvg(deployment2['gasData'])

      const measure: GasMeasure = {
        key: deployment1['name'] + '_constructor',
        contract: deployment1['name'],
        name: '_constructor',
        numberOfCalls:
          deployment2['gasData'].length - deployment1['gasData'].length,
        numberOfCallsPercent: this.getPercentOf(
          deployment2['gasData'].length,
          deployment1['gasData'].length
        ),
        min: gasData2.min - gasData1.min,
        minPercent: this.getPercentOf(gasData2.min, gasData1.min),
        max: gasData2.max - gasData1.max,
        maxPercent: this.getPercentOf(gasData2.max, gasData1.max),
        average: gasData2.avg - gasData1.avg,
        averagePercent: this.getPercentOf(gasData2.avg, gasData1.avg),
      }
      this.differences.push(measure)
    }

    in1Deploys.forEach(i => {
      this.differences.push({
        key: i.name + '_constructor',
        contract: i.name,
        name: '_constructor',
        numberOfCalls: 'MISSING_AFTER',
        numberOfCallsPercent: 'MISSING_AFTER',
        min: 'MISSING_AFTER',
        minPercent: 'MISSING_AFTER',
        max: 'MISSING_AFTER',
        maxPercent: 'MISSING_AFTER',
        average: 'MISSING_AFTER',
        averagePercent: 'MISSING_AFTER',
      })
    })
    in2Deploys.forEach(i => {
      this.differences.push({
        key: i.name + '_constructor',
        contract: i.name,
        name: '_constructor',
        numberOfCalls: 'MISSING_BEFORE',
        numberOfCallsPercent: 'MISSING_BEFORE',
        min: 'MISSING_BEFORE',
        minPercent: 'MISSING_BEFORE',
        max: 'MISSING_BEFORE',
        maxPercent: 'MISSING_BEFORE',
        average: 'MISSING_BEFORE',
        averagePercent: 'MISSING_BEFORE',
      })
    })

    this.sortDifferences()
  }

  private sortDifferences() {
    this.differences.sort((a, b) => {
      const contractCompare = a.contract.localeCompare(b.contract)
      if (contractCompare == 0) return a.name.localeCompare(b.name)
      else return contractCompare
    })
  }

  /**
   * Get only changed measurements
   */
  public onlyDifferent() {
    return this.differences.filter(
      a => a.average != 0 || a.min != 0 || a.max != 0 || a.numberOfCalls != 0
    )
  }

  /**
   * Write the calculated diffs to csv format. Will overwrite existing file.
   * @param outputFileLocation
   */
  public writeDiffsCsv(outputFileLocation: string) {
    // print all the measures to file
    const differences = this.differences
    let data =
      'key,contract,function,numberOfCalls,min,max,average,numberOfCalls%,min%,max%,average%|\n'
    for (const diff of differences) {
      data += `${diff.key},${diff.contract},${diff.name},${diff.numberOfCalls},${diff.min},${diff.max},${diff.average},${diff.numberOfCallsPercent},${diff.minPercent},${diff.maxPercent},${diff.averagePercent}\n`
    }

    fs.writeFileSync(outputFileLocation, data)
  }

  /**
   * Write the calculated diffs to github markdown. Will overwrite existing file.
   * @param outputFileLocation
   */
  public writeDiffsGithubMd(outputFileLocation: string) {
    // print only difference measures to a file
    const differences = this.onlyDifferent()
    let data = ''
    if (differences.length === 0) {
      data += `${differences.length} methods had a different gas cost.`
    } else {
      data += `<details><summary>${differences.length} method${
        differences.length !== 1 ? 's' : ''
      } had a different gas cost.</summary>\\n\\n`
      data +=
        '|key|contract|function|numberOfCalls|min|max|average|numberOfCalls%|min%|max%|average%|\\n'
      data += '|---|---|---|---|---|---|---|---|---|---|---|\\n'
      for (const diff of differences) {
        data += `|${diff.key}|${diff.contract}|${diff.name}|${diff.numberOfCalls}|${diff.min}|${diff.max}|${diff.average}|${diff.numberOfCallsPercent}|${diff.minPercent}|${diff.maxPercent}|${diff.averagePercent}|\\n`
      }
      data += '</details>'
    }
    fs.writeFileSync(outputFileLocation, data)
  }

  /**
   * Write the calculated diff to a console table.
   */
  public writeDiffsToConsole() {
    const differences = this.onlyDifferent()
    if (differences.length == 0) {
      console.log('No difference between reports.')
    } else {
      console.table(differences)
    }
  }
}

const main = async (args: {
  gasReport1?: string
  gasReport2?: string
  outputFile?: string
}) => {
  const gasReport1Name = args.gasReport1
  if (!gasReport1Name)
    throw new Error(
      'Missing gasReport1 argument, should be the path to a gasReporterOutput.json'
    )

  const gasReport2Name = args.gasReport2
  if (!gasReport2Name)
    throw new Error(
      'Missing gasReport2 argument, should be the path to another gasReporterOutput.json'
    )

  const gasDiffReporter = new GasDiffReporter(gasReport1Name, gasReport2Name)
  gasDiffReporter.calcDeploymentDiffs()
  gasDiffReporter.calcFunctionDiffs()

  if (args.outputFile) {
    gasDiffReporter.writeDiffsCsv(args.outputFile + '.csv')
    gasDiffReporter.writeDiffsGithubMd(args.outputFile + '.githubmd')
  } else {
    gasDiffReporter.writeDiffsToConsole()
  }
}
main(args).catch(console.error)
