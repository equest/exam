package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/equest/exam/internal/http-server/handlers"
)

var fn = func(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(1 * time.Second)
	log.Println(r.Context().Err())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test:ok"))
	return nil
}

func TestX(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// f := r.Form

		var f1 map[string]string
		json.NewDecoder(r.Body).Decode(&f1)
		log.Printf("%#v", f1)
	}))
	x := url.Values{
		"id":       []string{"10"},
		"password": []string{"my-secret"},
	}

	res, err := s.Client().PostForm(s.URL, x)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
}

func Test_Handler_Request(t *testing.T) {
	executed := false
	server := httptest.NewServer(handlers.Handler(func(w http.ResponseWriter, r *http.Request) error {
		time.Sleep(500 * time.Millisecond)
		executed = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test:ok"))
		return nil
	}))

	c := server.Client()
	c.Timeout = 1 * time.Second
	res, err := c.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %v, got %v", http.StatusOK, res.StatusCode)
	}
	if !executed {
		t.Fatal("expected executed")
	}
}

func Test_Handler_RequestCancelled(t *testing.T) {
	executed := false
	server := httptest.NewServer(handlers.Handler(func(w http.ResponseWriter, r *http.Request) error {
		time.Sleep(500 * time.Millisecond)
		executed = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test:ok"))
		return nil
	}))

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	tr := &http.Transport{}
	c := server.Client()
	c.Transport = tr

	c.Do(req)
	tr.CancelRequest(req)
	if executed {
		t.Fatal("expected !executed")
	}
}
