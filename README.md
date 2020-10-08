<h1 align="center"><img src="https://raw.githubusercontent.com/ory/meta/master/static/banners/keto.svg" alt="ORY Keto - Open Source & Cloud Native Access Control Server"></h1>

<h4 align="center">    
    <a href="https://www.ory.sh/chat">Chat</a> |
    <a href="https://community.ory.sh/">Forums</a> |
    <a href="http://eepurl.com/di390P">Newsletter</a><br/><br/>
    <a href="https://www.ory.sh/docs/keto/">Guide</a> |
    <a href="https://www.ory.sh/docs/keto/sdk/api">API Docs</a> |
    <a href="https://godoc.org/github.com/ory/keto">Code Docs</a><br/><br/>
    <a href="https://opencollective.com/ory">Support this project!</a>
</h4>

# This is the next step for ORY Keto :tada:

Be part of our journey to build the next-gen Keto based on
[Google's Zanzibar paper](https://research.google/pubs/pub48190/).

The following is just a high level view on the paper, trying to grasp all the concepts.

## ACL Language

The ACL is represented by `object#relation@user`, while `user` can be a single user,
or a set of users represented by `object#relation` (e.g. users with editing rights on some object).

## Managing content update, important for update ordering

1. > A Zanzibar client (e.g. Google Docs) requests an opaque consistency token called zookie for each content version via a `content change` ACL check.

   The client has to store the token together with the content change.
2. "The client sends this zookie in subsequent ACL check requests to ensure that the check snapshot is at least as
fresh as the timestamp for the content version."

## Zookies

> A zookie is an opaque byte sequence encoding a globally meaningful timestamp that reflects an ACL write,
a client content version, or a read snapshot.

## DB Schema

### Namespaces

Zanzibar clients have to configure their namespace:
1. configure relations
    - optimization via defining relations on relations within the relation definition
2. storage parameters (type of object IDs, sharding)

## API

### Read Relation Tuples

Purpose: display group member ship, shared with, ...

Querying by object or user, optionally constrained by the relation name.

> -> clients can look up a specific membership entry, read all entries in an ACL or group, or look up all groups with a given user as a direct member

> All tuplesets in a read request are processed at a single snapshot

### Write

(Over-)Write one or more tuples. This endpoint should support a locking mechanism to allow
the detection of races.

### Watch

Get all changes starting from a specific zookie. The response contains a zookie with a timestamp of the
response. Watching can be resumed with the response zookie and will not miss any updates.

### Check

1. view check
    The request contains a userset (`object#relation`), authentication token and a zookie corresponding to the object version.
2. content change check
    No zookie in the request, the ACL has to be evaluated at the latest snapshot. Returns the new zookie for the object version.

### Expand

Like read but expands all indirect references.

# Architecture

## Storage

Relation tuples are stored in a database per namespace. Old versions are garbage collected to allow historic evaluation within a certain window.
There is also a global changelog used for the watcher API and optimizations. Changes are commited to both the namespace database and the changelog
in one transaction.
Namespace configuration is stored in an extra database with two tables, one for the config and one for a changelog to allow hot
reloading.
Lastly, all of that is replicated and sharded across multiple locations around the world.

## Serving

This section omits performance measures for the moment and only focuses on consistency and correctness.

### Namespace Config Consistency

Every server in a cluster uses the config changelog to update it's local namespace config. There is a globally
managed list of versions available on every server that ensures servers can continue to respond even without the
config DB. An incoming request will be handled with one of the config versions from that list.

### Check Evaluation

The algorithm is recursive, concurrent and cancels redundant branches.

### Leopard Specialized Indexing

This indexing service is specialized to determine whether a user `U` is member of a group `G`.
To accomplish this, it maps every group to all it's direct or indirect subgroups and every user to
it's direct groups. If `(MEMBER2GROUP(U) ∩ GROUP2GROUP(G)) != 0`, i.e. there exists a group `G'` that is a
subgroup of `G` to any extend, and the user is member of `G'`. This can also be modeled as a reachability problem
on a graph of groups.
The index is build periodically offline. To ensure consistency, the list of changes between the last index build and
the requested timestamp is merged into the index for evaluation.
