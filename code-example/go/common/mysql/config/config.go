package config

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

var localConfig Config

func init() {
	localConfig = Config{
		Host:         "localhost",
		Port:         3306,
		User:         "shonn",
		Password:     "123456",
		Schema:       "shonn_github_db",
		DialTimeout:  1000,
		ReadTimeout:  3000,
		WriteTimeout: 3000,
		MaxIdleConns: 50,
		MaxOpenConns: 100,
		MaxLifetime:  3600000,
	}
}

func LocalConfig() *Config {
	return &localConfig
}

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	Schema       string
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
}

func (c *Config) MysqlConfig() *mysql.Config {
	return &mysql.Config{
		User:                 c.User,
		Passwd:               c.Password,
		Addr:                 fmt.Sprintf("%v:%v", c.Host, c.Port),
		DBName:               c.Schema,
		Params:               map[string]string{"charset": "utf8mb4"},
		Net:                  "tcp",
		MultiStatements:      true,
		AllowNativePasswords: true,
		ReadTimeout:          time.Duration(c.ReadTimeout) * time.Millisecond,
		WriteTimeout:         time.Duration(c.WriteTimeout) * time.Millisecond,
		Timeout:              time.Duration(c.DialTimeout) * time.Millisecond,
	}
}
