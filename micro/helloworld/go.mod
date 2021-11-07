module micro/helloworld

go 1.15

require (
	github.com/evanphx/json-patch/v5 v5.5.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/micro/v3 v3.6.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/grpc v1.41.0 => google.golang.org/grpc v1.26.0
