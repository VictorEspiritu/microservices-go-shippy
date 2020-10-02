// Package main provides ...
package main

import (
	"context"
	cs "github.com/VictorEspiritu/shippy-service-consignment/proto/consignment"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	"github.com/pkg/errors"
	"log"
)

type handler struct {
	repository
	vesselClient vs.VesselService
}

func (this *handler) CreateConsignment(ctx context.Context, req *cs.Consignment, res *cs.Response) error {

	log.Println("HANDLER:: CreateConsignment", req)
	vesselResp, err := this.vesselClient.FindAvailable(context.Background(), &vs.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if vesselResp == nil {
		return errors.New("HANDLER:: Error Fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}

	req.VesselId = vesselResp.Vessel.Id
	if err := this.repository.Create(ctx, MarshallConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (this *handler) GetConsignments(ctx context.Context, req *cs.GetRequest, res *cs.Response) error {
	log.Println("HANDLER:: GetConsignments")
	consignments, err := this.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnMarshallConsignmentCollection(consignments)

	return nil
}
