import { Context, Namespace } from "@ory/keto-namespace-types"

class User implements Namespace {
  related: {}
}

class Video implements Namespace {
  related: {
    owners: User[]
    parents: Video[]
  }

  permits = {
    view: (ctx: Context) => this.permits.isOwner(ctx),
    isOwner: (ctx: Context) =>
      this.related.owners.includes(ctx.subject) ||
      this.related.parents.traverse((p) => p.permits.isOwner(ctx)),
  }
}
