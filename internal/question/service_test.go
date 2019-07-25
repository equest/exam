package question_test

import (
	"context"
	"testing"

	"github.com/equest/exam/internal/question"
	"github.com/equest/exam/internal/test"
)

func Test_List(t *testing.T) {
	s := test.GetServices().Questions
	ctx := context.TODO()
	data, err := s.List(ctx, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 5 {
		t.Fatal("expecting ID with non zero value")
	}
}

func Test_Create(t *testing.T) {
	s := test.GetServices().Questions
	ctx := context.TODO()
	q := &question.Question{
		Kind:         question.QuestionKindMultipleChoices,
		Heading:      "Answer this question",
		Body:         "What is 1 + 1?",
		Footer:       "Please fill with correct answer",
		Answer:       "2",
		ReferenceURI: "question_test",
	}
	err := s.Create(ctx, q)
	if err != nil {
		t.Fatal(err)
	}
	if q.ID == 0 {
		t.Fatal("expecting ID with non zero value")
	}
}

func Test_Update(t *testing.T) {
	s := test.GetServices().Questions
	ctx := context.TODO()
	q := &question.Question{
		Kind:         question.QuestionKindMultipleChoices,
		Heading:      "Answer this question",
		Body:         "What is 1 + 1?",
		Footer:       "Please fill with correct answer",
		Answer:       "2",
		ReferenceURI: "question_test",
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
