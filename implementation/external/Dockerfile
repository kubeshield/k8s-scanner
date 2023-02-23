FROM ubuntu

WORKDIR /external
COPY . .
COPY --chown=65532:65532 certs/tls.crt certs/tls.key /certs/

RUN chmod +x /external/bin/provider
USER 65532:65532

ENTRYPOINT ["/external/bin/provider"]
