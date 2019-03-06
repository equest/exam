package config_test

import (
	"os"
	"testing"

	"github.com/payfazz/authfazz/config"
)

type env string

const (
	environment         env = "AUTHFAZZ_ENV"
	dbConnectionString  env = "AUTHFAZZ_CONSTR_DB"
	redisAddress        env = "AUTHFAZZ_REDIS_ADDRESS"
	redisPassword       env = "AUTHFAZZ_REDIS_PASSWORD"
	redisDB             env = "AUTHFAZZ_REDIS_DB"
	accessTokenDuration env = "AUTHFAZZ_ACCESSTOKEN_DURATION"
	sendfazzHost        env = "AUTHFAZZ_SENDFAZZ_HOST"
	sendfazzSender      env = "AUTHFAZZ_SENDFAZZ_SENDER"
)

func TestEnvironmentVariablesNotSet(t *testing.T) {
	envVars := getCurrentEnv()
	defer setEnv(envVars)

	os.Setenv(string(environment), "")
	os.Setenv(string(dbConnectionString), "")
	os.Setenv(string(redisAddress), "")
	os.Setenv(string(redisPassword), "")
	os.Setenv(string(redisDB), "")
	os.Setenv(string(accessTokenDuration), "")
	os.Setenv(string(sendfazzHost), "")
	os.Setenv(string(sendfazzSender), "")
	_, err := config.GetConfiguration()
	if err == nil {
		t.Fatal("should error when environment not set")
	}
}
func TestJsonConfigurationNotFound(t *testing.T) {
	envVars := getCurrentEnv()
	defer setEnv(envVars)

	os.Setenv(string(environment), "abc123")
	os.Setenv(string(dbConnectionString), "")
	os.Setenv(string(redisAddress), "")
	os.Setenv(string(redisPassword), "")
	os.Setenv(string(redisDB), "")
	os.Setenv(string(accessTokenDuration), "")
	os.Setenv(string(sendfazzHost), "")
	os.Setenv(string(sendfazzSender), "")
	config, err := config.GetConfiguration()
	if err == nil {
		t.Fatal("should error when environment not found")
	}
	if config != nil {
		t.Fatalf("config should fail to load when json configuration file not found")
	}
}

func TestTestConfiguration(t *testing.T) {
	envVars := getCurrentEnv()
	defer setEnv(envVars)

	os.Setenv(string(environment), "test")
	os.Setenv(string(dbConnectionString), "testDbConnectionString")
	os.Setenv(string(redisAddress), "testRedisAddress")
	os.Setenv(string(redisPassword), "testRedisPassword")
	os.Setenv(string(redisDB), "1")
	os.Setenv(string(accessTokenDuration), "12345")
	os.Setenv(string(sendfazzHost), "testSendfazzHost")
	os.Setenv(string(sendfazzSender), "testSendfazzSender")
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get test configuration", err)
	}
	if config.Environment != "test" {
		t.Fatalf("Environment configuration value is not valid. expected : '%s' but got '%s'", "test", config.Environment)
	}
	if config.DBConnectionString != "testDbConnectionString" {
		t.Fatalf("AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'", "testAuthConStr", config.DBConnectionString)
	}
	if config.Redis.Address != "testRedisAddress" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "testRedisAddress", config.Redis.Address)
	}
	if config.Redis.Password != "testRedisPassword" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "testRedisPassword", config.Redis.Password)
	}
	if config.Redis.DB != 1 {
		t.Fatalf("RedisDB configuration value is not valid. expected : '%v' but got '%v'", 1, config.Redis.DB)
	}
	if config.AccessTokenDuration != 12345 {
		t.Fatalf("AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'", 12345, config.AccessTokenDuration)
	}
	if config.Sendfazz.Host != "testSendfazzHost" {
		t.Fatalf("SendfazzHost configuration value is not valid. expected : '%s' but got '%s'", "testSendfazzHost", config.Sendfazz.Host)
	}
	if config.Sendfazz.Sender != "testSendfazzSender" {
		t.Fatalf("SendfazzSender configuration value is not valid. expected : '%s' but got '%s'", "testSendfazzSender", config.Sendfazz.Sender)
	}
}
func TestDevelopmentConfiguration(t *testing.T) {
	envVars := getCurrentEnv()
	defer setEnv(envVars)

	os.Setenv(string(environment), "development")
	os.Setenv(string(dbConnectionString), "developmentDbConnectionString")
	os.Setenv(string(redisAddress), "developmentRedisAddress")
	os.Setenv(string(redisPassword), "developmentRedisPassword")
	os.Setenv(string(redisDB), "1")
	os.Setenv(string(accessTokenDuration), "23456")
	os.Setenv(string(sendfazzHost), "developmentSendfazzHost")
	os.Setenv(string(sendfazzSender), "developmentSendfazzSender")
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get development configuration", err)
	}
	if config.Environment != "development" {
		t.Fatalf("Environment configuration value is not valid. expected : '%s' but got '%s'", "development", config.Environment)
	}
	if config.DBConnectionString != "developmentDbConnectionString" {
		t.Fatalf("AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'", "developmentAuthConStr", config.DBConnectionString)
	}
	if config.Redis.Address != "developmentRedisAddress" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "developmentRedisAddress", config.Redis.Address)
	}
	if config.Redis.Password != "developmentRedisPassword" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "developmentRedisPassword", config.Redis.Password)
	}
	if config.Redis.DB != 1 {
		t.Fatalf("RedisDB configuration value is not valid. expected : '%v' but got '%v'", 1, config.Redis.DB)
	}
	if config.AccessTokenDuration != 23456 {
		t.Fatalf("AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'", 23456, config.AccessTokenDuration)
	}
	if config.Sendfazz.Host != "developmentSendfazzHost" {
		t.Fatalf("SendfazzHost configuration value is not valid. expected : '%s' but got '%s'", "developmentSendfazzHost", config.Sendfazz.Host)
	}
	if config.Sendfazz.Sender != "developmentSendfazzSender" {
		t.Fatalf("SendfazzSender configuration value is not valid. expected : '%s' but got '%s'", "developmentSendfazzSender", config.Sendfazz.Sender)
	}
}
func TestStagingConfiguration(t *testing.T) {
	envVars := getCurrentEnv()
	defer setEnv(envVars)

	os.Setenv(string(environment), "staging")
	os.Setenv(string(dbConnectionString), "stagingDbConnectionString")
	os.Setenv(string(redisAddress), "stagingRedisAddress")
	os.Setenv(string(redisPassword), "stagingRedisPassword")
	os.Setenv(string(redisDB), "1")
	os.Setenv(string(accessTokenDuration), "34567")
	os.Setenv(string(sendfazzHost), "stagingSendfazzHost")
	os.Setenv(string(sendfazzSender), "stagingSendfazzSender")
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get staging configuration", err)
	}
	if config.Environment != "staging" {
		t.Fatalf("Environment configuration value is not valid. expected : '%s' but got '%s'", "staging", config.Environment)
	}
	if config.DBConnectionString != "stagingDbConnectionString" {
		t.Fatalf("AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'", "stagingAuthConStr", config.DBConnectionString)
	}
	if config.Redis.Address != "stagingRedisAddress" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "stagingRedisAddress", config.Redis.Address)
	}
	if config.Redis.Password != "stagingRedisPassword" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "stagingRedisPassword", config.Redis.Password)
	}
	if config.Redis.DB != 1 {
		t.Fatalf("RedisDB configuration value is not valid. expected : '%v' but got '%v'", 1, config.Redis.DB)
	}
	if config.AccessTokenDuration != 34567 {
		t.Fatalf("AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'", 34567, config.AccessTokenDuration)
	}
	if config.Sendfazz.Host != "stagingSendfazzHost" {
		t.Fatalf("SendfazzHost configuration value is not valid. expected : '%s' but got '%s'", "stagingSendfazzHost", config.Sendfazz.Host)
	}
	if config.Sendfazz.Sender != "stagingSendfazzSender" {
		t.Fatalf("SendfazzSender configuration value is not valid. expected : '%s' but got '%s'", "stagingSendfazzSender", config.Sendfazz.Sender)
	}
}
func TestProductionConfiguration(t *testing.T) {
	envVars := getCurrentEnv()
	defer setEnv(envVars)

	os.Setenv(string(environment), "production")
	os.Setenv(string(dbConnectionString), "productionDbConnectionString")
	os.Setenv(string(redisAddress), "productionRedisAddress")
	os.Setenv(string(redisPassword), "productionRedisPassword")
	os.Setenv(string(redisDB), "1")
	os.Setenv(string(accessTokenDuration), "45678")
	os.Setenv(string(sendfazzHost), "productionSendfazzHost")
	os.Setenv(string(sendfazzSender), "productionSendfazzSender")
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get production configuration", err)
	}
	if config.Environment != "production" {
		t.Fatalf("Environment configuration value is not valid. expected : '%s' but got '%s'", "production", config.Environment)
	}
	if config.DBConnectionString != "productionDbConnectionString" {
		t.Fatalf("AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'", "productionAuthConStr", config.DBConnectionString)
	}
	if config.Redis.Address != "productionRedisAddress" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "productionRedisAddress", config.Redis.Address)
	}
	if config.Redis.Password != "productionRedisPassword" {
		t.Fatalf("RedisAddress configuration value is not valid. expected : '%s' but got '%s'", "productionRedisPassword", config.Redis.Password)
	}
	if config.Redis.DB != 1 {
		t.Fatalf("RedisDB configuration value is not valid. expected : '%v' but got '%v'", 1, config.Redis.DB)
	}
	if config.AccessTokenDuration != 45678 {
		t.Fatalf("AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'", 45678, config.AccessTokenDuration)
	}
	if config.Sendfazz.Host != "productionSendfazzHost" {
		t.Fatalf("SendfazzHost configuration value is not valid. expected : '%s' but got '%s'", "productionSendfazzHost", config.Sendfazz.Host)
	}
	if config.Sendfazz.Sender != "productionSendfazzSender" {
		t.Fatalf("SendfazzSender configuration value is not valid. expected : '%s' but got '%s'", "productionSendfazzSender", config.Sendfazz.Sender)
	}
}

func getCurrentEnv() map[string]string {
	envVars := make(map[string]string)
	envVars[string(environment)] = os.Getenv(string(environment))
	envVars[string(dbConnectionString)] = os.Getenv(string(dbConnectionString))
	envVars[string(redisAddress)] = os.Getenv(string(redisAddress))
	envVars[string(redisPassword)] = os.Getenv(string(redisPassword))
	envVars[string(redisDB)] = os.Getenv(string(redisDB))
	envVars[string(accessTokenDuration)] = os.Getenv(string(accessTokenDuration))
	envVars[string(sendfazzHost)] = os.Getenv(string(sendfazzHost))
	envVars[string(sendfazzSender)] = os.Getenv(string(sendfazzSender))
	return envVars
}

func setEnv(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}
