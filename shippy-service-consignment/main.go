package main

import (
	"context"
	"log"
	//"net"
	//"sync"
	cs "github.com/VictorEspiritu/shippy-service-consignment/proto/consignment"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	"github.com/micro/go-micro/v2"
)

type repository interface {
	Create(*cs.Consignment) (*cs.Consignment, error)
	GetAll() []*cs.Consignment
}

type Repository struct {
	consignments []*cs.Consignment
}

func (this *Repository) Create(consignment *cs.Consignment) (*cs.Consignment, error) {
	updated := append(this.consignments, consignment)
	this.consignments = updated
	log.Println("REPO - Consignment created:", consignment)
	return consignment, nil
}

func (this *Repository) GetAll() []*cs.Consignment {
	return this.consignments
}

type consignmentService struct {
	repo         repository
	vesselClient vs.VesselService
}

func (this *consignmentService) CreateConsignment(ctx context.Context, req *cs.Consignment, res *cs.Response) error {

	vesselResp, err := this.vesselClient.FindAvailable(context.Background(), &vs.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	req.VesselId = vesselResp.Vessel.Id

	consignment, err := this.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (this *consignmentService) GetConsignments(ctx context.Context, req *cs.GetRequest, res *cs.Response) error {
	consignments := this.repo.GetAll()
	res.Consignments = consignments

	return nil
}

func main() {

	repo := &Repository{}
	//listen, err := net.Listen("tcp", port)
	//if err != nil {
	//   log.Fatal("ERROR Failed to listen %v", err)
	//}
	//server := grpc.NewServer()
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	vesselClient := vs.NewVesselService("shippy.service.vessel", service.Client())
	service.Init()

	log.Println("Running Service Discovery.... ")
	err := cs.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo, vesselClient})
	if err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
