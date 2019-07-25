package question

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/equest/exam/internal/graph"
)

// Available question kinds
const (
	QuestionKindMultipleChoices = "multiple-choices"
	QuestionKindEssay           = "essay"
)

// Node labels
const (
	NodeLabel = "Question"
)

// Question is question
type Question struct {
	graph.Node
	Kind         string `json:"kind"`
	Heading      string `json:"heading"`
	Body         string `json:"body"`
	Footer       string `json:"footer"`
	Answer       string `json:"answer"`
	ReferenceURI string `json:"referenceURI"`
	Public       bool   `json:"public"`
}

// Service ...
type Service interface {
	Create(ctx context.Context, q *Question) error
	Update(ctx context.Context, id int, q *Question) error
	List(ctx context.Context, page int, size int) ([]Question, error)
}

// GraphService ...
type GraphService struct {
	graph *graph.Service
}

// NewGraphService return new question servie instance
func NewGraphService(graph *graph.Service) Service {
	return &GraphService{
		graph: graph,
	}
}

// Create ...
func (s *GraphService) Create(ctx context.Context, q *Question) error {
	return s.graph.CreateNode(ctx, NodeLabel, q)
}

// Update ...
func (s *GraphService) Update(ctx context.Context, id int, q *Question) error {
	return s.graph.UpdateNode(ctx, id, NodeLabel, q)
}

// List ...
func (s *GraphService) List(ctx context.Context, page int, size int) ([]Question, error) {
	offset := (page - 1) * size
	if offset < 0 {
		offset = 0
	}
	nodes, err := s.graph.Read(ctx, func(tx neo4j.Transaction) (interface{}, error) {
		cypher := fmt.Sprintf(`MATCH (n:%s) 
			RETURN n 
			ORDER BY n.createdAt DESC 
			SKIP %v 
			LIMIT %v`, NodeLabel, offset, size)

		result, err := tx.Run(cypher, nil)
		if err != nil {
			return nil, err
		}

		arr := []interface{}{}
		for result.Next() {
			recs := result.Record()
			for i := range recs.Values() {
				node := recs.GetByIndex(i).(neo4j.Node)
				arr = append(arr, node.Props())
			}
		}
		return arr, nil
	})
	if err != nil {
		return nil, err
	}
	questions := []Question{}
	for _, node := range nodes.([]interface{}) {
		var question Question
		bs, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bs, &question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}
