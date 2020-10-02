module vessel

go 1.14

replace github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel => ./proto/vessel/
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
)
