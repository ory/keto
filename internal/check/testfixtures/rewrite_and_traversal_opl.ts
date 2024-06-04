import { Namespace, Context, SubjectSet } from "@ory/keto-namespace-types"

class User implements Namespace {}

class Role implements Namespace {
    related: {
        members: (User | Role)[]
    }
}

class SuperUsers implements Namespace {
    related: {
        admins: SubjectSet<Role, "members">[]  // annotier admins
    }
}

class Comment implements Namespace {
    related: {
        parents: Group[]
    }

    permits = {
        delete: (ctx: Context): boolean => this.related.parents.traverse((group) => group.permits.update(ctx)),
        update: (ctx: Context): boolean => this.permits.delete(ctx),
    }
}

class Group implements Namespace {
    related: {
        supers: SuperUsers[]
    }

    permits = {
        delete: (ctx: Context): boolean => this.related.supers.traverse((supe) => supe.related.admins.includes(ctx.subject)),  // delete a group
        update: (ctx: Context): boolean => this.permits.delete(ctx), // if can delete then can update
    }
}
