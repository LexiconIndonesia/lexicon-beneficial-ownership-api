package main

import (
	"fmt"
	"os"
	"strconv"
)

func loadEnvString(key string, result *string) {
	s, ok := os.LookupEnv(key)

	if !ok {
		return
	}
	*result = s
}

func loadEnvUint(key string, result *uint) {
	s, ok := os.LookupEnv(key)

	if !ok {
		return
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	*result = uint(n)
}

/* Configuration */

/*
*

	MYSQL Configuration
*/
type mysqlConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m mysqlConfig) ConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Database)
}

func defaultMysqlConfig() mysqlConfig {
	return mysqlConfig{
		Host:     "localhost",
		Port:     3306,
		Database: "document",
		Username: "root",
		Password: "password",
	}
}

func (m *mysqlConfig) loadFromEnv() {
	loadEnvString("APP_MYSQL_HOST", &m.Host)
	loadEnvUint("APP_MYSQL_PORT", &m.Port)
	loadEnvString("APP_MYSQL_DB_NAME", &m.Database)
	loadEnvString("APP_MYSQL_USERNAME", &m.Username)
	loadEnvString("APP_MYSQL_PASSWORD", &m.Password)
}

type listenConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (l listenConfig) Addr() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func defaultListenConfig() listenConfig {
	return listenConfig{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (l *listenConfig) loadFromEnv() {

	loadEnvString("APP_LISTEN_HOST", &l.Host)
	loadEnvUint("APP_LISTEN_PORT", &l.Port)

}

type config struct {
	Listen listenConfig `json:"listen"`
	MySql  mysqlConfig  `json:"mysql"`
}

func (c *config) loadFromEnv() {
	c.Listen.loadFromEnv()
	c.MySql.loadFromEnv()
}

func defaultConfig() config {
	return config{
		Listen: defaultListenConfig(),
		MySql:  defaultMysqlConfig(),
	}
}
