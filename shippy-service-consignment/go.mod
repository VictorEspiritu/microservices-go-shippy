module main

replace github.com/VictorEspiritu/shippy-service-consignment/proto/consignment => ./proto/consignment

go 1.14

require (
	github.com/VictorEspiritu/shippy-service-consignment/proto/consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.32.0
)
