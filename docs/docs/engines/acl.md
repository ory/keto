---
id: acl
title: Access Control Lists (ACL)
---

An
[Access Control List (ACL)](https://en.wikipedia.org/wiki/Access_control_list)
is a matrix of users and permissions:

|       | blog_post.create | blog_post.delete | blog_post.modify | blog_post.read |
| ----- | ---------------- | ---------------- | ---------------- | -------------- |
| Alice | yes              | yes              | yes              | yes            |
| Bob   | no               | no               | no               | yes            |
| Peter | yes              | no               | yes              | yes            |

In the example above, Alice has the permission to create a blog post
`(blog_post.create)` while Bob does not. All three (Alice, Bob, Peter) can read
blog posts.

Similarly, you could create a matrix of resources (e.g. blog articles) and each
user's permissions (`c` for `create`, `m` for `modify`, etc) with regards to
that resource:

|       | blog_post.1 | blog_post.2 | blog_post.3 | blog_post.4 |
| ----- | ----------- | ----------- | ----------- | ----------- |
| alice | c,r,m,d     | c,r,m,d     | c,r,m,d     | c,r,m,d     |
| bob   | r           | r           | r           | r           |
| peter | c,r,m,d     | r           | c,r,m,d     | r           |

ACLs are common in applications with few subjects like filesystems (`chmod` /
`chown`).

**Benefits:**

- Fine-grained control that can be fine-tuned per identity and permission.
- Works really well in systems where each identity has a different set of
  permissions.

**Shortcomings:**

- As the number of identities and resources grows over time, the matrix becomes
  large and hard to maintain.
- If many identities have the same permissions, choose a system like RBAC.

**Implementation status:** Access Control Lists are currently not implemented
but will be first-class citizens in the future. To bump this in priority, please
upvote [this GitHub ticket](https://github.com/ory/keto/issues/61).
