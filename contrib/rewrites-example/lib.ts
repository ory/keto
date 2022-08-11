/// <reference no-default-lib="true"/>

type Context = { subject: never }

interface Namespace {
  related?: { [relation: string]: Namespace[] }
  permits?: { [method: string]: (ctx: Context) => boolean }
}

interface Array<Namespace> {
  includes(element: Namespace): boolean
  traverse(iteratorfn: (element: Namespace) => boolean): boolean
}

type SubjectSet<
  A extends Namespace,
  R extends keyof A["related"],
> = A["related"][R] extends Array<infer T> ? T : never
