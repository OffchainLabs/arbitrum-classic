import fs from 'fs'

// read the gas report

const loadJson = (fileName: string) => {
  const rawdata = fs.readFileSync(fileName)
  return JSON.parse(rawdata.toString())
}

interface GasMeasure {
  contract: string
  name: string
  min: number
  max: number
  average: number
  numberOfCalls: number
}

const toMinMaxAvg = (data: number[]) => {
  const min = Math.min(...data)
  const max = Math.max(...data)
  const avg = Math.round(data.reduce((a, b) => a + b, 0) / data.length)

  return {
    min,
    max,
    avg,
  }
}

const main = async (args: string[]) => {
  const gasReport1Name = args[2]
  if (!gasReport1Name)
    throw new Error(
      'Missing 1st argument, should be the path to a gasReporterOutput.json'
    )

  const gasReport2Name = args[3]
  if (!gasReport2Name)
    throw new Error(
      'Missing 2nd argument, should be the path to another gasReporterOutput.json'
    )

  const output = args[4]
  if (!output)
    throw new Error(
      'Missing 3nd argument, should be the path an output csv file location'
    )

  const differences: GasMeasure[] = []

  const gasReport1 = loadJson(gasReport1Name)
  const gasReport2 = loadJson(gasReport2Name)

  const gasReport1Methods = gasReport1['info']['methods']
  const gasReport2Methods = gasReport2['info']['methods']

  for (const methodKey of Object.keys(gasReport1Methods)) {
    if (
      gasReport1Methods[methodKey]['gasData'].length > 0 &&
      // CHRIS: we should check from both sides and include if either side contains a call
      gasReport2Methods[methodKey] &&
      gasReport2Methods[methodKey]['gasData'].length > 0
    ) {
      const method1 = gasReport1Methods[methodKey]
      const method2 = gasReport2Methods[methodKey]

      const gasData1 = toMinMaxAvg(method1['gasData'])
      const gasData2 = toMinMaxAvg(method2['gasData'])

      const measure: GasMeasure = {
        contract: method1['contract'],
        name: method1['method'],
        numberOfCalls: method2['numberOfCalls'] - method1['numberOfCalls'],
        min: gasData2.min - gasData1.min,
        max: gasData2.max - gasData1.max,
        average: gasData2.avg - gasData1.avg,
      }

      differences.push(measure)
    }
  }

  const gasReport1Deployments = gasReport1['info']['deployments']
  const gasReport2Deployments = gasReport2['info']['deployments']

  for (const methodKey of Object.keys(gasReport1Deployments)) {
    if (
      gasReport1Deployments[methodKey]['gasData'].length > 0 &&
      // CHRIS: we should check from both sides and include if either side contains a call
      gasReport2Deployments[methodKey] &&
      gasReport2Deployments[methodKey]['gasData'].length > 0
    ) {
      const deployment1 = gasReport1Deployments[methodKey]
      const deployment2 = gasReport2Deployments[methodKey]

      const gasData1 = toMinMaxAvg(deployment1['gasData'])
      const gasData2 = toMinMaxAvg(deployment2['gasData'])

      const measure: GasMeasure = {
        contract: deployment1['name'],
        name: '_constructor',
        numberOfCalls:
          deployment2['gasData'].length - deployment1['gasData'].length,
        min: gasData2.min - gasData1.min,
        max: gasData2.max - gasData1.max,
        average: gasData2.avg - gasData1.avg,
      }

      differences.push(measure)
    }
  }

  // print all the measures to file
  let data = 'contract,method,numberOfCalls,min,max,average\n'
  for (const diff of differences.sort((a, b) => {
    const contractCompare = a.contract.localeCompare(b.contract)
    if (contractCompare == 0) return a.name.localeCompare(b.name)
    else return contractCompare
  })) {
    data += `${diff.contract},${diff.name},${diff.numberOfCalls},${diff.min},${diff.max},${diff.average}\n`
  }

  fs.writeFileSync(output, data)
}

main(process.argv).catch(console.error)
