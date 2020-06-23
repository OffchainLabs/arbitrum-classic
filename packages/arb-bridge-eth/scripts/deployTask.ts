import { task } from '@nomiclabs/buidler/config'
import deploy_contracts from './deployLib'

export default function setupDeployTask() {
  task('deploy', 'Prints the list of accounts', async (taskArgs, bre) => {
    const contracts = await deploy_contracts(bre)
  })
}
