package entity

import "go-match/api/response"

type Segmentation struct {
	Key      string
	Value    string
	CircleID string
}

func (d Segmentation) ToResponse() response.IdentifyResponse {
	return response.IdentifyResponse{
		CircleId: d.CircleID,
	}
}
