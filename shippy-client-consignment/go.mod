module main


replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
replace github.com/VictorEspiritu/shippy-service-consignment/proto/consignment => ../shippy-service-consignment/proto/consignment

go 1.14

require (
	github.com/VictorEspiritu/shippy-service-consignment/proto/consignment v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.32.0
)
