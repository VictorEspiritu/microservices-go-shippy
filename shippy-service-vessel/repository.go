// Package main provides ...
package main

import (
	"context"
	vs "github.com/VictorEspiritu/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

type Specification struct {
	Capacity  int32
	MaxWeight int32
}

type Vessel struct {
	ID        string
	Capacity  int32
	Name      string
	Available bool
	OwnerID   string
	MaxWeight int32
}

func MarshallSpecification(spec *vs.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnMarshallSpecification(spec *Specification) *vs.Specification {
	return &vs.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func MarshallVessel(vessel *vs.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

func UnMarshallVessel(vessel *Vessel) *vs.Vessel {
	return &vs.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

func (this *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{
			{"$gte", spec.Capacity},
			{"$gte", spec.MaxWeight},
		},
	}}
	vessel := &Vessel{}
	if err := this.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	log.Println("REPO:: FindAvailable", vessel)
	return vessel, nil
}

func (this *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	insert, err := this.collection.InsertOne(ctx, vessel)

	log.Println("REPO:: Create", insert)
	return err
}
