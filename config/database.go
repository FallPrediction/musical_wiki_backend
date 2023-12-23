package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var Pgsql = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
