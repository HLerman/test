package provider

import (
	"database/sql"
	"fmt"

	"github.com/HLerman/test/internal/implementation/db"
	_ "github.com/lib/pq"
)

type PostgresConf struct {
	User     string
	Password string
	Database string
	Url      string
	Port     string
}

func ProvidePostgres(cfg PostgresConf) db.Db {
	d, err := sql.Open("postgres", conStr(cfg))

	if err != nil {
		panic(err)
	}

	return db.NewDb(d)
}

func conStr(cfg PostgresConf) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Url, cfg.Port, cfg.Database)
}
