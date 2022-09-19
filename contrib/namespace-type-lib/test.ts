import { Namespace, SubjectSet, Context } from '@ory/keto-namespace-types'

class User implements Namespace {
  related: {
    friends: User[]
  }
}

class Group implements Namespace {
  related: {
    members: (User | Group)[]
  }
}

class File implements Namespace {
  related: {
    viewers: (User | SubjectSet<Group, "members">)[]
  }

  // Some comment
  permits = {
    view: (ctx: Context): boolean =>
      this.related.viewers.traverse((p) => (p instanceof User ? p.related.friends.includes(ctx.subject) : p.related.members.traverse())) ||
      this.related.viewers.includes(ctx.subject)
  }
}
