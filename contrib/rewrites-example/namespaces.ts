import { Namespace, SubjectSet, Context } from '@ory/keto-namespace-types'

// Declare new namespaces as classes that implement `Namespaces`
class User implements Namespace {
  related: {
    // Define relations to other objects here.
    // Examples:
    //
    // parents: (File | Folder)[]
    // viewers: SubjectSet<Group, "members">[]
  }

  permits = {
    // Define permissions here. These can be derived from the relations above.
    // Examples:
    //
    // view: (ctx: Context): boolean =>
    //  this.related.viewers.includes(ctx.subject) ||
    //  this.related.parents.traverse((p) => p.permits.view(ctx)),
  }
}

class Group implements Namespace {
  related: {
    // Define relations to other objects here.
    // Examples:
    //
    // parents: (File | Folder)[]
    // viewers: SubjectSet<Group, "members">[]
  }

  permits = {
    // Define permissions here. These can be derived from the relations above.
    // Examples:
    //
    // view: (ctx: Context): boolean =>
    //  this.related.viewers.includes(ctx.subject) ||
    //  this.related.parents.traverse((p) => p.permits.view(ctx)),
  }
}

class Folder implements Namespace {
  related: {
    // Define relations to other objects here.
    // Examples:
    //
    // parents: (File | Folder)[]
    // viewers: SubjectSet<Group, "members">[]
  }

  permits = {
    // Define permissions here. These can be derived from the relations above.
    // Examples:
    //
    // view: (ctx: Context): boolean =>
    //  this.related.viewers.includes(ctx.subject) ||
    //  this.related.parents.traverse((p) => p.permits.view(ctx)),
  }
}

class File implements Namespace {
  related: {
    // Define relations to other objects here.
    // Examples:
    //
    // parents: (File | Folder)[]
    // viewers: SubjectSet<Group, "members">[]
  }

  permits = {
    // Define permissions here. These can be derived from the relations above.
    // Examples:
    //
    // view: (ctx: Context): boolean =>
    //  this.related.viewers.includes(ctx.subject) ||
    //  this.related.parents.traverse((p) => p.permits.view(ctx)),
  }
}

