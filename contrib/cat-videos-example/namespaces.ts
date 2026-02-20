// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { Context, Namespace, SubjectSet } from "@ory/keto-namespace-types"

class User implements Namespace {}

class Video implements Namespace {
  related: {
    view: (User | SubjectSet<Video, "owner">)[]
    owner: User[]
  }
}
