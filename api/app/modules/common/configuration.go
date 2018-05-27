package common

import (
	"encoding/json"
	"os"
	"time"
)

//Configuration provides the configuration model
type Configuration struct {
	JWTKey, JWTIssuer, AuthHeader string
	JWTLife                       time.Duration
}

//LoadConfiguration loads configuration from json file
func (config *Configuration) LoadConfiguration(filename string) (err error) {
	file, err := os.Open("config/" + filename)
	defer file.Close()
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
	}
	return err

}
