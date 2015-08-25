package config

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/joho/godotenv"
)

// ErrEnvVarNotFound is an error that is raised when an environment variable is missing.
type ErrEnvVarNotFound string

func (envVar ErrEnvVarNotFound) Error() string {
	return fmt.Sprintf("%s was not found in the environment variables", string(envVar))
}

// ErrUnableToParseIntWithDefault is raises when converting a environment variable to int raises an error
type ErrUnableToParseIntWithDefault struct {
	key    string
	raw    string
	defVal int
}

func (e ErrUnableToParseIntWithDefault) Error() string {
	return fmt.Sprintf(
		"unable to parse .env variable '%s' with value '%s' as integer, setting to default '%d'",
		e.key,
		e.raw,
		e.defVal,
	)
}

// ErrUnableToParseInt is raises when converting a environment variable to int raises an error
type ErrUnableToParseInt struct {
	key string
	raw string
}

func (e ErrUnableToParseInt) Error() string {
	return fmt.Sprintf(
		"unable to parse .env variable '%s' with value '%s' as integer",
		e.key,
		e.raw,
	)
}

// InitEnv initializes the environment variables.
func InitEnv() {
	if IsTest() {
		return
	}

	// Load the environment variables first from `.env.default` and then from `.env`, allowing `.env` to override when
	// necessary.
	err := godotenv.Load(".env", ".env.default")
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Environment: %s", Env())
}

// Get simply returns the environment variable as a string, or an empty string when undefined.
func Get(key string) string {
	return os.Getenv(key)
}

// GetString returns the environment variable as a string, or the default value when undefined.
func GetString(key, defVal string) string {
	val := Get(key)
	if val == "" {
		return defVal
	}
	return val
}

// MustGetString returns the environment variable as a string, or logs a fatal error when undefined.
func MustGetString(key string) string {
	val := Get(key)
	if val == "" {
		log.Fatal(ErrEnvVarNotFound(key))
	}
	return val
}

// GetInt returns the environment variable as a int, or the default value when undefined.
func GetInt(key string, defVal int) int {
	raw := os.Getenv(key)
	if raw == "" {
		return defVal
	}
	val, err := strconv.Atoi(raw)
	if err != nil {
		log.Warn(
			ErrUnableToParseIntWithDefault{
				key:    key,
				raw:    raw,
				defVal: defVal,
			},
		)
	}
	return val
}

// MustGetInt returns the environment variable as a string, or logs a fatal error when undefined.
func MustGetInt(key string) int {
	raw := os.Getenv(key)
	if raw == "" {
		log.Fatal(ErrEnvVarNotFound(key))
	}
	val, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatal(
			ErrUnableToParseInt{
				key: key,
				raw: raw,
			},
		)
	}
	return val
}
