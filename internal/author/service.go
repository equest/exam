package author

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/equest/exam/internal/graph"
)

// Const
const (
	NodeLabel = "Author"
)

// Author ...
type Author struct {
	graph.Node
	Name string `json:"name"`
}

// Service ...
type Service interface {
	Create(ctx context.Context, q *Author) error
	Update(ctx context.Context, id int, q *Author) error
	List(ctx context.Context, page int, size int) ([]Author, error)
}

// GraphService ...
type GraphService struct {
	graph *graph.Service
}

// NewGraphService return new author servie instance
func NewGraphService(graph *graph.Service) Service {
	return &GraphService{
		graph: graph,
	}
}

// Create ...
func (s *GraphService) Create(ctx context.Context, q *Author) error {
	return s.graph.CreateNode(ctx, NodeLabel, q)
}

// Update ...
func (s *GraphService) Update(ctx context.Context, id int, q *Author) error {
	return s.graph.UpdateNode(ctx, id, NodeLabel, q)
}

// List ...
func (s *GraphService) List(ctx context.Context, page int, size int) ([]Author, error) {
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
	authors := []Author{}
	for _, node := range nodes.([]interface{}) {
		var author Author
		bs, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bs, &author)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}
