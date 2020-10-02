// Package main provides ...
package main

import (
	"context"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func dummyData(ctx context.Context, repo repository) {
	vessels := []*vs.Vessel{
		{Id: "vessel003", Name: "Mackbook 30", MaxWeight: 200000, Capacity: 300000},
		{Id: "vessel006", Name: "Iphone 300", MaxWeight: 400000, Capacity: 600000},
		{Id: "vessel009", Name: "IMac 110", MaxWeight: 800000, Capacity: 900000},
	}

	for _, v := range vessels {
		repo.Create(ctx, MarshallVessel(v))
	}
}

func main() {

	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	service.Init()

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	client, err := CreateClient(context.Background(), host, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}
	handler := &handler{repository}

	dummyData(context.Background(), repository)

	log.Println("MAIN:: Starting Vessel Service")
	if err := vs.RegisterVesselServiceHandler(service.Server(), handler); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
