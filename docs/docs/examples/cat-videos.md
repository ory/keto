---
id: cat-videos-example
title: Cat Videos Application Example
---

This example describes a video sharing service. Videos are organized in
directories. Every directory has an owner and every video has the same owner as
it's parent directory. The owner has elevated privileges about the video which
are not modeled individually in Ory Keto. The only other privilege we are
modeling in this example is view access. Every owner has view access to their
objects, but this can be granted to other users as well. The application
interprets the special `*` user ID as any user (even anonymous). Note that Ory
Keto does not interpret this subject any different from other subjects. It also
does not know anything about directory structures or induced ownership.

## Starting the Example

First, [install Keto](../install.md).

Now you can start the example using either `docker-compose` or a bash script.
The bash script requires you to have the `keto` binary in your `$PATH`, while
docker automatically gets the required images.

```shell
# clone the repository if you don't have it yet
git clone git@github.com:ory/keto.git && cd keto

docker-compose -f contrib/cat-videos-example/docker-compose.yml up
# or
./contrib/cat-videos-example/up.sh

# output: all initially created relation tuples

# NAMESPACE       OBJECT          RELATION NAME   SUBJECT
# videos          /cats/1.mp4     owner           videos:/cats#owner
# videos          /cats/1.mp4     view            videos:/cats/1.mp4#owner
# videos          /cats/1.mp4     view            *
# videos          /cats/2.mp4     owner           videos:/cats#owner
# videos          /cats/2.mp4     view            videos:/cats/2.mp4#owner
# videos          /cats           owner           cat lady
# videos          /cats           view            videos:/cats#owner
```

## State of the System

At the current state only one user (with the username `cat lady`) has added
videos. Both of them are in the `/cats` directory which `cat lady` owns.
`/cats/1.mp4` is viewable by anyone (`*`), while `/cats/2.mp4` has no extra
sharing options and can therefore only be viewed by its owner, `cat lady`. Have
a look at all the relation tuple definitions in the
`contrib/cat-videos-example/relation-tuples` directory.

## Simulating the Client

Now you can open a second terminal to run the queries against, just like the
video service client would do. We will use the Keto CLI client in this example,
but you can also use a tool like `curl` or [Postman](https://www.postman.com/)
with [Keto's REST API](../reference/api.mdx) to achieve the same result.

:::info

If you want to run the Keto CLI within **Docker**, set the alias

```shell
alias keto="docker run -it --network cat-videos-example_default -e KETO_GRPC_URL=\"keto:4467\" oryd/keto:latest"
```

in your terminal session.

:::

Set the remote endpoint so that the Keto CLI knows where to connect to (not
necessary if using Docker):

```shell
export KETO_GRPC_URL="127.0.0.1:4467"
```

First off, we get a request by an anonymous user that would like to view
`/cats/2.mp4`. The client now has to ask Keto if that operation should be
allowed or denied.

```shell
# Is "*" allowed to "view" the object "videos":"/cats/2.mp4"?
keto check "*" view videos /cats/2.mp4
# output:

# false
```

We already discussed that this request should not be allowed, but it is always
good to see this in action.

Now `cat lady` wants to change some view permissions of `/cats/1.mp4`. For this,
the video service application has to show all users that are currently allowed
to view the video. It uses Keto's [Expand API](/TODO) to get these data:

```shell
# Who is allowed to "view" the object "videos":"/cats/2.mp4"?
keto expand view videos /cats/1.mp4
# outupt:

# ∪ videos:/cats/1.mp4#view
# ├─ ∪ videos:/cats/1.mp4#owner
# │  ├─ ∪ videos:/cats#owner
# │  │  ├─ ☘ cat lady️
# ├─ ☘ *️
```

Here we can see what the requested user set expands to. In the first branch we
see that every owner of the object is allowed to view it
(`videos:/cats/1.mp4#owner`). In the next step we see that the owners of the
object are actually the owners of `/cats` (`videos:/cats#owner`). Finally, we
see that `cat lady` is the owner of `/cats`. Note that there is no direct
relation tuple that would grant `cat lady` view access on `/cats/1.mp4`. This is
indirectly defined via the ownership relation.

The special user `*` on the other hand was directly granted view access on the
object, as it is a first-level leaf of the expansion tree.

<!--TODO-->

Updating the view permissions will be added here at a later stage.
