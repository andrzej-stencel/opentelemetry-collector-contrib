module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver

go 1.23.0

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/google/go-cmp v0.6.0
	github.com/mongodb-forks/digest v1.1.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.120.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.120.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.120.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.120.0
	github.com/stretchr/testify v1.10.0
	go.mongodb.org/atlas v0.37.0
	go.opentelemetry.io/collector/component v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/component/componentstatus v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/config/configopaque v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/config/configretry v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/config/configtls v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/confmap v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/consumer v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/extension v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/filter v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/pdata v1.26.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/receiver v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/semconv v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
)

require (
	go.opentelemetry.io/collector/component/componenttest v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/consumer/consumertest v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/extension/xextension v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/receiver/receivertest v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/scraper v0.120.1-0.20250219144032-c2af75d88e89
	go.opentelemetry.io/collector/scraper/scraperhelper v0.120.1-0.20250219144032-c2af75d88e89
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/leodido/go-syslog/v4 v4.2.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20190525184631-5f46317e436b // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.120.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.120.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/featuregate v1.26.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/pipeline v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.120.1-0.20250219144032-c2af75d88e89 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/grpc v1.70.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza => ../../pkg/stanza

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage => ../../extension/storage

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
