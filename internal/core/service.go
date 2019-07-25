package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/equest/exam/internal/author"
	"github.com/equest/exam/internal/graph"
	"github.com/equest/exam/internal/organization"
	"github.com/equest/exam/internal/question"
	"github.com/equest/exam/internal/subject"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// HasSubject ...
type HasSubject struct {
	graph.Node
	Year int `json:"year"`
}

// Service ...
type Service interface {
	CreateQuestion(ctx context.Context, q *question.Question, subject *subject.Subject, author *author.Author, owner *organization.Organization) (interface{}, error)
}

// NewGraphService return new question servie instance
func NewGraphService(graph *graph.Service) Service {
	return &GraphService{
		graph: graph,
	}
}

// GraphService ...
type GraphService struct {
	graph *graph.Service
}

// CreateQuestion ...
func (g *GraphService) CreateQuestion(ctx context.Context, q *question.Question, s *subject.Subject, a *author.Author, o *organization.Organization) (interface{}, error) {
	cypher := fmt.Sprintf(`
		MERGE (question:%s {id:%v})
		ON CREATE SET question = $props, question.id = id(question), question.createdAt = datetime(), question.updatedAt = datetime() 
		
		WITH question

		MATCH (subject:%s {id:%v})
		MATCH (author:%s {id:%v})
		MATCH (organization:%s {id:%v})
 
		MERGE (question)-[hassubject:HAS_SUBJECT]->(subject)
		ON CREATE SET hassubject = $hassubject, hassubject.id = id(hassubject), hassubject.createdAt = datetime(), hassubject.updatedAt = datetime()
 
		MERGE (question)-[authoredby:AUTHORED_BY]->(author)
		ON CREATE SET authoredby = $authoredby, authoredby.id = id(authoredby), authoredby.createdAt = datetime(), authoredby.updatedAt = datetime()

		MERGE (question)-[ownedby:OWNED_BY]->(organization)
		ON CREATE SET ownedby = $ownedby, ownedby.id = id(ownedby), ownedby.createdAt = datetime(), ownedby.updatedAt = datetime() 

		RETURN question`,
		question.NodeLabel, q.ID,
		subject.NodeLabel, s.ID,
		author.NodeLabel, a.ID,
		organization.NodeLabel, o.ID)

	hasSubject, err := graph.Map(&HasSubject{
		Year: 2019,
	})
	authoredBy, err := graph.Map(graph.Node{})
	ownedBy, err := graph.Map(graph.Node{})
	qprops, err := graph.Map(q)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"props":      qprops,
		"hassubject": hasSubject,
		"authoredby": authoredBy,
		"ownedby":    ownedBy,
	}

	relation, err := g.graph.Write(ctx, func(tx neo4j.Transaction) (interface{}, error) {
		log.Print(cypher)
		result, err := tx.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		err = result.Err()
		if err != nil {
			return nil, err
		}

		for result.Next() {
			recs := result.Record()
			node := recs.GetByIndex(0).(neo4j.Node)
			return node.Props(), err
		}
		return nil, nil
	})
	if err != nil {
		return nil, err
	}
	var rel question.Question
	bs, err := json.Marshal(relation)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &rel)
	if err != nil {
		return nil, err
	}
	return &rel, err
}
