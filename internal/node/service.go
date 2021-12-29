package node

import (
	"fmt"
	"go-match/internal/entity/segmentation"
	"go-match/internal/eval"
)

const (
	StartKeyDelimiter = "_"
	EndKeyDelimiter   = "_"
)

type Service struct {
	Repository Repository
}

func (s Service) CreateSimpleKv(node segmentation.Node, circleId string) error {
	return s.Repository.CreateSimpleKV(node, circleId)
}

func (s Service) Identify(key string, value interface{}) ([]DB, error) {
	nodes, err := s.Repository.FindSimpleKV(key, value)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (s Service) CreateRegularKey(node segmentation.Node, id string, key *string) {
	if node.Type == segmentation.Clause {
		for _, clause := range node.Clauses {
			s.CreateRegularKey(clause, id, key)
		}
	} else {
		*key += fmt.Sprintf("%s%s", node.Content.Key, EndKeyDelimiter)
	}
}

func (s Service) CreateRegularValue(node segmentation.Node, id string, value *string) {
	if node.Type == segmentation.Clause {
		for _, clause := range node.Clauses {
			s.CreateRegularValue(clause, id, value)
		}
	} else {
		*value += node.Expression()
	}
}

func (s Service) CreateRegular(key string, value string, id string) error {
	return s.Repository.CreateRegular(key, value, id)
}

func (s Service) IdentifyRegular(data map[string]interface{}) ([]DB, error) {
	regularMatched := make([]DB, 0)
	regularNodes, err := s.Repository.FindRegular()
	if err != nil {
		return nil, err
	}
	for _, regularNode := range regularNodes {
		exp := regularNode.Value
		mapParameters := make(map[string]interface{})
		for k, v := range data {
			mapParameters[k] = v
		}
		matched, err := eval.EvalExpression(exp, mapParameters)
		if err != nil {
			return nil, err
		}
		if matched {
			regularMatched = append(regularMatched, regularNode)
		}
	}
	return regularMatched, nil
}
