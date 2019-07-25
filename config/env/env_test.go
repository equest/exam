package env_test

import (
	"os"
	"testing"

	"github.com/equest/exam/config/env"
	"github.com/equest/exam/config"
)

func Test_Config_DBConnectionString(t *testing.T) {
	val := "postgress://user:password@host:port/db?key=value"
	os.Setenv(env.DBConnectionString, val)
	dbConStr := config.DBConnectionString()
	if dbConStr != val {
		t.Fatalf("expected %s got %s", val, dbConStr)
	}
}
