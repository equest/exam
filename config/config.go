package config

import "sync"

type config interface {
	DBConnectionString() string
	GraphDBHost() string
	GraphDBUser() string
	GraphDBPassword() string
	GraphDBRealm() string
	AlertSlackToken() string
	AlertSlackChannel() string
}

var cfg config
var once sync.Once
var isSet bool

// SetConfig sets active config to c
// this method should be called by config implementor
// and can only be called once
func SetConfig(c config) {
	if isSet {
		panic("config is set more than once")
	}
	once.Do(func() {
		cfg = c
	})
}

// DBConnectionString ...
func DBConnectionString() string {
	return cfg.DBConnectionString()
}

// AlertSlackToken ...
func AlertSlackToken() string {
	return cfg.AlertSlackToken()
}

// AlertSlackChannel ...
func AlertSlackChannel() string {
	return cfg.AlertSlackChannel()
}

// GraphDBHost ...
func GraphDBHost() string {
	return cfg.GraphDBHost()
}

// GraphDBUser ...
func GraphDBUser() string {
	return cfg.GraphDBUser()
}

// GraphDBPassword ...
func GraphDBPassword() string {
	return cfg.GraphDBPassword()
}

// GraphDBRealm ...
func GraphDBRealm() string {
	return cfg.GraphDBRealm()
}
