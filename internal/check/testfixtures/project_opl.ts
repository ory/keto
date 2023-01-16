// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import {Namespace, Context} from '@ory/keto-namespace-types'

class User implements Namespace {}

class Project implements Namespace {
  related: {
    owner: User[]
    developer: User[]
  }

  permits = {
    isOwner: (ctx: Context) => this.related.owner.includes(ctx.subject),
    isOwnerOrDeveloper: (ctx: Context) =>
      this.related.owner.includes(ctx.subject) ||
      this.related.developer.includes(ctx.subject),
    writeCollaborator: (ctx: Context) =>
      this.permits.isOwner(ctx),
    readCollaborator: (ctx: Context) =>
      this.permits.isOwnerOrDeveloper(ctx),
    deleteProject: (ctx: Context) => this.permits.isOwner(ctx),
    writeProject: (ctx: Context) =>
      this.permits.isOwnerOrDeveloper(ctx),
    readProject: (ctx: Context) =>
      this.permits.isOwnerOrDeveloper(ctx),
  }
}
