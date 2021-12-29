package handler

import (
	"github.com/labstack/echo/v4"
	"go-match/api/request"
	"go-match/internal/entity/segmentation"
	"go-match/internal/node"
	"net/http"
)

type Segmentation struct {
	Service node.Service
}

func (s Segmentation) Create(c echo.Context) error {
	nodes := make([]segmentation.Node, 0)
	var segmentationRequest = &request.SegmentationRequest{}
	err := c.Bind(segmentationRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	decomposeSegmentation(segmentationRequest.Node, &nodes)
	for _, node := range nodes {
		if segmentationRequest.Type == request.SimpleKV {
			err := s.Service.CreateSimpleKv(node, segmentationRequest.CircleID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
		} else {
			regularKey := ""
			regularValue := ""
			s.Service.CreateRegularKey(node, segmentationRequest.CircleID, &regularKey)
			s.Service.CreateRegularValue(node, segmentationRequest.CircleID, &regularValue)
			err := s.Service.CreateRegular(regularKey, regularValue[:len(regularValue)-3], segmentationRequest.CircleID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.NoContent(http.StatusCreated)
}

func decomposeSegmentation(nodeRequest request.NodeRequest, segmentations *[]segmentation.Node) {

	if nodeRequest.Type == request.Clause && nodeRequest.LogicalOperator == request.OR {
		for _, clause := range nodeRequest.Clauses {
			decomposeSegmentation(clause, segmentations)
		}
	} else {
		node, err := nodeRequest.ToNode()
		if err == nil {
			*segmentations = append(*segmentations, node)
		}
	}
}
