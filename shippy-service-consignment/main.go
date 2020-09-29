package main

import (
	"context"
	"log"
	//"net"
	//"sync"
	cs "github.com/VictorEspiritu/shippy-service-consignment/proto/consignment"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	"github.com/micro/go-micro/v2"
)

//const (
//   port = ":50051"
//)

type repository interface {
	Create(*cs.Consignment) (*cs.Consignment, error)
	GetAll() []*cs.Consignment
}

type Repository struct {
	//mu           sync.RWMutex
	consignments []*cs.Consignment
}

func (this *Repository) Create(consignment *cs.Consignment) (*cs.Consignment, error) {
	//this.mu.Lock()
	updated := append(this.consignments, consignment)
	this.consignments = updated
	//this.mu.Unlock()
	//log.Println("REPO - Consignment created:", consignment)
	//log.Println("REPO - ALL:", this.consignments)
	return consignment, nil
}

func (this *Repository) GetAll() []*cs.Consignment {
	return this.consignments
}

type consignmentService struct {
	repo repository
}

func (this *consignmentService) CreateConsignment(ctx context.Context, req *cs.Consignment, res *cs.Response) error {
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
	service.Init()

	//cs.RegisterShippingServiceServer(server, &service{repo})
	//reflection.Register(server)

	//log.Println("Running Service on port: ", port)

	//err = server.Serve(listen)
	//if err != nil {
	//   log.Fatalf("ERROR Failed to serve: %v", err)
	//}
	log.Println("Running Service Discovery.... ")
	err := cs.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo})
	if err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
