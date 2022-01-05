package creds

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	creds            *auth = nil
	CONFIG_NOT_FOUND error = errors.New("File \"auth.json\" not found in current directory.")
	CONFIG_NOT_INIT  error = errors.New("Config file is not yet initiated.")
)

type auth struct {
	Token  string `json:"token"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func Init(path string) error {
	fnPath := "auth.json"
	if path != "" || strings.TrimSpace(path) != "" {
		fnPath = path
	}

	file, err := os.Open(fnPath)
	if err != nil {
		return CONFIG_NOT_FOUND
	}

	log.Println("Successfully initialized authorization credentials.")

	var auth auth
	byteV, _ := ioutil.ReadAll(file)

	if jErr := json.Unmarshal(byteV, &auth); jErr != nil {
		return errors.New("File is not json.")
	}

	creds = &auth

	defer file.Close()

	return nil
}

// Uninitializes creds file. Hacky way of testing
// but doesn't really hurt that much in actual code.
func UnInit() {
	creds = nil
}

func Token() string {
	panicIfNotInit()

	return creds.Token
}

func Key() string {
	panicIfNotInit()

	return creds.Key
}

func Secret() string {
	panicIfNotInit()

	return creds.Secret
}

func panicIfNotInit() {
	if creds == nil || creds.Token == "" {
		panic(CONFIG_NOT_INIT.Error())
	}
}
