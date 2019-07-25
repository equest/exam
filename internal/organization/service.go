package organization

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/equest/exam/internal/graph"
)

// Const
const (
	TypeSchool  = "School"
	TypePrivate = "Private"
	NodeLabel   = "Organization"
)

// Organization ...
type Organization struct {
	graph.Node
	Type string `json:"type"`
	Name string `json:"name"`
}

// Service ...
type Service interface {
	Create(ctx context.Context, q *Organization) error
	Update(ctx context.Context, id int, q *Organization) error
	List(ctx context.Context, page int, size int) ([]Organization, error)
}

// GraphService ...
type GraphService struct {
	graph *graph.Service
}

// NewGraphService return new organization servie instance
func NewGraphService(graph *graph.Service) Service {
	return &GraphService{
		graph: graph,
	}
}

// Create ...
func (s *GraphService) Create(ctx context.Context, q *Organization) error {
	return s.graph.CreateNode(ctx, NodeLabel, q)
}

// Update ...
func (s *GraphService) Update(ctx context.Context, id int, q *Organization) error {
	return s.graph.UpdateNode(ctx, id, NodeLabel, q)
}

// List ...
func (s *GraphService) List(ctx context.Context, page int, size int) ([]Organization, error) {
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
	organizations := []Organization{}
	for _, node := range nodes.([]interface{}) {
		var organization Organization
		bs, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bs, &organization)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}
	return organizations, nil
}
