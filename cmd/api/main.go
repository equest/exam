package main

import (
	"context"
	"fmt"

	"github.com/equest/exam/internal/v1"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "Standar123", ""))
	if err != nil {
		panic(err)
	}
	s := v1.NewQuestionService(driver)
	ctx := context.TODO()
	q := &v1.Question{}
	err = s.CreateQuestion(ctx, q)
	if err != nil {
		panic(err)
	}
	fmt.Print("success")
}
