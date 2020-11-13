# To compile this image manually run:
#
# $ packr; GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build; docker build . -t oryd/keto:latest; rm keto; packr clean

# Use distroless/static as minimal base image that contains:
# - ca-certificates
# - A /etc/passwd entry for a root user
# - A /tmp directory
# - tzdata
#
# Refer to https://github.com/GoogleContainerTools/distroless for more details.
FROM gcr.io/distroless/static:latest
USER 1000
COPY keto /usr/bin/keto
CMD ["keto", "serve"]