package config

import (
	"fmt"
	"os"
)

type pg struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (pg pg) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", pg.host, pg.user, pg.password, pg.dbname, pg.port)
}

func (pg pg) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pg.user, pg.password, pg.host, pg.port, pg.dbname)
}

func NewPg() pg {
	return pg{
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	}
}
