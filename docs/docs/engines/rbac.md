---
id: rbac
title: Role Based Access Control (RBAC)
---

[Role Based Access Control (RBAC)](https://en.wikipedia.org/wiki/Role-based_access_control)
maps subjects to roles and roles to permissions. The goal of (H)RBAC is to make
permission management convenient by grouping subjects in roles and assigning
permissions roles. This type of access control is common in web applications
where one often encounters roles such as "administrator", "moderator", and so
on.

In **Hierarchical Role Based Access Control (HRBAC)** roles can inherit
permissions from other roles. The "administrator" role, for example, could
inherit all permissions from the "moderator" role. This reduces duplication and
management complexity around defining privileges.

Let's come back to Alice, Bob, Peter, and blog posts and the matrix from the ACL
example. This time we model the access rights using (H)RBAC and the roles
"reader", "author", and "admin":

![(H)RBAC Example](/images/docs/keto/rbac.png)

`Admin` inherits all privileges from `author`, which inherits from `reader`.
Only `Alice` (or rather her role `admin`) can delete blog posts, whereas
`author` can create and modify blog posts.

(H)RBAC is everywhere. If you ever installed a forum software such as
[phpBB](https://www.phpbb.com/support/docs/en/3.1/ug/adminguide/permissions_roles/)
or [Wordpress](https://codex.wordpress.org/Roles_and_Capabilities), you have
definitely encountered ACL, (H)RBAC, or both.

(H)RBAC reduces management complexity & overhead with large user/subject bases.
Sometimes however, even (H)RBAC is not enough. An example is when you need to
express ownership (e.g. `bob` can only modify his own blog posts), have
attributes (e.g. `bob` works in department `blog`), or in multi-tenant
environments.

**Benefits:**

- Reduces management complexity when many identities share similar permissions.
- Role hierarchies can reduce redundancy even further.
- Is well established and easily understood by many developers as it is a
  de-facto standard for web applications.

**Shortcomings:**

- Has no concept of context:
  - There is no concept of ownership: _Dan is the author of article "Hello
    World" and is thus allowed to update it_.
  - There is no concept of environment: _Dan is allowed to access accounting
    services when the request comes from IP 10.0.0.3_.
  - There is no concept of tenants: _Dan is allowed to access resources on the
    "dan's test" tenant_.

**Implementation status:** (Hierarchical) Role Based Access Control is currently
not implemented but will be first-class citizens in the future. To bump this in
priority, please upvote
[this GitHub ticket](https://github.com/ory/keto/issues/60).
