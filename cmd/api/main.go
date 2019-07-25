package main

import apiserver "github.com/equest/exam/internal/api-server"
import _ "github.com/equest/exam/config/env"

func main() {
	// starts the api server
	s := apiserver.NewServer()
	s.Serve("8080")
}
