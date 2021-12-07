// we don't expose this method in shared lib so we don't risk breaking this
// method because of a dependency needed in lib
export const getPackagePath = (packageName: string) => {
  const path = require.resolve(`${packageName}/package.json`)
  return path.substr(0, path.indexOf('package.json'))
}
