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

	PgSQL Configuration
*/
type pgSqlConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	SslMode  string `json:"ssl_mode"`
}

func (p pgSqlConfig) ConnStr() string {
	return fmt.Sprintf("host=%s port=%d database=%s sslmode=%s", p.Host, p.Port, p.Database, p.SslMode)
}

func defaultPgSql() pgSqlConfig {
	return pgSqlConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "todo",
		SslMode:  "disable",
	}
}

func (p *pgSqlConfig) loadFromEnv() {
	loadEnvString("APP_PGSQL_HOST", &p.Host)
	loadEnvUint("APP_PGSQL_PORT", &p.Port)
	loadEnvString("APP_PGSQL_DB_NAME", &p.Database)
	loadEnvString("APP_PGSQL_SSLMODE", &p.SslMode)
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
	PgSql  pgSqlConfig  `json:"pgsql"`
}

func (c *config) loadFromEnv() {
	c.Listen.loadFromEnv()
	c.PgSql.loadFromEnv()
}

func defaultConfig() config {
	return config{
		Listen: defaultListenConfig(),
		PgSql:  defaultPgSql(),
	}
}
