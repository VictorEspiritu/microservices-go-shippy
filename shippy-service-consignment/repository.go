// Package main provides ...
package main

import (
	"context"
	cs "github.com/VictorEspiritu/shippy-service-consignment/proto/consignment"
	//"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Consignment struct {
	ID          string     `json:id`
	Weight      int32      `json:weight`
	Description string     `json:description`
	Containers  Containers `json:containers`
	VesselID    string     `json:vessel_id`
}

type Container struct {
	ID         string `json:id`
	CustomerID string `json:customer_d`
	UserID     string `json:user_id`
}

type Containers []*Container

func MarshallContainerCollection(containers []*cs.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshallContainer(container))
	}
	return collection
}

func MarshallContainer(container *cs.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

func MarshallConsignment(consignment *cs.Consignment) *Consignment {
	containers := MarshallContainerCollection(consignment.Containers)
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  containers,
		VesselID:    consignment.VesselId,
	}
}

func MarshallConsignmentCollection(consignments []*cs.Consignment) []*Consignment {
	collection := make([]*Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, MarshallConsignment(consignment))
	}
	return collection
}

func UnMarshallConsignment(consignment *Consignment) *cs.Consignment {
	return &cs.Consignment{
		Id:          consignment.ID,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  UnMarshallContainerCollection(consignment.Containers),
		VesselId:    consignment.VesselID,
	}
}

func UnMarshallContainerCollection(containers []*Container) []*cs.Container {
	collection := make([]*cs.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnMarshallContainer(container))
	}
	return collection
}

func UnMarshallContainer(container *Container) *cs.Container {
	return &cs.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

func UnMarshallConsignmentCollection(consignments []*Consignment) []*cs.Consignment {
	collection := make([]*cs.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnMarshallConsignment(consignment))
	}
	return collection
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (this *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	osave, err := this.collection.InsertOne(ctx, consignment)
	log.Println("REPO:: Consignment Create", osave)
	return err
}

func (this *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cursor, err := this.collection.Find(ctx, bson.D{})

	if err != nil {
		log.Println("REPO: Error to getAll consignments: ", err)
		return nil, err
	}

	log.Println("REPO:: GetAll2", cursor)
	var consignments []*Consignment
	for cursor.Next(ctx) {
		var consignment *Consignment
		if err := cursor.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	log.Println("REPO:: Consignments GetAll", consignments)
	return consignments, err
}
