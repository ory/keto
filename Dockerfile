FROM golang:1.15-alpine AS builder

RUN apk -U --no-cache add build-base git gcc bash

WORKDIR /go/src/github.com/ory/keto

ADD go.mod go.mod
ADD go.sum go.sum

ENV GO111MODULE on
ENV CGO_ENABLED 1

RUN go mod download

RUN go build -o /usr/bin/pkger github.com/markbates/pkger/cmd/pkger

ADD . .

RUN /usr/bin/pkger && go build -tags sqlite -o /usr/bin/keto

FROM alpine:3.12

RUN addgroup -S ory; \
    adduser -S ory -G ory -D  -h /home/ory -s /bin/nologin; \
    chown -R ory:ory /home/ory

RUN apk add -U --no-cache ca-certificates

COPY --from=builder /usr/bin/keto /usr/bin/keto

# By creating the sqlite folder as the ory user, the mounted volume will be owned by ory:ory, which
# is required for read/write of SQLite.
RUN mkdir -p /var/lib/sqlite
RUN chown ory:ory /var/lib/sqlite
VOLUME /var/lib/sqlite

# Exposing the ory home directory to simplify passing in the configuration.
VOLUME /home/ory

EXPOSE 4466 4467

USER ory

ENTRYPOINT ["keto"]

CMD ["serve"]
