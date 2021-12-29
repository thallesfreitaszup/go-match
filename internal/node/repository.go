package node

import (
	"context"
	"fmt"
	"go-match/internal/entity/segmentation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateSimpleKV(segmentation segmentation.Node, circleId string) error
	CreateRegular(key, value, circleId string) error
	FindSimpleKV(key, value interface{}) ([]DB, error)
	FindRegular() ([]DB, error)
}

type RepositoryImpl struct {
	Client *mongo.Client
}

func (r RepositoryImpl) CreateSimpleKV(node segmentation.Node, circleId string) error {
	collection := r.Client.Database("matcher").Collection("node")
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
	collection := r.Client.Database("matcher").Collection("node")
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

func (r RepositoryImpl) FindSimpleKV(key interface{}, value interface{}) ([]DB, error) {
	var nodeDB DB
	nodeArray := make([]DB, 0)

	collection := r.Client.Database("matcher").Collection("node")
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

func (r RepositoryImpl) FindRegular() ([]DB, error) {
	var nodeDB DB
	nodeArray := make([]DB, 0)

	collection := r.Client.Database("matcher").Collection("node")
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
