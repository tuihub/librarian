module github.com/tuihub/librarian

go 1.23.0

require (
	entgo.io/ent v0.14.1
	github.com/BurntSushi/toml v1.4.0
	github.com/PuerkitoBio/goquery v1.9.2
	github.com/ThreeDotsLabs/watermill v1.3.7
	github.com/ThreeDotsLabs/watermill-redisstream v1.4.0
	github.com/ThreeDotsLabs/watermill-sql/v3 v3.0.3
	github.com/blevesearch/bleve/v2 v2.4.2
	github.com/bufbuild/protovalidate-go v0.6.5
	github.com/cayleygraph/cayley v0.7.7
	github.com/cayleygraph/quad v1.3.0
	github.com/dchest/captcha v1.0.0
	github.com/dgraph-io/ristretto v0.1.1
	github.com/fullstorydev/grpchan v1.1.1
	github.com/getsentry/sentry-go v0.28.1
	github.com/go-co-op/gocron/v2 v2.11.0
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20240627104009-3198e0b83bf2
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20240627104009-3198e0b83bf2
	github.com/go-kratos/kratos/v2 v2.8.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.6.0
	github.com/hashicorp/consul/api v1.29.4
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/invopop/jsonschema v0.12.0
	github.com/johannesboyne/gofakes3 v0.0.0-20240701191259-edd0227ffc37
	github.com/lib/pq v1.10.9
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/meilisearch/meilisearch-go v0.28.0
	github.com/minio/minio-go/v7 v7.0.76
	github.com/redis/go-redis/v9 v9.6.1
	github.com/sony/sonyflake v1.2.0
	github.com/stretchr/testify v1.9.0
	github.com/tuihub/protos v0.4.21
	github.com/zhihu/norm v0.1.11
	go.opentelemetry.io/otel v1.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.5.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.5.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.29.0
	go.opentelemetry.io/otel/log v0.5.0
	go.opentelemetry.io/otel/sdk v1.29.0
	go.opentelemetry.io/otel/sdk/log v0.5.0
	go.opentelemetry.io/otel/sdk/metric v1.29.0
	go.uber.org/ratelimit v0.3.1
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.26.0
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948
	google.golang.org/grpc v1.66.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	ariga.io/atlas v0.19.1-0.20240203083654-5948b60a8e43 // indirect
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.34.2-20240717164558-a6c49f84cc0f.2 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/Rican7/retry v0.3.1 // indirect
	github.com/RoaringBitmap/roaring v1.9.3 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/aws-sdk-go v1.49.5 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/bits-and-blooms/bitset v1.12.0 // indirect
	github.com/blevesearch/bleve_index_api v1.1.10 // indirect
	github.com/blevesearch/geo v0.1.20 // indirect
	github.com/blevesearch/go-faiss v1.0.20 // indirect
	github.com/blevesearch/go-porterstemmer v1.0.3 // indirect
	github.com/blevesearch/gtreap v0.1.1 // indirect
	github.com/blevesearch/mmap-go v1.0.4 // indirect
	github.com/blevesearch/scorch_segment_api/v2 v2.2.15 // indirect
	github.com/blevesearch/segment v0.9.1 // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/blevesearch/upsidedown_store_api v1.0.2 // indirect
	github.com/blevesearch/vellum v1.0.10 // indirect
	github.com/blevesearch/zapx/v11 v11.3.10 // indirect
	github.com/blevesearch/zapx/v12 v12.3.10 // indirect
	github.com/blevesearch/zapx/v13 v13.3.10 // indirect
	github.com/blevesearch/zapx/v14 v14.3.10 // indirect
	github.com/blevesearch/zapx/v15 v15.3.13 // indirect
	github.com/blevesearch/zapx/v16 v16.1.5 // indirect
	github.com/bufbuild/protocompile v0.4.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/facebook/fbthrift/thrift/lib/go/thrift v0.0.0-20240106173148-f66811508108 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/gobuffalo/logger v1.0.7 // indirect
	github.com/gobuffalo/packd v1.0.2 // indirect
	github.com/gobuffalo/packr/v2 v2.8.3 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/geo v0.0.0-20210211234256-740aa86cb551 // indirect
	github.com/golang/glog v1.2.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/cel-go v0.21.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.16.1 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jhump/protoreflect v1.15.1 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/markbates/errx v1.1.0 // indirect
	github.com/markbates/oncer v1.0.0 // indirect
	github.com/markbates/safe v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20190728182440-6a916e37a237 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/rs/cors v1.8.3 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/ryszard/goskiplist v0.0.0-20150312221310-2dfbae5fcf46 // indirect
	github.com/shabbyrobe/gocovmerge v0.0.0-20190829150210-3e036491d500 // indirect
	github.com/shirou/gopsutil/v3 v3.23.6 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sony/gobreaker v1.0.0 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/vesoft-inc/nebula-go/v2 v2.6.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	github.com/zclconf/go-cty v1.12.1 // indirect
	go.etcd.io/bbolt v1.3.7 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240822170219-fc7c04adadcd // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240822170219-fc7c04adadcd // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

replace github.com/cayleygraph/cayley v0.7.7 => github.com/cayleygraph/cayley v0.7.7-0.20220304214302-275a7428fb10
