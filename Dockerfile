# To compile this image manually run:
#
# $ packr; GO111MODULE=on GOOS=linux GOARCH=amd64 go build; docker build -t oryd/keto:latest .; rm keto; packr clean
FROM alpine:3.9

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY keto /usr/bin/keto

USER 1000

ENTRYPOINT ["keto"]

CMD ["serve"]
