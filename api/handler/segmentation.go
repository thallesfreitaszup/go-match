package handler

import (
	"github.com/labstack/echo/v4"
	"go-match/api/request"
	"go-match/internal/domain/segmentation"
	"go-match/internal/segmentation/service"
	"net/http"
)

type Segmentation struct {
	Service service.Segmentation
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
		err := s.createNode(node, segmentationRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.NoContent(http.StatusCreated)
}

func (s Segmentation) createNode(node segmentation.Node, segmentationRequest *request.SegmentationRequest) error {
	regularKey := ""
	if segmentationRequest.Type == request.SimpleKV {
		err := s.Service.CreateSimpleKv(node, segmentationRequest.CircleID)
		if err != nil {
			return err
		}
	} else {
		s.Service.CreateRegularKey(node, segmentationRequest.CircleID, &regularKey)
		err := s.Service.CreateRegular(regularKey, node.Expression(), segmentationRequest.CircleID)
		if err != nil {
			return err
		}
	}
	return nil
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
