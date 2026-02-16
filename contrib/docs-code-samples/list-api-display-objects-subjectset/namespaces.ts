// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { Context, Namespace } from "@ory/keto-namespace-types"

class User implements Namespace {}

class Chat implements Namespace {
  related: {
    member: User[]
  }
}
