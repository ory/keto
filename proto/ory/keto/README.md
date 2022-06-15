# Notes

> ORY Keto is still a `sandbox` project and the included APIs are unstable until we reach `v1` and release `v1.0.0` of Keto!
>
> Older API versions, such as `v1alpha2`, will still get support for a reasonable amount of time after release of `v1`!

This directory contains the ProtoBuf & gRPC definitions for the Access Control APIs.

**ACL is the flexible and scalable "base system" all other access control schemes built upon.**

## Directory layout

```
ory
└── keto
    └── acl
        └── v1* - "Base" API definitions
```

- `acl/v1*`
  - model definitions
  - read
  - write
  - check
  - expand
