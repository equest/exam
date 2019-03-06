package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/GeertJohan/go.rice"
)

type env string

const (
	environment          env = "AUTHFAZZ_ENV"
	dbConnectionString   env = "AUTHFAZZ_CONSTR_DB"
	redisAddress         env = "AUTHFAZZ_REDIS_ADDRESS"
	redisPassword        env = "AUTHFAZZ_REDIS_PASSWORD"
	redisDB              env = "AUTHFAZZ_REDIS_DB"
	accessTokenDuration  env = "AUTHFAZZ_ACCESSTOKEN_DURATION"
	sendfazzHost         env = "AUTHFAZZ_SENDFAZZ_HOST"
	sendfazzSender       env = "AUTHFAZZ_SENDFAZZ_SENDER"
	loggersSlackEndpoint env = "AUTHFAZZ_LOGGERS_SLACK_ENDPOINT"
	loggersSlackEnabled  env = "AUTHFAZZ_LOGGERS_SLACK_ENABLED"
)

// Config , contains application configuration
type Config struct {
	Environment         string    `json:"environment"`
	DBConnectionString  string    `json:"dbConnectionString"`
	Redis               *Redis    `json:"redis"`
	AccessTokenDuration uint64    `json:"accessTokenDuration"`
	Sendfazz            *Sendfazz `json:"sendfazz"`
	Loggers             *Loggers  `json:"loggers"`
}

// Redis , redis configuration
type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// Sendfazz , sendfazz configuration
type Sendfazz struct {
	Host   string `json:"host"`
	Sender string `json:"sender"`
}

// Loggers ...
type Loggers struct {
	Slack *LoggerSlack `json:"slack"`
}

// LoggerSlack ...
type LoggerSlack struct {
	Endpoint string `json:"endpoint"`
	Enabled  bool   `json:"enabled"`
}

var config *Config

// GetConfiguration , get application configuration based on set environment
func GetConfiguration() (*Config, error) {
	if config != nil {
		return config, nil
	}
	env := os.Getenv(string(environment))

	config, err := load(env)
	if err != nil {
		return nil, err
	}

	overrideDBConnectionString(config)
	overrideRedisConfig(config)
	overrideAccessTokenDuration(config)
	overrideSendfazzConfig(config)
	overrideLoggersConfig(config)

	return config, nil
}

// load , load configuration from config.$ENV.json file
func load(env string) (*Config, error) {
	if env == "" {
		return nil, fmt.Errorf("environment variable '%s' is not set", environment)
	}
	config := &Config{
		Environment:         env,
		AccessTokenDuration: 28800,
		Loggers: &Loggers{
			Slack: &LoggerSlack{},
		},
	}

	templateBox, errRice := rice.FindBox("./")
	if errRice != nil {
		log.Fatal(errRice)
	}

	js, errJS := templateBox.String(fmt.Sprintf("config.%s.json", config.Environment))
	if errJS != nil {
		return nil, errJS
	}

	errUnmarshal := json.Unmarshal([]byte(js), config)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	return config, nil
}

// overrideDBConnectionString , replace config's db connection string value from env variable
func overrideDBConnectionString(config *Config) {
	dbConnectionString := os.Getenv(string(dbConnectionString))
	if dbConnectionString != "" {
		config.DBConnectionString = dbConnectionString
	}
}

// overrideRedisConfig , replace config's redis values from env variables
func overrideRedisConfig(config *Config) {
	if config.Redis == nil {
		config.Redis = &Redis{}
	}
	redisAddress := os.Getenv(string(redisAddress))
	if redisAddress != "" {
		config.Redis.Address = redisAddress
	}
	redisPassword := os.Getenv(string(redisPassword))
	if redisPassword != "" {
		config.Redis.Password = redisPassword
	}
	redisDB := os.Getenv(string(redisDB))
	rDB, err := strconv.ParseInt(redisDB, 10, 64)
	if err == nil {
		config.Redis.DB = int(rDB)
	}
}

// overrideSendfazzConfig , replace config's sendfazz host value from env variable
func overrideSendfazzConfig(config *Config) {
	if config.Sendfazz == nil {
		config.Sendfazz = &Sendfazz{}
	}
	csendfazzHost := os.Getenv(string(sendfazzHost))
	if csendfazzHost != "" {
		config.Sendfazz.Host = csendfazzHost
	}
	sendfazzSender := os.Getenv(string(sendfazzSender))
	if sendfazzSender != "" {
		config.Sendfazz.Sender = sendfazzSender
	}
}

// overrideAccessTokenDuration , replace config's AccessTokenDuration value from env variable
func overrideAccessTokenDuration(config *Config) {
	accessTokenDuration := os.Getenv(string(accessTokenDuration))
	atDuration, err := strconv.ParseInt(accessTokenDuration, 10, 64)
	if err == nil {
		config.AccessTokenDuration = uint64(atDuration)
	}
}

func overrideLoggersConfig(config *Config) {
	loggersSlackEndpoint := os.Getenv(string(loggersSlackEndpoint))
	if loggersSlackEndpoint != "" {
		config.Loggers.Slack.Endpoint = loggersSlackEndpoint
	}
	loggersSlackEnabled := os.Getenv(string(loggersSlackEnabled))
	enabled, err := strconv.ParseBool(loggersSlackEnabled)
	if err != nil {
		config.Loggers.Slack.Enabled = false
	}
	config.Loggers.Slack.Enabled = enabled
}
