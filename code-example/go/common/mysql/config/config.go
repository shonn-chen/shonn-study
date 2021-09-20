package config

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

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
