package handler

import (
	"github.com/labstack/echo/v4"
	"go-match/api/request"
	"go-match/api/response"
	"go-match/internal/node"
	"net/http"
)

type Identify struct {
	Service node.Service
}

func (i Identify) Identify(context echo.Context) error {
	var identifyRequest = &request.IdentifyRequest{}
	var responseArray = make([]response.IdentifyResponse, 0)
	err := context.Bind(&identifyRequest)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	for key, value := range identifyRequest.RequestData {
		nodes, error := i.Service.Identify(key, value)
		if error != nil {
			return context.JSON(http.StatusInternalServerError, err)
		}
		responseArray = append(responseArray, toResponseArray(nodes)...)
	}
	regularMatched, err := i.Service.IdentifyRegular(identifyRequest.RequestData)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	responseArray = append(responseArray, toResponseArray(regularMatched)...)
	return context.JSON(http.StatusOK, responseArray)
}

func toResponseArray(nodes []node.DB) []response.IdentifyResponse {
	var responseArray = make([]response.IdentifyResponse, 0)
	for _, node := range nodes {
		responseArray = append(responseArray, node.ToResponse())
	}
	return responseArray
}
