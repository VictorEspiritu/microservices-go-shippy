module main

go 1.14

replace github.com/VictorEspiritu/shippy-service-consignment/proto/consignment => ./proto/consignment

replace github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel => ../shippy-service-vessel/proto/vessel/

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/VictorEspiritu/shippy-service-consignment/proto/consignment v0.0.0-00010101000000-000000000000
	github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.6.3 // indirect
	go.mongodb.org/mongo-driver v1.4.1
	google.golang.org/grpc v1.32.0
)
