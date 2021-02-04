---
id: objects
title: Objects
---

Objects are identifiers for some kind of application objects. They can represent
e.g. a file, network port, physical item, ... . It is up to the application to
map its objects to an unambiguous identifier. We recommend the usage of UUIDs as
they provide a high entropy. It is however totally possible to use e.g. URLs or
opaque tokens of any kind. ORY Keto will consider objects equal iff their string
representation is equal.
