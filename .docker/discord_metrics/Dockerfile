FROM golang:1.22.2 AS builder
WORKDIR /build

COPY metrics/ metrics/

COPY otelcol-builder.yaml otelcol-builder.yaml

RUN go run go.opentelemetry.io/collector/cmd/builder@latest --config=otelcol-builder.yaml
RUN go build -o /bin/app main.go

FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=builder /bin/app /app/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/app"]