---
id: index
title: Access Control Engines - Introduction
---

Whatever your system looks like, you probably have a concept of permissions
which models who is allowed to do what (access control). ORY Keto provides you
with battle-tested, best practice access control concepts. Please note that ORY
Keto doesn't support all access control mechanisms while in "sandbox" mode.

This chapter introduces the most widely used Access Control Policies. Before we
do that, let's cover some of the basics.

Every app that has users usually assigns permissions to these users ("Bob and
Alice are allowed to write blog posts"). There are various established best
practices for assigning one or more permissions to one or more users. In the
context of access control, you'll often encounter **users**, **identities** or
**subjects**. They typically include users, robots, cronjobs, services, etc.
