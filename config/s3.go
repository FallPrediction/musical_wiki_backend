package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Bucket = os.Getenv("AWS_BUCKET")
var Endpoint = os.Getenv("AWS_ENDPOINT")
var Region = os.Getenv("AWS_REGION")
