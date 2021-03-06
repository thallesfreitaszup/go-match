package service

import (
	"fmt"
	"go-match/internal/domain/segmentation"
	"go-match/internal/eval"
	"go-match/internal/segmentation/entity"
	"go-match/internal/segmentation/repository"
	"log"
)

const (
	StartKeyDelimiter = "_"
	EndKeyDelimiter   = "_"
)

type Segmentation struct {
	Repository repository.Segmentation
}

func (s Segmentation) CreateSimpleKv(node segmentation.Node, circleId string) error {
	return s.Repository.CreateSimpleKV(node, circleId)
}

func (s Segmentation) Identify(key string, value interface{}) ([]entity.Segmentation, error) {
	nodes, err := s.Repository.FindSimpleKV(key, value)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (s Segmentation) CreateRegularKey(node segmentation.Node, id string, key *string) {
	if node.Type == segmentation.Clause {
		for _, clause := range node.Clauses {
			s.CreateRegularKey(clause, id, key)
		}
	} else {
		*key += fmt.Sprintf("%s%s", node.Content.Key, EndKeyDelimiter)
	}
}

func (s Segmentation) CreateRegularValue(node segmentation.Node, id string, value *string) {
	if node.Type == segmentation.Clause {
		for _, clause := range node.Clauses {
			s.CreateRegularValue(clause, id, value)
		}
	} else {
		*value += node.Expression()
	}
}

func (s Segmentation) CreateRegular(key string, value string, id string) error {
	return s.Repository.CreateRegular(key, value, id)
}

func (s Segmentation) IdentifyRegular(data map[string]interface{}) ([]entity.Segmentation, error) {
	regularMatched := make([]entity.Segmentation, 0)
	regularNodes, err := s.Repository.FindRegular()
	if err != nil {
		return nil, err
	}
	for _, regularNode := range regularNodes {
		exp := regularNode.Value
		matched, err := eval.Expression(exp, data)
		if err != nil {
			log.Println("Error evaluating exp: ", err)
		}
		if matched {
			regularMatched = append(regularMatched, regularNode)
		}
	}
	return regularMatched, nil
}
