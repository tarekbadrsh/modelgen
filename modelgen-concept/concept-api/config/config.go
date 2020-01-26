
package config

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/tarekbadrshalaan/goStuff/configuration"
)

var readConfigOnce sync.Once

var c Config

// Config : application configuration
type Config struct {
	DBConnectionString string `json:"DB_CONNECTION_STRING" envconfig:"DB_CONNECTION_STRING"`
	DBEngine           string `json:"DB_ENGINE" envconfig:"DB_ENGINE"`
	WebAddress         string `json:"API_ADDRESS" envconfig:"API_ADDRESS"`
	WebPort            int    `json:"API_PORT" envconfig:"API_PORT"`
}

// GetConfigs : Get application configuration
func GetConfigs() (c Config, err error) {
	readConfigOnce.Do(func() {
		jsonerr := configuration.JSON("config.json", &c)
		if jsonerr != nil {
			fmt.Println(jsonerr)
			// get configuration from environment variables
			err = envconfig.Process("", &c)
			if err != nil {
				err = fmt.Errorf("Error while initiating app configuration : %v", err)
			}
		}
		if c.DBConnectionString == "" {
			err = fmt.Errorf("No Database Connectionstring found")
		}
	})
	return c, err
}

