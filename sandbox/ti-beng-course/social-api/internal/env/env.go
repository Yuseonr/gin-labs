package env

// bisa pake library env eksternal, tapi ini cara ngelakuin manual

import (
	"os"
	"strconv"
)

// Get string value enviroment variable dari native golang
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

// Get Int value enviroment variable
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return  fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}