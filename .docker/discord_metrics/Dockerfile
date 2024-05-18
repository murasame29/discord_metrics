FROM golang:1.22.2 AS builder
WORKDIR /build

COPY metrics/ metrics/

COPY otelcol-builder.yaml otelcol-builder.yaml

RUN go run go.opentelemetry.io/collector/cmd/builder@latest --config=otelcol-builder.yaml
RUN go build -o /bin/app .

FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY otel-config.yaml otel-config.yaml
COPY --from=builder /bin/app /app/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


CMD ["/app/app","--config","otel-config.yaml"]