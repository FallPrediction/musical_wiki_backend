package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var Bucket = os.Getenv("AWS_BUCKET")
var Endpoint = os.Getenv("AWS_ENDPOINT")
