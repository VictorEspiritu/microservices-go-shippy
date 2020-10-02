package main

import (
	"context"
	"fmt"
	cs "github.com/VictorEspiritu/shippy-service-consignment/proto/consignment"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
)

const (
	defaultHost = "mongodb://datastore:27017"
)

func main() {
	//repo := &Repository{}
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignment")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vs.NewVesselService("shippy.service.vessel", service.Client())
	handler := &handler{repository, vesselClient}

	log.Println("MAIN:: Running Service Discovery.... ")

	if err := cs.RegisterShippingServiceHandler(service.Server(), handler); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
