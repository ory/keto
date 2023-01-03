// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { Namespace, SubjectSet, Context } from "@ory/keto-namespace-types"

// This test is not really a valid config, but rather a check of the types.
class User implements Namespace {
  related: {
    friends: User[]
  }
}

class Group implements Namespace {
  related: {
    members: (User | Group)[]
  }

  permits = {
    isMember: (ctx: Context): boolean =>
      this.related.members.traverse((m) =>
        m instanceof User ? m == ctx.subject : m.permits.isMember(ctx),
      ),
  }
}

class File implements Namespace {
  related: {
    viewers: (User | SubjectSet<Group, "members">)[]
  }

  permits = {
    view: (ctx: Context): boolean =>
      this.related.viewers.traverse((p) =>
        p instanceof User
          ? p.related.friends.includes(ctx.subject)
          : p.permits.isMember(ctx),
      ) || this.related.viewers.includes(ctx.subject),
  }
}
