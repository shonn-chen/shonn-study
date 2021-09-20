package client

import (
	"database/sql"
	"log"
	"time"

	"github.com/shonn-study/code-example/go/common/mysql/config"
)

func New(cfg *config.Config) *sql.DB {
	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = 1000
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 3000
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 3000
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 50
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 100
	}
	if cfg.MaxLifetime == 0 {
		cfg.MaxLifetime = 3600000
	}
	dsn := cfg.MysqlConfig().FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open mysql client fail, dsn:%s, error:%s", dsn, err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("ping mysql fail, dsn:%s, err:%s", dsn, err.Error())
	}
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Millisecond)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	return db
}
