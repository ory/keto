# Notes

> ORY Keto is still a `sandbox` project and
the included APIs are unstable until we reach `v1` 
and release `v1.0.0` of Keto!
>
> Older API versions, such as `v1alpha1`, will still
> get support for a reasonable amount of time after release
> of `v1`!

This directory contains the ProtoBuf & gRPC definitions
for the Access Control APIs.
    
**ACL is the flexible and scalable "base system"
all other access control schemes built upon.**

## Directory layout

```shell script
keto
└── acl / rbac / ...
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