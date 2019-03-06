package v1

import (
	"context"
)

// Available question kinds
const (
	QuestionKindMultipleChoices = "multiple-choices"
	QuestionKindEssay           = "essay"
)

// Question is question
type Question struct {
	Node
	Kind    string `json:"kind"`
	Heading string `json:"heading"`
	Body    string `json:"body"`
	Footer  string `json:"footer"`
	Answer  string `json:"answer"`
}

// QuestionService ...
type QuestionService struct {
	graph *GraphService
}

// NewQuestionService return new question servie instance
func NewQuestionService(graph *GraphService) *QuestionService {
	return &QuestionService{
		graph: graph,
	}
}

// Create ...
func (s *QuestionService) Create(ctx context.Context, q *Question) error {
	return s.graph.CreateNode(ctx, "Question", q)
}

// Update ...
func (s *QuestionService) Update(ctx context.Context, id int, q *Question) error {
	return s.graph.UpdateNode(ctx, id, "Question", q)
}
