import yargs from 'yargs/yargs'

const argv = yargs(process.argv.slice(2)).options({
  address: {
    type: 'string',
  },
}).argv

export default argv
