package env

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/equest/exam/config"
)

func init() {
	once.Do(func() {
		config.SetConfig(&env{})
	})
}

type env struct{}

var e *env
var once sync.Once

// Env configuration keys
const (
	DBConnectionString = "EXAM_DB_CONNECTIONSTRING"
	AlertSlackToken    = "EXAM_ALERT_SLACK_TOKEN"
	AlertSlackChannel  = "EXAM_ALERT_SLACK_CHANNEL"
	GraphDBHost        = "EXAM_GRAPHDB_HOST"
	GraphDBUser        = "EXAM_GRAPHDB_USER"
	GraphDBPassword    = "EXAM_GRAPHDB_PASSWORD"
	GraphDBRealm       = "EXAM_GRAPHDB_REALM"
)

func (e *env) DBConnectionString() string {
	return getStringOrDefault(DBConnectionString, "")
}
func (e *env) AlertSlackToken() string {
	return getStringOrDefault(AlertSlackToken, "")
}
func (e *env) AlertSlackChannel() string {
	return getStringOrDefault(AlertSlackChannel, "")
}
func (e *env) GraphDBHost() string {
	return getStringOrDefault(GraphDBHost, "bolt://localhost:7687")
}
func (e *env) GraphDBUser() string {
	return getStringOrDefault(GraphDBUser, "neo4j")
}
func (e *env) GraphDBPassword() string {
	return getStringOrDefault(GraphDBPassword, "Standar123")
}
func (e *env) GraphDBRealm() string {
	return getStringOrDefault(GraphDBRealm, "")
}

// Helper methods
func getStringOrDefault(key, def string) string {
	return getEnvOrDefault(key, def)
}
func getIntOrDefault(key string, def int) int {
	v := getEnvOrDefault(key, fmt.Sprint(def))
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return def
	}
	return int(i)
}

func getBooleanOrDefault(key string, def bool) bool {
	v := getEnvOrDefault(key, fmt.Sprint(def))
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}

func getTime(key string) time.Time {
	v := getEnvOrDefault(key, "1970-01-01T00:00:00+07:00")
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		now := time.Now()
		return time.Date(1970, 1, 1, 0, 0, 0, 0, now.Location())
	}
	return t
}

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
