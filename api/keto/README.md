# Notes

> ORY Keto is still a `sandbox` project.
This makes the included api version `v1` subject
to have breaking changes until the `v1.0.0` release of Keto!

This directory contains the ProtoBuf & gRPC definitions
for the Access Control APIs.

This includes:
- ACL
- Soon:
    - RBAC
    - ABAC
    
**ACL is the flexible and scalable "base system"
where all other access control schemes built upon.**

## Directory layout

```shell script
keto
└── acl / rbac / abac
    ├── node
    │   └── v1 - Intercommunication API (cluster internal)
    ├── admin
    │   └── v1 - Admin API definitions
    └── v1 - "Base" API definitions
```

- `admin` - API for critical administrative tasks
  - namespace config management
  - retrieval of cluster system statistics / analysis
  - etc.
- `base/v1` - Base APIs / models
  - read/write/watch/... services
- `node` - Intercommunication of cluster nodes