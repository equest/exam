package core_test

import (
	"context"
	"testing"

	"github.com/equest/exam/internal/author"
	"github.com/equest/exam/internal/organization"
	"github.com/equest/exam/internal/question"
	"github.com/equest/exam/internal/subject"
	"github.com/equest/exam/internal/test"
)

func Test_Create_Relation(t *testing.T) {
	ctx := context.TODO()
	// subject
	ss := test.GetServices().Subjects
	s := &subject.Subject{
		Code: "Math",
		Name: "Mathematic",
	}
	err := ss.Create(ctx, s)
	if err != nil {
		t.Fatal(err)
	}
	// author
	as := test.GetServices().Authors
	a := &author.Author{
		Name: "Galileo",
	}
	err = as.Create(ctx, a)
	if err != nil {
		t.Fatal(err)
	}
	// organization
	os := test.GetServices().Organizations
	o := &organization.Organization{
		Type: organization.TypePrivate,
		Name: "Alexandria",
	}
	err = os.Create(ctx, o)
	if err != nil {
		t.Fatal(err)
	}

	cs := test.GetServices().Core
	q := &question.Question{
		Kind:         question.QuestionKindMultipleChoices,
		Heading:      "Answer this question",
		Body:         "What is 1 + 1?",
		Footer:       "Please fill with correct answer",
		Answer:       "2",
		ReferenceURI: "https://en.wikipedia.org/wiki/Addition",
	}
	_, err = cs.CreateQuestion(ctx, q, s, a, o)
	if err != nil {
		t.Fatal(err)
	}

}
