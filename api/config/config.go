package config

import (
	"os"
	"strings"
)

var GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
var PORT = os.Getenv("PORT")
var ALLOWED_EXTENSION_CLIENT_IDS = getEnvAsSlice("ALLOWED_EXTENSION_CLIENT_IDS")

func getEnvAsSlice(env string) []string {
	str := os.Getenv(env)
	return strings.Split(str, ",")
}
