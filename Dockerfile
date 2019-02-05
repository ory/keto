FROM golang:1.11-alpine

ARG git_tag
ARG git_commit

RUN apk add --no-cache git build-base

WORKDIR /go/src/github.com/ory/keto

RUN GO111MODULE=off go get -u github.com/gobuffalo/packr/packr

ENV GO111MODULE=on

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

ADD . .

RUN go mod verify
RUN packr
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -X github.com/ory/keto/cmd.Version=$git_tag -X github.com/ory/keto/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/keto/cmd.GitHash=$git_commit" -a -installsuffix cgo -o keto github.com/ory/keto

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/github.com/ory/keto/keto /usr/bin/keto

ENTRYPOINT ["keto"]

CMD ["serve"]
