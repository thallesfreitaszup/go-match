package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-match/api/handler"
	"go-match/internal/node"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type app struct {
	server *echo.Echo
}

func (a app) Server() http.Handler {
	return a.server
}

func newApp() app {
	return app{
		server: echo.New(),
	}
}

func (a app) start() {
	a.registerRoutes()
	a.server.Logger.Fatal(a.server.Start(":8080"))
}

func (a app) registerRoutes() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	segmentationHandler := handler.Segmentation{
		Service: node.Service{Repository: node.RepositoryImpl{
			client,
		}},
	}

	identifyHandler := handler.Identify{
		Service: node.Service{Repository: node.RepositoryImpl{
			client,
		}},
	}
	a.server.POST("/segmentation", segmentationHandler.Create)
	a.server.POST("/identify", identifyHandler.Identify)
}
