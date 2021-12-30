package node

import "go-match/api/response"

type DB struct {
	Key      string
	Value    string
	CircleID string
}

func (d DB) ToResponse() response.IdentifyResponse {
	return response.IdentifyResponse{
		CircleId: d.CircleID,
	}
}
