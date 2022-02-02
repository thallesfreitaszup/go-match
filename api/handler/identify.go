package handler

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/labstack/echo/v4"
	"go-match/api/request"
	"go-match/internal/segmentation/entity"
	"go-match/internal/segmentation/service"
	"log"
	"net/http"
)

type Identify struct {
	Service service.Segmentation
}

func (i Identify) Identify(context echo.Context) error {
	var identifyRequest = &request.IdentifyRequest{}
	var responseArray = make([]interface{}, 0)
	err := context.Bind(&identifyRequest)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	for key, value := range identifyRequest.RequestData {
		nodes, err := i.Service.Identify(key, value)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err)
		}
		responseArray = append(responseArray, toResponseArray(nodes)...)
	}
	regularMatched, err := i.Service.IdentifyRegular(identifyRequest.RequestData)
	if err != nil {
		log.Println(err)
		return context.JSON(http.StatusInternalServerError, err)
	}
	responseArray = append(responseArray, toResponseArray(regularMatched)...)
	return context.JSON(http.StatusOK, mapset.NewSetFromSlice(responseArray))
}

func toResponseArray(nodes []entity.Segmentation) []interface{} {
	var responseArray = make([]interface{}, 0)
	for _, node := range nodes {
		responseArray = append(responseArray, node.ToResponse())
	}
	return responseArray
}
