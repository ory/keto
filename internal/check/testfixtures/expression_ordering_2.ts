import { Context, Namespace, SubjectSet } from "@ory/keto-namespace-types"

class User implements Namespace {}

class LegalEntityUserGroup implements Namespace {
  related: {
    // predefined roles
    owners: User[]
    administrators: User[]
    creators: User[]
    viewers: User[]
    members: User[]
  }
}

class BlockedUserGroup implements Namespace {
  related: {
    blockedUsers: User[]
  }
}

class LegalEntity implements Namespace {
  related: {
    // pre-deifned roles
    owners: SubjectSet<LegalEntityUserGroup, "owners">[]
    administrators: SubjectSet<LegalEntityUserGroup, "administrators">[]
    creators: SubjectSet<LegalEntityUserGroup, "creators">[]
    viewers: SubjectSet<LegalEntityUserGroup, "viewers">[]
    members: SubjectSet<LegalEntityUserGroup, "members">[]
    blockedUsers: SubjectSet<BlockedUserGroup, "blockedUsers">[]
  }

  permits = {
    own: (ctx: Context): boolean =>
      !this.related.blockedUsers.includes(ctx.subject) &&
      this.related.owners.includes(ctx.subject),

    itemView: (ctx: Context): boolean =>
      !this.related.blockedUsers.includes(ctx.subject) &&
      (this.related.owners.includes(ctx.subject) || // ⚠️ move this line up and it it won't work.
        this.related.administrators.includes(ctx.subject) ||
        this.related.creators.includes(ctx.subject) ||
        this.related.viewers.includes(ctx.subject) ||
        this.related.members.includes(ctx.subject)),
  }
}
