package config

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    ".gitignore",
		FileModTime: time.Unix(1533889203, 0),
		Content:     string("keys"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "config.development.json",
		FileModTime: time.Unix(1538459324, 0),
		Content:     string("{\n    \"environment\": \"development\",\n    \"dbConnectionString\": \"host=localhost port=5432 user=postgres password=Standar123 dbname=authfazz sslmode=disable\",\n    \"redis\": {\n        \"address\": \"localhost:6789\",\n        \"password\": \"\",\n        \"db\": 0\n    },\n    \"accessTokenDuration\": 28800\n}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "config.go",
		FileModTime: time.Unix(1538473895, 0),
		Content:     string("package config\n\nimport (\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"log\"\n\t\"os\"\n\t\"strconv\"\n\n\t\"github.com/GeertJohan/go.rice\"\n)\n\ntype env string\n\nconst (\n\tenvironment         env = \"AUTHFAZZ_ENV\"\n\tdbConnectionString  env = \"AUTHFAZZ_CONSTR_DB\"\n\tredisAddress        env = \"AUTHFAZZ_REDIS_ADDRESS\"\n\tredisPassword       env = \"AUTHFAZZ_REDIS_PASSWORD\"\n\tredisDB             env = \"AUTHFAZZ_REDIS_DB\"\n\taccessTokenDuration env = \"AUTHFAZZ_ACCESSTOKEN_DURATION\"\n)\n\n// Config , contains application configuration\ntype Config struct {\n\tEnvironment         string `json:\"environment\"`\n\tDBConnectionString  string `json:\"dbConnectionString\"`\n\tRedis               *Redis `json:\"redis\"`\n\tAccessTokenDuration int64  `json:\"accessTokenDuration\"`\n}\n\n// Redis , redis configuration\ntype Redis struct {\n\tAddress  string `json:\"address\"`\n\tPassword string `json:\"password\"`\n\tDB       int    `json:\"db\"`\n}\n\nvar config *Config\n\n// GetConfiguration , get application configuration based on set environment\nfunc GetConfiguration() (*Config, error) {\n\tif config != nil {\n\t\treturn config, nil\n\t}\n\n\tconfig := &Config{\n\t\tAccessTokenDuration: 28800,\n\t}\n\tconfig.Environment = os.Getenv(string(environment))\n\n\tif config.Environment == \"\" {\n\t\treturn nil, fmt.Errorf(\"environment variable '%s' is not set\", environment)\n\t}\n\n\ttemplateBox, errRice := rice.FindBox(\"./\")\n\tif errRice != nil {\n\t\tlog.Fatal(errRice)\n\t}\n\n\tjs, errJS := templateBox.String(fmt.Sprintf(\"config.%s.json\", config.Environment))\n\tif errJS != nil {\n\t\treturn nil, errJS\n\t}\n\n\terrUnmarshal := json.Unmarshal([]byte(js), config)\n\tif errUnmarshal != nil {\n\t\treturn nil, errUnmarshal\n\t}\n\n\tdbConnectionString := os.Getenv(string(dbConnectionString))\n\tif dbConnectionString != \"\" {\n\t\tconfig.DBConnectionString = dbConnectionString\n\t}\n\n\taccessTokenDuration := os.Getenv(string(accessTokenDuration))\n\tatDuration, err := strconv.ParseInt(accessTokenDuration, 10, 64)\n\tif err == nil {\n\t\tconfig.AccessTokenDuration = atDuration\n\t}\n\n\tif config.Redis == nil {\n\t\tconfig.Redis = &Redis{}\n\t}\n\tredisAddress := os.Getenv(string(redisAddress))\n\tif redisAddress != \"\" {\n\t\tconfig.Redis.Address = redisAddress\n\t}\n\tredisPassword := os.Getenv(string(redisPassword))\n\tif redisPassword != \"\" {\n\t\tconfig.Redis.Password = redisPassword\n\t}\n\tredisDB := os.Getenv(string(redisDB))\n\trDB, err := strconv.ParseInt(redisDB, 10, 64)\n\tif err == nil {\n\t\tconfig.Redis.DB = int(rDB)\n\t}\n\n\treturn config, nil\n}\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "config.production.json",
		FileModTime: time.Unix(1538459315, 0),
		Content:     string("{\n    \"environment\": \"production\",\n    \"dbConnectionString\": \"host=localhost port=5432 user=postgres password=Standar123 dbname=authfazz sslmode=disable\",\n    \"redis\": {\n        \"address\": \"localhost:6789\",\n        \"password\": \"\",\n        \"db\": 0\n    },\n    \"accessTokenDuration\": 28800\n}"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "config.staging.json",
		FileModTime: time.Unix(1538459311, 0),
		Content:     string("{\n    \"environment\": \"staging\",\n    \"dbConnectionString\": \"host=localhost port=5432 user=postgres password=Standar123 dbname=authfazz sslmode=disable\",\n    \"redis\": {\n        \"address\": \"localhost:6789\",\n        \"password\": \"\",\n        \"db\": 0\n    },\n    \"accessTokenDuration\": 28800\n}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "config.test.json",
		FileModTime: time.Unix(1538459307, 0),
		Content:     string("{\n    \"environment\": \"test\",\n    \"dbConnectionString\": \"host=localhost port=5432 user=postgres password=Standar123 dbname=authfazz sslmode=disable\",\n    \"redis\": {\n        \"address\": \"localhost:6789\",\n        \"password\": \"\",\n        \"db\": 0\n    },\n    \"accessTokenDuration\": 28800\n}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "config_test.go",
		FileModTime: time.Unix(1538470229, 0),
		Content:     string("package config_test\n\nimport (\n\t\"os\"\n\t\"testing\"\n\n\t\"github.com/payfazz/authfazz/config\"\n)\n\ntype env string\n\nconst (\n\tenvironment         env = \"AUTHFAZZ_ENV\"\n\tdbConnectionString  env = \"AUTHFAZZ_CONSTR_DB\"\n\tredisAddress        env = \"AUTHFAZZ_REDIS_ADDRESS\"\n\tredisPassword       env = \"AUTHFAZZ_REDIS_PASSWORD\"\n\tredisDB             env = \"AUTHFAZZ_REDIS_DB\"\n\taccessTokenDuration env = \"AUTHFAZZ_ACCESSTOKEN_DURATION\"\n)\n\nfunc TestEnvironmentVariablesNotSet(t *testing.T) {\n\tenvVars := getCurrentEnv()\n\tdefer setEnv(envVars)\n\n\tos.Setenv(string(environment), \"\")\n\tos.Setenv(string(dbConnectionString), \"\")\n\tos.Setenv(string(redisAddress), \"\")\n\tos.Setenv(string(redisPassword), \"\")\n\tos.Setenv(string(redisDB), \"\")\n\tos.Setenv(string(accessTokenDuration), \"\")\n\t_, err := config.GetConfiguration()\n\tif err == nil {\n\t\tt.Fatal(\"should error when environment not set\")\n\t}\n}\nfunc TestJsonConfigurationNotFound(t *testing.T) {\n\tenvVars := getCurrentEnv()\n\tdefer setEnv(envVars)\n\n\tos.Setenv(string(environment), \"abc123\")\n\tos.Setenv(string(dbConnectionString), \"\")\n\tos.Setenv(string(redisAddress), \"\")\n\tos.Setenv(string(redisPassword), \"\")\n\tos.Setenv(string(redisDB), \"\")\n\tos.Setenv(string(accessTokenDuration), \"\")\n\tconfig, err := config.GetConfiguration()\n\tif err == nil {\n\t\tt.Fatal(\"should error when environment not found\")\n\t}\n\tif config != nil {\n\t\tt.Fatalf(\"config should fail to load when json configuration file not found\")\n\t}\n}\n\nfunc TestTestConfiguration(t *testing.T) {\n\tenvVars := getCurrentEnv()\n\tdefer setEnv(envVars)\n\n\tos.Setenv(string(environment), \"test\")\n\tos.Setenv(string(dbConnectionString), \"testDbConnectionString\")\n\tos.Setenv(string(redisAddress), \"testRedisAddress\")\n\tos.Setenv(string(redisPassword), \"testRedisPassword\")\n\tos.Setenv(string(redisDB), \"1\")\n\tos.Setenv(string(accessTokenDuration), \"12345\")\n\tconfig, err := config.GetConfiguration()\n\tif err != nil {\n\t\tt.Fatalf(\"an error '%s' was not expected when get test configuration\", err)\n\t}\n\tif config.Environment != \"test\" {\n\t\tt.Fatalf(\"Environment configuration value is not valid. expected : '%s' but got '%s'\", \"test\", config.Environment)\n\t}\n\tif config.DBConnectionString != \"testDbConnectionString\" {\n\t\tt.Fatalf(\"AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'\", \"testAuthConStr\", config.DBConnectionString)\n\t}\n\tif config.Redis.Address != \"testRedisAddress\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"testRedisAddress\", config.Redis.Address)\n\t}\n\tif config.Redis.Password != \"testRedisPassword\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"testRedisPassword\", config.Redis.Password)\n\t}\n\tif config.Redis.DB != 1 {\n\t\tt.Fatalf(\"RedisDB configuration value is not valid. expected : '%v' but got '%v'\", 1, config.Redis.DB)\n\t}\n\tif config.AccessTokenDuration != 12345 {\n\t\tt.Fatalf(\"AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'\", 12345, config.AccessTokenDuration)\n\t}\n}\nfunc TestDevelopmentConfiguration(t *testing.T) {\n\tenvVars := getCurrentEnv()\n\tdefer setEnv(envVars)\n\n\tos.Setenv(string(environment), \"development\")\n\tos.Setenv(string(dbConnectionString), \"developmentDbConnectionString\")\n\tos.Setenv(string(redisAddress), \"developmentRedisAddress\")\n\tos.Setenv(string(redisPassword), \"developmentRedisPassword\")\n\tos.Setenv(string(redisDB), \"1\")\n\tos.Setenv(string(accessTokenDuration), \"23456\")\n\tconfig, err := config.GetConfiguration()\n\tif err != nil {\n\t\tt.Fatalf(\"an error '%s' was not expected when get development configuration\", err)\n\t}\n\tif config.Environment != \"development\" {\n\t\tt.Fatalf(\"Environment configuration value is not valid. expected : '%s' but got '%s'\", \"development\", config.Environment)\n\t}\n\tif config.DBConnectionString != \"developmentDbConnectionString\" {\n\t\tt.Fatalf(\"AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'\", \"developmentAuthConStr\", config.DBConnectionString)\n\t}\n\tif config.Redis.Address != \"developmentRedisAddress\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"developmentRedisAddress\", config.Redis.Address)\n\t}\n\tif config.Redis.Password != \"developmentRedisPassword\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"developmentRedisPassword\", config.Redis.Password)\n\t}\n\tif config.Redis.DB != 1 {\n\t\tt.Fatalf(\"RedisDB configuration value is not valid. expected : '%v' but got '%v'\", 1, config.Redis.DB)\n\t}\n\tif config.AccessTokenDuration != 23456 {\n\t\tt.Fatalf(\"AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'\", 23456, config.AccessTokenDuration)\n\t}\n}\nfunc TestStagingConfiguration(t *testing.T) {\n\tenvVars := getCurrentEnv()\n\tdefer setEnv(envVars)\n\n\tos.Setenv(string(environment), \"staging\")\n\tos.Setenv(string(dbConnectionString), \"stagingDbConnectionString\")\n\tos.Setenv(string(redisAddress), \"stagingRedisAddress\")\n\tos.Setenv(string(redisPassword), \"stagingRedisPassword\")\n\tos.Setenv(string(redisDB), \"1\")\n\tos.Setenv(string(accessTokenDuration), \"34567\")\n\tconfig, err := config.GetConfiguration()\n\tif err != nil {\n\t\tt.Fatalf(\"an error '%s' was not expected when get staging configuration\", err)\n\t}\n\tif config.Environment != \"staging\" {\n\t\tt.Fatalf(\"Environment configuration value is not valid. expected : '%s' but got '%s'\", \"staging\", config.Environment)\n\t}\n\tif config.DBConnectionString != \"stagingDbConnectionString\" {\n\t\tt.Fatalf(\"AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'\", \"stagingAuthConStr\", config.DBConnectionString)\n\t}\n\tif config.Redis.Address != \"stagingRedisAddress\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"stagingRedisAddress\", config.Redis.Address)\n\t}\n\tif config.Redis.Password != \"stagingRedisPassword\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"stagingRedisPassword\", config.Redis.Password)\n\t}\n\tif config.Redis.DB != 1 {\n\t\tt.Fatalf(\"RedisDB configuration value is not valid. expected : '%v' but got '%v'\", 1, config.Redis.DB)\n\t}\n\tif config.AccessTokenDuration != 34567 {\n\t\tt.Fatalf(\"AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'\", 34567, config.AccessTokenDuration)\n\t}\n}\nfunc TestProductionConfiguration(t *testing.T) {\n\tenvVars := getCurrentEnv()\n\tdefer setEnv(envVars)\n\n\tos.Setenv(string(environment), \"production\")\n\tos.Setenv(string(dbConnectionString), \"productionDbConnectionString\")\n\tos.Setenv(string(redisAddress), \"productionRedisAddress\")\n\tos.Setenv(string(redisPassword), \"productionRedisPassword\")\n\tos.Setenv(string(redisDB), \"1\")\n\tos.Setenv(string(accessTokenDuration), \"45678\")\n\tconfig, err := config.GetConfiguration()\n\tif err != nil {\n\t\tt.Fatalf(\"an error '%s' was not expected when get production configuration\", err)\n\t}\n\tif config.Environment != \"production\" {\n\t\tt.Fatalf(\"Environment configuration value is not valid. expected : '%s' but got '%s'\", \"production\", config.Environment)\n\t}\n\tif config.DBConnectionString != \"productionDbConnectionString\" {\n\t\tt.Fatalf(\"AuthConnectionString configuration value is not valid. expected : '%s' but got '%s'\", \"productionAuthConStr\", config.DBConnectionString)\n\t}\n\tif config.Redis.Address != \"productionRedisAddress\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"productionRedisAddress\", config.Redis.Address)\n\t}\n\tif config.Redis.Password != \"productionRedisPassword\" {\n\t\tt.Fatalf(\"RedisAddress configuration value is not valid. expected : '%s' but got '%s'\", \"productionRedisPassword\", config.Redis.Password)\n\t}\n\tif config.Redis.DB != 1 {\n\t\tt.Fatalf(\"RedisDB configuration value is not valid. expected : '%v' but got '%v'\", 1, config.Redis.DB)\n\t}\n\tif config.AccessTokenDuration != 45678 {\n\t\tt.Fatalf(\"AccessTokenDuration configuration value is not valid. expected : '%v' but got '%v'\", 45678, config.AccessTokenDuration)\n\t}\n}\n\nfunc getCurrentEnv() map[string]string {\n\tenvVars := make(map[string]string)\n\tenvVars[string(environment)] = os.Getenv(string(environment))\n\tenvVars[string(dbConnectionString)] = os.Getenv(string(dbConnectionString))\n\tenvVars[string(redisAddress)] = os.Getenv(string(redisAddress))\n\tenvVars[string(redisPassword)] = os.Getenv(string(redisPassword))\n\tenvVars[string(redisDB)] = os.Getenv(string(redisDB))\n\tenvVars[string(accessTokenDuration)] = os.Getenv(string(accessTokenDuration))\n\treturn envVars\n}\n\nfunc setEnv(envs map[string]string) {\n\tfor key, value := range envs {\n\t\tos.Setenv(key, value)\n\t}\n}\n"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    "keys/private_key",
		FileModTime: time.Unix(1533621180, 0),
		Content:     string("-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA4w5xhil8YFSLptRxzQsiJgQm7DxfVx7nEFAndQDw/7a1VfIf\nhhzZlUYx6u+57kP4+JPhqLMl9hEPnJh2DMPV4wrQAOSe6pDK5UP/xZQx8ygy70lG\nfJ6MVo7mkXKaofKobOhkFIOhqtLU/6CrzFl+KdFIsD7pt+FxV6mMmPbnAvDN+hF5\nNwU6N61WGAZER8z7SSTgayGpuHdUKCdPwfuiUIEX3GxhskzV/ROiS+R/NbQZlsfm\nQqcBJ5FxhOtAVevi9s7x6LLTSQKopuuunSTTtu3ys/hs5m6AqNPPkLKqp6R8iXF1\nLg0DMeQlFHYwEo3oRweMNhfYRzC3ukioSf+GuwIDAQABAoIBADlemeKLMujoE80Y\nWpSzXnJ6lBcWfgR2Q23EwuN2VG5YDONlZP+u5G8qKEyzO6hvNkYgn2DPuyS8VNR9\nVT6OcMmIHtxK57he01UwZDzY3/IPUydQvWWZbd4lBy7y5Q1MUbAK29avF7cgxD6+\nqwncBtusDJCzpLwYU1oR9ftkTyRXl8WzHUQ+/QILNnSCDsTrP8JsVaVxbd6FhKKn\n5sSyqM+dX7mtvVAOcj0OJSHZiit7fk5QG9Pi/5iP4pCdZf42sImsr++2GFOezfJd\nH5UU+ujTf+b4oGirnqgEDRrSr5IyykagWc07D2KJgyPzrkfFDxoB5C/ZC3C6C9AA\nXwzd+GECgYEA5SPDfCMVBRFkYBoxKgbWEElquGiPMDSe+p6QSlX24UXFv8gzdtbT\nf33d27v2cpIOWYym3Er5JiSFq6oCr1cg9+mLP/tNc50sHrdHb8vRfn190nawFJHa\neOe0b3ZePUtAxdd1HaZgq4bNnLYSbi//spdHuu6E1jZrzcmbvIm7PJECgYEA/awp\nrILMDvqHuGNlVr+kdcGfmFxA8y9Z1tZHLgqNjPQQlaOuyJn1cfYbIqghMLjk//Au\nVQ5gfKLc2abHQaVQ2dLqV846eNQvr+cnLQUrUqk41IZuN0HTMbvLHgOLkQNdsUMs\n1TmmPeMxh9X9cLqp7mZoY5CeWeWFOe3EJA1dZIsCgYEAklbf3yUMpJrx7wprQbrx\n9Z7dwH5OjGve6JJh9oemT0LfQ1dZvtj+ZBr/mPkXMR6keX6Bhol/S2Ph1ruSUWck\n0A/gdfFKCr9jUQ6eWgDif5UnyUUxuUFZNQRN0S3Yi+7GpFOxIUmDzagfIqmJZcPT\n2rwQ/IqeXayN9vR+ONABu3ECgYAECn4PdXXytyL6WPsASsU/6vmz36RZO2Pe/ELe\nBOUEXc7100mxgGJckmMURkFhGVDsktLqH/SBh8ak4PdDoHKNRcLd6zcbPaYU00XY\nfcCW7IMvP4T59F586FTwAXZztO4FKODJ9MUlLz1WwJ3s8cxLM+5tx5v+Kp3YsmTx\nfhUCyQKBgDCEkFexrqC2a1rHLh+pwTyvnE4JCVNt72FF8L51aEsG5tGGFvTvgUN6\nIlRCYASNhUK/3+hu337uOSolKXu0W+dFnp1/OLo6sUkuhxWGx3YLwGJygjSrOl5f\n3wIikQ0U/RjRr+/pI0/yw/w3Xcr7iUjei6SBxkiIeZL/749EcLNB\n-----END RSA PRIVATE KEY-----"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    "keys/public_key.pub",
		FileModTime: time.Unix(1533882821, 0),
		Content:     string("-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4w5xhil8YFSLptRxzQsi\nJgQm7DxfVx7nEFAndQDw/7a1VfIfhhzZlUYx6u+57kP4+JPhqLMl9hEPnJh2DMPV\n4wrQAOSe6pDK5UP/xZQx8ygy70lGfJ6MVo7mkXKaofKobOhkFIOhqtLU/6CrzFl+\nKdFIsD7pt+FxV6mMmPbnAvDN+hF5NwU6N61WGAZER8z7SSTgayGpuHdUKCdPwfui\nUIEX3GxhskzV/ROiS+R/NbQZlsfmQqcBJ5FxhOtAVevi9s7x6LLTSQKopuuunSTT\ntu3ys/hs5m6AqNPPkLKqp6R8iXF1Lg0DMeQlFHYwEo3oRweMNhfYRzC3ukioSf+G\nuwIDAQAB\n-----END PUBLIC KEY-----"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    "rice-box.go",
		FileModTime: time.Unix(1538477359, 0),
		Content:     string(""),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1537773519, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // ".gitignore"
			file3, // "config.development.json"
			file4, // "config.go"
			file5, // "config.production.json"
			file6, // "config.staging.json"
			file7, // "config.test.json"
			file8, // "config_test.go"
			filec, // "rice-box.go"

		},
	}
	dir9 := &embedded.EmbeddedDir{
		Filename:   "keys",
		DirModTime: time.Unix(1533882814, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filea, // "keys/private_key"
			fileb, // "keys/public_key.pub"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir9, // "keys"

	}
	dir9.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`./`, &embedded.EmbeddedBox{
		Name: `./`,
		Time: time.Unix(1537773519, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":     dir1,
			"keys": dir9,
		},
		Files: map[string]*embedded.EmbeddedFile{
			".gitignore":              file2,
			"config.development.json": file3,
			"config.go":               file4,
			"config.production.json":  file5,
			"config.staging.json":     file6,
			"config.test.json":        file7,
			"config_test.go":          file8,
			"keys/private_key":        filea,
			"keys/public_key.pub":     fileb,
			"rice-box.go":             filec,
		},
	})
}
