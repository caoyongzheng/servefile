function getPaths(path = '') {
  const splitpaths = path.split('/')
  const paths = []
  for (let i = 0; i < splitpaths.length; i++) {
    paths.push('children', splitpaths[i])
  }
  return paths.filter(s => s)
}

export default {
  getPaths
}
