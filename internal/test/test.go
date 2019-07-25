package test

import (
	"sync"

	"github.com/equest/exam/config"
	_ "github.com/equest/exam/config/env"
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

var a *app.App
var onceApp sync.Once

// GetApp ...
func GetApp() *app.App {
	onceApp.Do(func() {
		a = &app.App{
			Services: GetServices(),
		}
	})
	return a
}

var services *app.Services
var onceServices sync.Once

// GetServices ...
func GetServices() *app.Services {
	onceServices.Do(func() {
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
		services = &app.Services{
			Auths:         auths,
			Graph:         graph,
			Core:          core.NewGraphService(graph),
			Questions:     question.NewGraphService(graph),
			Subjects:      subject.NewGraphService(graph),
			Authors:       author.NewGraphService(graph),
			Organizations: organization.NewGraphService(graph),
		}
	})
	return services
}
