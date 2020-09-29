---
id: index
slug: /
title: Introduction
---

ORY Keto is a permission server that implements best practice access control
mechanisms. If you came looking for the answer to the question:

- Is a certain user allowed to modify this blog article?
- Is this service allowed to print that document?
- Is a member of the ACME organisation allowed to modify data of one of their
  tenants?
- Is this process allowed to execute that worker when coming from IP 10.0.0.2
  between 4pm and 5pm on a Monday?
- ...

ORY Keto provides various access control engines:

- Available today:
  - ORY-flavored Access Control Policies with exact, glob, and regexp matching
    strategies
- Available soon:
  - [Access Control Lists](https://en.wikipedia.org/wiki/Access_control_list)
  - [Role-based Access Control](https://en.wikipedia.org/wiki/Role-based_access_control)
  - Role Based Access Control with Context (Google/Kubernetes-flavored)
  - Amazon Web Services Identity & Access Management Policies (AWS IAM Policies)

Each mechanism is powered by a decision engine implemented on top of the
[Open Policy Agent](https://www.openpolicyagent.org/) and provides well-defined
management and authorization REST API endpoints.
