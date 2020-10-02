// Package main provides ...
package main

import (
	"context"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	"log"
)

type handler struct {
	repository
}

func (this *handler) FindAvailable(ctx context.Context, req *vs.Specification, res *vs.Response) error {
	log.Println("HANDLER:: FindAvailable", req)
	vessel, err := this.repository.FindAvailable(ctx, MarshallSpecification(req))
	if err != nil {
		return err
	}
	res.Vessel = UnMarshallVessel(vessel)
	return nil
}

func (this *handler) Create(ctx context.Context, req *vs.Vessel, res *vs.Response) error {
	log.Println("HANDLER:: Create")
	if err := this.repository.Create(ctx, MarshallVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
