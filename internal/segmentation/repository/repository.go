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
	collection := r.GetCollection()
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
	collection := r.GetCollection()
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
	segmentationArray := make([]entity.Segmentation, 0)
	collection := r.GetCollection()

	filter := bson.D{
		{"key", key},
		{"value", value},
	}

	data, err := collection.Find(context.TODO(), filter)

	if err == mongo.ErrNoDocuments {
		return segmentationArray, nil
	} else if err != nil {
		return segmentationArray, err
	}

	err = r.mapToSegmentation(data, &segmentationArray)
	if err != nil {
		return segmentationArray, err
	}
	return segmentationArray, nil
}

func (r RepositoryImpl) FindRegular() ([]entity.Segmentation, error) {
	segmentationArray := make([]entity.Segmentation, 0)
	filter := bson.D{
		{"type", "REGULAR"},
	}

	collection := r.GetCollection()

	data, err := collection.Find(context.TODO(), filter)

	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
		return segmentationArray, nil
	} else if err != nil {
		return segmentationArray, err
	}

	err = r.mapToSegmentation(data, &segmentationArray)
	if err != nil {
		return segmentationArray, err
	}

	return segmentationArray, nil
}

func (r RepositoryImpl) GetCollection() *mongo.Collection {
	return r.Client.Database("matcher").Collection("segmentation")
}

func (r RepositoryImpl) mapToSegmentation(find *mongo.Cursor, segArray *[]entity.Segmentation) error {
	var segmentationDB entity.Segmentation

	for find.Next(context.TODO()) {
		err := find.Decode(&segmentationDB)
		if err != nil {
			return err
		}
		*segArray = append(*segArray, segmentationDB)
	}

	return nil
}
