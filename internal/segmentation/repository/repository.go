package repository

import (
	"context"
	"fmt"
	"go-match/internal/domain/segmentation"
	"go-match/internal/segmentation/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Segmentation interface {
	CreateSimpleKV(segmentation segmentation.Node, circleId string) error
	CreateRegular(key, value, circleId string) error
	FindSimpleKV(key, value interface{}) ([]entity.Segmentation, error)
	FindRegular() ([]entity.Segmentation, error)
}

type RepositoryImpl struct {
	Client *mongo.Client
}

func (r RepositoryImpl) CreateSimpleKV(node segmentation.Node, circleId string) error {
	collection := r.Client.Database("matcher").Collection("segmentation")
	_, err := collection.InsertOne(context.TODO(), bson.D{
		{"key", node.Content.Key},
		{"value", node.Content.Value},
		{"circleId", circleId},
		{"type", "SIMPLE_KV"},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) CreateRegular(key, value, circleId string) error {
	collection := r.Client.Database("matcher").Collection("segmentation")
	_, err := collection.InsertOne(context.TODO(), bson.D{
		{"key", key},
		{"value", value},
		{"circleId", circleId},
		{"type", "REGULAR"},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) FindSimpleKV(key interface{}, value interface{}) ([]entity.Segmentation, error) {
	var nodeDB entity.Segmentation
	nodeArray := make([]entity.Segmentation, 0)

	collection := r.Client.Database("matcher").Collection("segmentation")
	filter := bson.D{
		{"key", key},
		{"value", value},
	}
	find, err := collection.Find(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
		return nodeArray, nil
	} else if err != nil {
		return nil, err
	}
	for find.Next(context.TODO()) {
		err := find.Decode(&nodeDB)
		if err != nil {
			return nil, err
		}
		nodeArray = append(nodeArray, nodeDB)
	}
	return nodeArray, nil
}

func (r RepositoryImpl) FindRegular() ([]entity.Segmentation, error) {
	var nodeDB entity.Segmentation
	nodeArray := make([]entity.Segmentation, 0)
	collection := r.Client.Database("matcher").Collection("segmentation")
	filter := bson.D{
		{"type", "REGULAR"},
	}
	find, err := collection.Find(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
		return nodeArray, nil
	} else if err != nil {
		return nil, err
	}
	for find.Next(context.TODO()) {
		err := find.Decode(&nodeDB)
		if err != nil {
			return nil, err
		}
		nodeArray = append(nodeArray, nodeDB)
	}
	return nodeArray, nil
}
