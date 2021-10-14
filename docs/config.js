const replaceInDir = ({ dir, replacer }) => {
  const { join } = require('path')
  const fs = require('fs')

  const walk = (dir) =>
    fs.readdirSync(dir).flatMap((fn) => {
      const fp = join(dir, fn)
      if (fs.statSync(fp).isDirectory()) {
        return walk(fp)
      }
      return fp
    })

  return new Proxy(
    { replacer },
    {
      get: (target, p) => {
        if (p === 'files') {
          return walk(dir)
        }
        return target[p]
      }
    }
  )
}

module.exports = {
  projectName: 'ORY Keto',
  projectSlug: 'keto',
  projectTagLine:
    'A cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.',
  updateTags: [
    replaceInDir({
      replacer: ({ content, next }) =>
        content.replace(
          /v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?/gi,
          `${next}`
        ),
      dir: 'docs/docs'
    }),
    replaceInDir({
      replacer: ({ content, next }) =>
        content.replace('version="master"', `version="${next}"`),
      dir: 'docs/docs'
    }),
    {
      replacer: ({ content, next }) =>
        content.replace(
          /oryd\/keto:v[0-9a-zA-Z.-]*-sqlite/g,
          `oryd/keto:${next}-sqlite`
        ),
      files: ['contrib/cat-videos-example/docker-compose.yml']
    }
  ],
  updateConfig: {
    src: '../.schema/config.schema.json',
    dst: 'docs/reference/configuration.md'
  }
}
