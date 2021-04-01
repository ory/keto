---
id: index
slug: /
title: Introduction
---

Ory Keto is the first and only open source implementation of "Zanzibar: Google's Consistent, Global Authorization System":

> Determining whether online users are authorized to access digital objects is central to preserving privacy. This paper
> presents the design, implementation, and deployment of Zanzibar, a global system for storing and evaluating access control lists.
> Zanzibar provides a uniform data model and configuration language for expressing a wide range of access control policies
> from hundreds of client services at Google, including Calendar, Cloud, Drive, Maps, Photos, and YouTube. Its authorization
> decisions respect causal ordering of user actions and thus provide external consistency amid changes to access control
> lists and object contents. Zanzibar scales to trillions of access control lists and millions of authorization requests
> per second to support services used by billions of people. It has maintained 95th-percentile latency of less than 10 milliseconds and availability of greater than 99.999% over 3 years of production use.
>
> [Source](https://research.google/pubs/pub48190/)

If you need to know if a user (or robot, car, service) is allowed to do something - Ory Keto is the right fit for you.

Currently, Ory Keto implements the basic API contracts for managing and checking relations ("permissions") with HTTP
and gRPC APIs. Future versions will include features such as userset rewrites (e.g. RBAC-style role-permission models),
Zookies, and more. An overview of what is implemented and upcoming can be found at [Implemented and Planned Features](implemented-planned-features.mdx).
