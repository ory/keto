/// <reference no-default-lib="true"/>

declare interface Boolean {}
declare interface String {}
declare interface Number {}
declare interface Function {}
declare interface Object {}
declare interface IArguments {}
declare interface RegExp {}

declare interface Array<T extends namespace> {
  includes(element: T): boolean
  traverse(iteratorfn: (element: T) => boolean): boolean
}

interface context {
  subject: never
}

interface namespace {
  related?: { [relation: string]: namespace[] }
  permits?: { [method: string]: (ctx: context) => boolean }
}

declare module "@ory/keto-namespace-types" {
  export type Context = context

  export type Namespace = namespace

  export type SubjectSet<
    A extends Namespace,
    R extends keyof A["related"],
  > = A["related"][R] extends Array<infer T> ? T : never
}
