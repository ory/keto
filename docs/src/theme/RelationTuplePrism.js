import Prism from 'prism-react-renderer/prism'

const delimiter = {
  punctuation: /[:#@()]/
}

const namespace = {
  pattern: /[^:#@()\n]+:/,
  inside: {
    ...delimiter,
    namespace: /.*/
  }
}

const object = {
  pattern: /[^:#@()\n]+#/,
  inside: {
    ...delimiter,
    symbol: /.*/
  }
}

const relation = {
  pattern: /[^:#@()\n]+/,
  alias: 'string'
}

const subjectID = {
  pattern: /@[^:#@()\n]+/,
  inside: {
    ...delimiter,
    keyword: /.*/
  }
}

const subjectSet = {
  pattern: /@\(([^:#@()\n]+:)?([^:#@()\n]+)#([^:#@()\n]*)\)/,
  inside: {
    punctuation: /[@()]/,
    namespace,
    object,
    relation
  }
}

Prism.languages['keto-relation-tuples'] = {
  comment: /\/\/.*(\n|$)/,
  'relation-tuple': {
    pattern: /([^:#@()\n]+:)?([^:#@()\n]+)#([^:#@()\n]+)@?((\(([^:#@()\n]+:)?([^:#@()\n]+)#([^:#@()\n]*)\))|([^:#@()\n]+))/,
    inside: {
      namespace,
      object,
      subjectID,
      subjectSet,
      relation
    }
  }
}

export default () => {}
