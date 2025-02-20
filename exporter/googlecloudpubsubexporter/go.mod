module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter

go 1.23.0

require (
	cloud.google.com/go/pubsub v1.47.0
	github.com/google/uuid v1.6.0
	github.com/googleapis/gax-go/v2 v2.14.1
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/component/componenttest v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/config/configretry v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/confmap v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/consumer v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/exporter v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/exporter/exportertest v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/pdata v1.26.1-0.20250219144032-c2af75d88e89
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
	google.golang.org/api v0.221.0
	google.golang.org/grpc v1.70.0
)

require (
	cloud.google.com/go/auth v0.14.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.7 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/iam v1.3.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/extension v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/featuregate v1.26.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/pipeline v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/receiver v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.58.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.58.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/oauth2 v0.26.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/time v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20250122153221-138b5a5a4fd4 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250124145028-65684f501c47 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
