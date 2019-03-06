package v1

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

// GraphService ...
type GraphService struct {
	driver neo4j.Driver
}

// NewGraphService return new graph service instance
func NewGraphService(driver neo4j.Driver) *GraphService {
	return &GraphService{
		driver: driver,
	}
}

// WriteTransaction ...
func (s *GraphService) WriteTransaction(ctx context.Context, work neo4j.TransactionWork) (interface{}, error) {
	session, err := s.driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.WriteTransaction(work)
}

// CreateNode ...
func (s *GraphService) CreateNode(ctx context.Context, label string, node interface{}) error {
	p, err := s.WriteTransaction(ctx, s.createNodeTxWork(label, node))
	if err != nil {
		return err
	}
	return s.copyValues(node, p)
}

// UpdateNode ...
func (s *GraphService) UpdateNode(ctx context.Context, id int, label string, node interface{}) error {
	p, err := s.WriteTransaction(ctx, s.updateNodeTxWork(id, label, node))
	if err != nil {
		return err
	}
	return s.copyValues(node, p)
}

// createNodeTxWork is a higher-order-function that returns function for creating node
func (s *GraphService) createNodeTxWork(label string, node interface{}) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		cypher := fmt.Sprintf(`CREATE (o:%s $props) 
					SET o.id = id(o),
					o.createdAt = datetime(),
					o.updatedAt = datetime()
				RETURN o`, label)

		bs, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}
		var props map[string]interface{}
		err = json.Unmarshal(bs, &props)
		if err != nil {
			return nil, err
		}

		result, err := tx.Run(cypher, map[string]interface{}{
			"props": props,
		})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			record := result.Record().GetByIndex(0).(neo4j.Node)
			return record.Props(), nil
		}
		return nil, result.Err()
	}
}

// updateNodeTxWork is a higher-order-function that returns function for creating node
func (s *GraphService) updateNodeTxWork(id int, label string, node interface{}) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		cypher := fmt.Sprintf(`MATCH (o:%s {id:$id}) 
					SET o = $props,
					o.id = $id,
					o.updatedAt = datetime()
				RETURN o`, label)

		bs, err := json.Marshal(node)
		if err != nil {
			return nil, err
		}
		var props map[string]interface{}
		err = json.Unmarshal(bs, &props)
		if err != nil {
			return nil, err
		}

		result, err := tx.Run(cypher, map[string]interface{}{
			"id":    id,
			"props": props,
		})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			record := result.Record().GetByIndex(0).(neo4j.Node)
			return record.Props(), nil
		}
		return nil, result.Err()
	}
}

func (s *GraphService) copyValues(destination, source interface{}) error {
	bs, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, destination)
}
