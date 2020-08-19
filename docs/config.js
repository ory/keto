module.exports = {
  projectName: 'ORY Keto',
  projectSlug: 'keto',
  projectTagLine:
    'A cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.',
  updateTags: [
    {
      image: 'oryd/keto',
      files: ['docs/docs/configure-deploy.md']
    },
    {
      replacer: ({ content, next }) =>
        content.replace(
          /v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?/gi,
          `${next}`
        ),
      files: ['docs/docs/install.md']
    }
  ],
  updateConfig: {
    src: '.schema/config.schema.json',
    dst: './docs/docs/reference/configuration.md'
  }
}
