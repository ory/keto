---
id: namespaces
title: Namespaces
---

ORY Keto knows the concept of namespaces to organize relation tuples. Namespaces
have a configuration that defines the relations, and some other important values
([see reference](/TODO)). Unlike other applications, ORY Keto does **not**
isolate namespaces. Their purpose is to split up the data into coherent
partitions, each with its corresponding configuration. Internally each namespace
has its own table in the database to allow setting individual storage specific
options.
