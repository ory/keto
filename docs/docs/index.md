---
id: index
slug: /
title: Introduction
---

Ory Keto is a permission server that implements best practice access control
mechanisms. If you came looking for the answer to the question:

- Is a certain user allowed to modify this blog article?
- Is this service allowed to print that document?
- Is a member of the ACME organisation allowed to modify data of one of their
  tenants?
- Is this process allowed to execute that worker when coming from IP 10.0.0.2
  between 4pm and 5pm on a Monday?
- ...

Ory Keto is build based on
[Google's Zanzibar research paper](https://research.google/pubs/pub48190/) and
provides an extensible ACL language.

Soon, there will be native support for:

- [Role-based Access Control](https://en.wikipedia.org/wiki/Role-based_access_control)
- Role Based Access Control with Context (Google/Kubernetes-flavored)
- [Attribute-based Access Control](https://en.wikipedia.org/wiki/Attribute-based_access_control)
- decision engines based on
  [Open Policy Agent](https://www.openpolicyagent.org/)
