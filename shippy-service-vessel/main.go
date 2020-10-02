// Package main provides ...
package main

import (
	"context"
	"errors"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
)

type Repository interface {
	FindAvailable(*vs.Specification) (*vs.Vessel, error)
}

type VesselRepository struct {
	vessels []*vs.Vessel
}

func (this *VesselRepository) FindAvailable(spec *vs.Specification) (*vs.Vessel, error) {
	log.Println("Vessel start: ", spec)
	for _, vessel := range this.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No Vessel found by that spec")
}

type VesselService struct {
	repo Repository
}

func (this *VesselService) FindAvailable(ctx context.Context, req *vs.Specification, res *vs.Response) error {
	vessel, err := this.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*vs.Vessel{
		&vs.Vessel{Id: "vessel003", Name: "Mc Book Pro", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	service.Init()

	log.Println("Starting Vessel Service")
	if err := vs.RegisterVesselServiceHandler(service.Server(), &VesselService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
