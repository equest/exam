package organization_test

import (
	"context"
	"testing"

	"github.com/equest/exam/internal/organization"
	"github.com/equest/exam/internal/test"
)

func Test_List(t *testing.T) {
	s := test.GetServices().Organizations
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
	s := test.GetServices().Organizations
	ctx := context.TODO()
	q := &organization.Organization{
		Type: organization.TypePrivate,
		Name: "organization_test",
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
	s := test.GetServices().Organizations
	ctx := context.TODO()
	q := &organization.Organization{
		Type: organization.TypePrivate,
		Name: "organization_test",
	}
	err := s.Create(ctx, q)
	if err != nil {
		t.Fatal(err)
	}
	if q.ID == 0 {
		t.Fatal("expecting ID with non zero value")
	}
	q.Name = q.Name + ":update"
	err = s.Update(ctx, q.ID, q)
	if err != nil {
		t.Fatal(err)
	}
}
