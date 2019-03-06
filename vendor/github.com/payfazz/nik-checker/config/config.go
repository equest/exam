package config

import (
	"os"
)

const (
	environment        = "PAYFAZZ_NIKCHECKER_ENV"
	dbConnectionString = "PAYFAZZ_NIKCHECKER_CONSTR_DB"
	slackToken         = "PAYFAZZ_NIKCHECKER_SLACK_TOKEN"
	slackAlertChannel  = "PAYFAZZ_NIKCHECKER_SLACK_ALERT_CHANNEL"
)

// Config contains application configuration
type Config struct {
	Environment        string
	DBConnectionString string
	SlackToken         string
	SlackAlertChannel  string
}

var config *Config

func getEnvOrDefault(env string, defaultVal string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultVal
	}
	return e
}

// GetConfiguration , get application configuration based on set environment
func GetConfiguration() (*Config, error) {
	if config != nil {
		return config, nil
	}

	// default configuration
	config := &Config{
		Environment:        getEnvOrDefault(environment, "dev"),
		DBConnectionString: getEnvOrDefault(dbConnectionString, "dbname=postgres host=127.0.0.1 sslmode=disable user=postgres"),
		SlackToken:         getEnvOrDefault(slackToken, ""),
		SlackAlertChannel:  getEnvOrDefault(slackAlertChannel, ""),
	}

	return config, nil
}
