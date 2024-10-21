module github.com/dirac-lee/ddd_demo_parking

go 1.20

require (
	code.byted.org/gopkg/context v0.0.1
	code.byted.org/gopkg/env v1.6.4
	code.byted.org/gopkg/logs v1.2.23
	code.byted.org/gopkg/logs/v2 v2.1.50
	code.byted.org/gorm/bytedgen v0.3.26
	code.byted.org/lang/gg v0.18.0
	code.byted.org/middleware/hertz v1.13.0
	code.byted.org/middleware/hertz_ext/v2 v2.1.9
	code.byted.org/oec/overpass_common_struct v0.0.0-20240622162336-57b41eecfedf
	code.byted.org/oec/status_code v0.0.0-20240621183437-c9234e2fb68b
	github.com/google/wire v0.6.0
	github.com/luci/go-render v0.0.0-20160219211803-9a04cc21af0f
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.7.2
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.5.6
	gorm.io/gen v0.3.26
	gorm.io/gorm v1.25.10
	gorm.io/plugin/dbresolver v1.5.1
)

require (
	code.byted.org/aiops/apm_vendor_byted v0.0.17 // indirect
	code.byted.org/aiops/metrics_codec v0.0.11 // indirect
	code.byted.org/aiops/monitoring-common-go v0.0.3 // indirect
	code.byted.org/bcc/bcc-go-client v0.0.28 // indirect
	code.byted.org/bcc/conf_engine v0.0.0-20221205092436-e41478ff92fd // indirect
	code.byted.org/bcc/pull_json_model v1.0.3 // indirect
	code.byted.org/bcc/tools v0.0.6 // indirect
	code.byted.org/bytedtrace/bytedtrace-client-go v1.0.54 // indirect
	code.byted.org/bytedtrace/bytedtrace-common/go v0.0.13 // indirect
	code.byted.org/bytedtrace/bytedtrace-conf-provider-client-go v0.0.24 // indirect
	code.byted.org/bytedtrace/interface-go v1.0.20 // indirect
	code.byted.org/bytedtrace/serializer-go v1.0.0 // indirect
	code.byted.org/duanyi.aster/gopkg v0.0.3 // indirect
	code.byted.org/gopkg/apm_vendor_interface v0.0.2 // indirect
	code.byted.org/gopkg/asynccache v0.0.0-20201112072351-d630cb60c767 // indirect
	code.byted.org/gopkg/consul v1.2.4 // indirect
	code.byted.org/gopkg/ctxvalues v0.4.0 // indirect
	code.byted.org/gopkg/debug v0.10.1 // indirect
	code.byted.org/gopkg/etcd_util v2.3.3+incompatible // indirect
	code.byted.org/gopkg/etcdproxy v0.1.1 // indirect
	code.byted.org/gopkg/lang v0.21.8 // indirect
	code.byted.org/gopkg/logid v0.0.0-20211104042040-f78600e482f2 // indirect
	code.byted.org/gopkg/metainfo v0.1.1 // indirect
	code.byted.org/gopkg/metrics v1.4.23 // indirect
	code.byted.org/gopkg/metrics/v3 v3.1.29 // indirect
	code.byted.org/gopkg/metrics/v4 v4.0.26 // indirect
	code.byted.org/gopkg/metrics_core v0.0.23 // indirect
	code.byted.org/gopkg/net2 v1.5.0 // indirect
	code.byted.org/gopkg/pid v0.0.18 // indirect
	code.byted.org/gopkg/pkg v0.0.0-20210817064112-6fe00340bb36 // indirect
	code.byted.org/gopkg/stats v1.2.12 // indirect
	code.byted.org/gopkg/tccclient v1.4.13 // indirect
	code.byted.org/gopkg/thrift v1.4.13 // indirect
	code.byted.org/hystrix/hystrix-go v0.0.0-20190214095017-a2a890c81cd5 // indirect
	code.byted.org/ies/starling_goclient v0.2.7 // indirect
	code.byted.org/iespkg/bytedkits-go/goext v0.4.0 // indirect
	code.byted.org/iespkg/retry-go v0.1.2 // indirect
	code.byted.org/kite/kitex v1.15.3 // indirect
	code.byted.org/kite/kitex-overpass-suite v0.0.24 // indirect
	code.byted.org/kite/kitutil v3.8.4+incompatible // indirect
	code.byted.org/kite/rpal v0.1.12 // indirect
	code.byted.org/lang/trace v0.0.2 // indirect
	code.byted.org/lidar/profiler v0.3.2 // indirect
	code.byted.org/lidar/profiler/hertz v0.0.0-20230801111316-7e5562fd8659 // indirect
	code.byted.org/lidar/profiler/kitex v0.0.0-20230801111316-7e5562fd8659 // indirect
	code.byted.org/log_market/gosdk v0.0.0-20220328031951-809cbf0ba485 // indirect
	code.byted.org/log_market/loghelper v0.1.10 // indirect
	code.byted.org/log_market/tracelog v0.1.4 // indirect
	code.byted.org/log_market/ttlogagent_gosdk v0.0.6 // indirect
	code.byted.org/log_market/ttlogagent_gosdk/v4 v4.0.51 // indirect
	code.byted.org/middleware/fic_client v0.2.8 // indirect
	code.byted.org/oec/lazyman v0.0.0 // indirect
	code.byted.org/oec/lazyman/starling v1.0.1 // indirect
	code.byted.org/oec/lazyman/strutil v1.0.0 // indirect
	code.byted.org/overpass/common v0.0.0-20230608113610-5b75c582b89e // indirect
	code.byted.org/overpass/oec_common_status_code_mgt v0.0.0-20240420181507-3cb739998564 // indirect
	code.byted.org/security/go-spiffe-v2 v1.0.3 // indirect
	code.byted.org/security/memfd v0.0.1 // indirect
	code.byted.org/security/sensitive_finder_engine v0.3.17 // indirect
	code.byted.org/security/zti-jwt-helper-golang v1.0.15 // indirect
	code.byted.org/service_mesh/shmipc v0.2.16 // indirect
	code.byted.org/trace/trace-client-go v1.3.6 // indirect
	code.byted.org/ttarch/byteconf-cel-go v0.0.3 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/antonmedv/expr v1.9.0 // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/bits-and-blooms/bloom/v3 v3.6.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20240202110943-5e26950c5e57 // indirect
	github.com/bytedance/sonic v1.11.2 // indirect
	github.com/caarlos0/env/v6 v6.2.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/choleraehyq/pid v0.0.18 // indirect
	github.com/cloudwego/configmanager v0.2.0 // indirect
	github.com/cloudwego/dynamicgo v0.2.0 // indirect
	github.com/cloudwego/fastpb v0.0.4 // indirect
	github.com/cloudwego/frugal v0.1.14 // indirect
	github.com/cloudwego/hertz v0.9.0 // indirect
	github.com/cloudwego/kitex v0.9.1 // indirect
	github.com/cloudwego/localsession v0.0.2 // indirect
	github.com/cloudwego/netpoll v0.6.0 // indirect
	github.com/cloudwego/thriftgo v0.3.6 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/pprof v0.0.0-20221103000818-d260c55eee4c // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hbollon/go-edlib v1.6.0 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/hertz-contrib/http2 v0.1.1 // indirect
	github.com/hertz-contrib/localsession v0.0.0-20230912121050-49d165b95cbf // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.2.0 // indirect
	github.com/nyaruka/phonenumbers v1.0.56 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/shirou/gopsutil/v3 v3.22.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/smartystreets/assertions v1.2.0 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	github.com/zeebo/errs v1.3.0 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20230525183740-e7c30c78aeb2 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.1.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20220801145646-83ce21fca29f // indirect
	google.golang.org/grpc v1.48.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/hints v1.1.0 // indirect
)
