package common

import (
	"encoding/json"
	"os"
	"time"
)

//Configuration provides the configuration model
type Configuration struct {
	Host, Database, Username, Password, BasicMember string
	MaximumFailedLogin                              int
	DropDB                                          bool `json:"DropDB,string"`
	JWTKey, JWTIssuer, ENV                          string
	JWTLife                                         time.Duration
}

//LoadConfiguration loads configuration from json file
func (config *Configuration) LoadConfiguration(filepath string) (err error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
	}
	return err
}
