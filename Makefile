mdatagen:
	go run github.com/open-telemetry/opentelemetry-collector-contrib/cmd/mdatagen@latest metrics/metadata.yaml

mbuild:
	go run go.opentelemetry.io/collector/cmd/builder@latest --config=otelcol-builder.yaml

build:
	docker compose build

run: build
	docker compose up