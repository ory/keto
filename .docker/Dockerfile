FROM alpine:3.16.0

RUN addgroup -S ory; \
    adduser -S ory -G ory -D  -h /home/ory -s /bin/nologin; \
    chown -R ory:ory /home/ory

RUN apk --no-cache --update-cache --upgrade --latest add ca-certificates

COPY keto /usr/bin/keto

# Exposing the ory home directory to simplify passing in keto configuration (e.g. if the file $HOME/keto.yaml
# exists, it will be automatically used as the configuration file).
VOLUME /home/ory

# Declare the standard ports used by keto (4433 for read service endpoint, 4434 for write service endpoint)
EXPOSE 4433 4434

USER ory

ENTRYPOINT ["keto"]
CMD ["serve"]
