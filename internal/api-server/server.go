package apiserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/equest/exam/config"
	"github.com/equest/exam/internal/app"
	"github.com/equest/exam/internal/auth"
	"github.com/equest/exam/internal/author"
	"github.com/equest/exam/internal/core"
	"github.com/equest/exam/internal/graph"
	"github.com/equest/exam/internal/organization"
	"github.com/equest/exam/internal/question"
	"github.com/equest/exam/internal/subject"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Server struct {
	App  *app.App
	http http.Server
}

func NewServer() *Server {
	app := buildApp()
	routes := buildRoutes(app)
	return &Server{
		App: app,
		http: http.Server{
			Handler: routes,
		},
	}
}

func buildApp() *app.App {
	return &app.App{
		Services: buildServices(),
	}
}

func buildServices() *app.Services {
	driver, err := neo4j.NewDriver(config.GraphDBHost(), neo4j.BasicAuth(config.GraphDBUser(), config.GraphDBPassword(), config.GraphDBRealm()))
	if err != nil {
		panic(err)
	}
	graph := graph.NewService(driver)
	auths := auth.NewAWSCognitoAuthService("ap-southeast-1", "ap-southeast-1_IEcEh5SHo")
	err = auths.LoadJWK()
	if err != nil {
		panic(err)
	}

	return &app.Services{
		Auths:         auths,
		Graph:         graph,
		Core:          core.NewGraphService(graph),
		Questions:     question.NewGraphService(graph),
		Subjects:      subject.NewGraphService(graph),
		Authors:       author.NewGraphService(graph),
		Organizations: organization.NewGraphService(graph),
	}
}

func buildClients() *app.Clients {
	return &app.Clients{}
}

// Serve make server to listen and serve on defined address
func (s *Server) Serve(address string) {
	log.Printf("About to listen on %s. Go to http://127.0.0.1:%s", address, address)
	s.http.Addr = fmt.Sprintf(":%s", address)
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	s.App.Services.Auths.LoadJWK()
}

// Shutdown shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
