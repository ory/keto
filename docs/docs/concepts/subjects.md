---
id: subjects
title: Subjects
---

In ORY Keto subjects are a recursive polymorphic datatype. They either refer to
a specific subject (e.g. user) by some application defined identifier, or a set
of subjects ⭮.

## Subject IDs

A subject ID can be any string up to 64 characters. It is up to the application
to map its users, devices, ... to an unambiguous identifier. We recommend the
usage of UUIDs as they provide a high entropy. It is however totally possible to
use usernames, opaque tokens or [strings with special meanings](/TODO) here as
well. ORY Keto will consider subject IDs equal iff their string representation
is equal.

## Subject Sets

A subject set is the set of all subjects that have a specific relation on an
[object](./objects). They empower ORY Keto to be as flexible as you need it
by defining indirections. They can be used to realize e.g. [RBAC](/TODO) or
[inheritance of relations](/TODO). Subject sets themselves can again indirect to
subject sets. For a performant evaluation of requests it is however required to
follow some [best practices](../performance). As a special case, subject
sets can also refer to an object by using the empty relation. Effectively, this
is interpreted as "any relation, even a non-existent one".

Subject sets also represent all intermediary nodes in
[the graph of relations](/TODO).
