module.exports = {
  entryPoints: ['./src/lib'],
  out: 'docs',
  exclude: ['./src/lib/abi'],
  toc: ['L2Bridge'], // TODO: noop, why
  excludeNotDocumented: true,
}
