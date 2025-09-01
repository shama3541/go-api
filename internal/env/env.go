package env

import (
	"os"
	"strconv"
)

func Getstring(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valAsInt, _ := strconv.Atoi(val)

	return valAsInt
}
