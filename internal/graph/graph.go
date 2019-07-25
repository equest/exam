package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Node base node struct
type Node struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Service ...
type Service struct {
	driver neo4j.Driver
}

// NewService return new graph service instance
func NewService(driver neo4j.Driver) *Service {
	return &Service{
		driver: driver,
	}
}

// CreateNode ...
func (s *Service) CreateNode(ctx context.Context, label string, node interface{}) error {
	work := func(tx neo4j.Transaction) (interface{}, error) {
		cypher := fmt.Sprintf(`
			CREATE (data:%s $props) 
				SET data.id = id(data),
				data.createdAt = datetime(),
				data.updatedAt = datetime()
			RETURN data`, label)

		params, err := Props(node)
		if err != nil {
			return nil, err
		}
		result, err := tx.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		err = result.Err()
		if err != nil {
			return nil, err
		}
		if result.Next() {
			record := result.Record().GetByIndex(0).(neo4j.Node)
			return record.Props(), nil
		}
		return nil, result.Err()
	}
	result, err := s.Write(ctx, work)
	if err != nil {
		return err
	}
	return copy(result, node)
}

// UpdateNode ...
func (s *Service) UpdateNode(ctx context.Context, id int, label string, node interface{}) error {
	work := func(tx neo4j.Transaction) (interface{}, error) {
		cypher := fmt.Sprintf(`
			MATCH (target:%s {id:$id}) 
				SET target = $props,
				target.id = $id,
				target.updatedAt = datetime()
			RETURN target`, label)

		params, err := Props(node)
		if err != nil {
			return nil, err
		}
		params["id"] = id

		result, err := tx.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		err = result.Err()
		if err != nil {
			return nil, err
		}

		if result.Next() {
			record := result.Record().GetByIndex(0).(neo4j.Node)
			return record.Props(), nil
		}
		return nil, result.Err()
	}
	result, err := s.Write(ctx, work)
	if err != nil {
		return err
	}
	return copy(result, node)
}

// Write ...
func (s *Service) Write(ctx context.Context, work neo4j.TransactionWork) (interface{}, error) {
	session, err := s.driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.WriteTransaction(work)
}

// Read ...
func (s *Service) Read(ctx context.Context, work neo4j.TransactionWork) (interface{}, error) {
	session, err := s.driver.Session(neo4j.AccessModeRead)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session.ReadTransaction(work)
}

func copy(source, destination interface{}) error {
	bs, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, destination)
}

// Props creates props
func Props(node interface{}) (map[string]interface{}, error) {
	props, err := Map(node)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"props": props,
	}, nil
}

// Map parses node to map[string]interface{}
func Map(node interface{}) (map[string]interface{}, error) {
	bs, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	var props map[string]interface{}
	err = json.Unmarshal(bs, &props)
	if err != nil {
		return nil, err
	}
	return props, err
}
