package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-match/api/handler"
	"go-match/internal/segmentation/repository"
	"go-match/internal/segmentation/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type app struct {
	server *echo.Echo
}

func (a app) Server() http.Handler {
	return a.server
}

func NewApp() app {
	newApp := app{
		server: echo.New(),
	}
	newApp.registerRoutes()
	return newApp
}

func (a app) Start() {

	a.server.Logger.Fatal(a.server.Start(":8080"))
}

func (a app) registerRoutes() {
	host := os.Getenv("MONGO_HOST")
	fmt.Println(host)
	if host == "" {
		host = "mongodb://localhost:27017"
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(host))
	if err != nil {
		log.Fatal(err)
	}
	segmentationHandler := handler.Segmentation{
		Service: service.Segmentation{Repository: repository.RepositoryImpl{
			Client: client,
		}},
	}

	identifyHandler := handler.Identify{
		Service: service.Segmentation{Repository: repository.RepositoryImpl{
			client,
		}},
	}
	a.server.POST("/segmentation", segmentationHandler.Create)
	a.server.POST("/identify", identifyHandler.Identify)
}
