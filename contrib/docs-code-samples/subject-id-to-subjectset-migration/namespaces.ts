// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { Context, Namespace } from "@ory/keto-namespace-types"

class User implements Namespace {}
class ApiKey implements Namespace {}

class File implements Namespace {
  related: {
    viewer: (User | ApiKey)[]
  }

  permits = {
    view: (ctx: Context): boolean => this.related.viewer.includes(ctx.subject),
  }
}
