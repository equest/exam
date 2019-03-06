package v1_test

import (
	"context"
	"testing"

	v1 "github.com/equest/exam/internal/v1"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func getGraphService() *v1.GraphService {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "Standar123", ""))
	if err != nil {
		panic(err)
	}
	return v1.NewGraphService(driver)
}

func TestCreate(t *testing.T) {
	g := getGraphService()
	s := v1.NewQuestionService(g)
	ctx := context.TODO()
	q := &v1.Question{
		Kind:    v1.QuestionKindMultipleChoices,
		Heading: "Answer this question",
		Body:    "What is 1 + 1?",
		Footer:  "Please fill with correct answer",
		Answer:  "2",
	}
	err := s.Create(ctx, q)
	if err != nil {
		t.Fatal(err)
	}
	if q.ID == 0 {
		t.Fatal("expecting ID with non zero value")
	}
}
func TestUpdate(t *testing.T) {
	g := getGraphService()
	s := v1.NewQuestionService(g)
	ctx := context.TODO()
	q := &v1.Question{
		Kind:    v1.QuestionKindMultipleChoices,
		Heading: "Answer this question",
		Body:    "What is 1 + 1?",
		Footer:  "Please fill with correct answer",
		Answer:  "2",
	}
	err := s.Create(ctx, q)
	if err != nil {
		t.Fatal(err)
	}
	if q.ID == 0 {
		t.Fatal("expecting ID with non zero value")
	}
	q.Body = "What is 6 - 1?"
	q.Answer = "5"

	err = s.Update(ctx, q.ID, q)
	if err != nil {
		t.Fatal(err)
	}
}
