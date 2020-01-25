package templates

// configTmpl : template of config
var configTmpl = `
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
	DBConnectionString string {{backQuote}}json:"DB_CONNECTION_STRING" envconfig:"DB_CONNECTION_STRING"{{backQuote}}
	DBEngine           string {{backQuote}}json:"DB_ENGINE" envconfig:"DB_ENGINE"{{backQuote}}
	WebAddress         string {{backQuote}}json:"API_ADDRESS" envconfig:"API_ADDRESS"{{backQuote}}
	WebPort            int    {{backQuote}}json:"API_PORT" envconfig:"API_PORT"{{backQuote}}
}

// GetConfigs : Get application configuration
func GetConfigs() Config {
	readConfigOnce.Do(func() {
		err := configuration.JSON("config.json", &c)
		if err != nil {
			fmt.Println(err)
			// get configuration from environment variables
			err := envconfig.Process("", &c)
			if err != nil {
				panic(fmt.Sprintf("Error while initiating app configuration : %v", err))
			}
			if c.DBConnectionString == "" {
				panic("No Database Connectionstring found")
			}
		}
	})
	return c
}

`
