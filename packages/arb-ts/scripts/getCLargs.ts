import yargs from 'yargs/yargs'

const argv = yargs(process.argv.slice(2)).options({
  address: {
    type: 'string',
  },
  l1Address: {
    type: 'string',
  },
  l2Address: {
    type: 'string',
  },
  txid: {
    type: 'string',
  },
}).argv

export default argv
