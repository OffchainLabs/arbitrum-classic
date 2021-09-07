import yargs from 'yargs/yargs'

const argv = yargs(process.argv.slice(2))
  .options({
    address: {
      type: 'string',
    },
    txid: {
      type: 'string',
    },
    networkID: {
      type: 'number',
    },
  })
  .parseSync()

export default argv
