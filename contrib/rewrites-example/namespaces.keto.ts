// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { Namespace, SubjectSet, Context } from "@ory/keto-namespace-types"

class User implements Namespace {
  related: {
    manager: User[]
  }
}

class Group implements Namespace {
  related: {
    members: (User | Group)[]
  }
}

class Folder implements Namespace {
  related: {
    parents: (File | Folder)[]
    viewers: SubjectSet<Group, "members">[]
  }

  permits = {
    view: (ctx: Context): boolean =>
      this.related.viewers.includes(ctx.subject) ||
      this.related.parents.traverse((p) => p.permits.view(ctx)),
  }
}

class File implements Namespace {
  related: {
    parents: (File | Folder)[]
    viewers: (User | SubjectSet<Group, "members">)[]
    owners: (User | SubjectSet<Group, "members">)[]
  }

  // Some comment
  permits = {
    view: (ctx: Context): boolean =>
      this.related.parents.traverse((p) => p.permits.view(ctx)) ||
      this.related.viewers.includes(ctx.subject) ||
      this.related.owners.includes(ctx.subject),

    edit: (ctx: Context) => this.related.owners.includes(ctx.subject),
  }
}
