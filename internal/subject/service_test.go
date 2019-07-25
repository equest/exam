package subject_test

import (
	"context"
	"testing"

	"github.com/equest/exam/internal/subject"
	"github.com/equest/exam/internal/test"
)

func Test_List(t *testing.T) {
	s := test.GetServices().Subjects
	ctx := context.TODO()
	_, err := s.List(ctx, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Create(t *testing.T) {
	s := test.GetServices().Subjects
	ctx := context.TODO()
	q := &subject.Subject{
		Code: "Math",
		Name: "Mathematic",
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
	s := test.GetServices().Subjects
	ctx := context.TODO()
	q := &subject.Subject{
		Code: "Math",
		Name: "Mathematic",
	}
	err := s.Create(ctx, q)
	if err != nil {
		t.Fatal(err)
	}
	if q.ID == 0 {
		t.Fatal("expecting ID with non zero value")
	}

	err = s.Update(ctx, q.ID, q)
	if err != nil {
		t.Fatal(err)
	}
}
