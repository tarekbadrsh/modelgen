package server_test

type config struct {
	DBConnectionString string `json:"DB_CONNECTION_STRING" envconfig:"DB_CONNECTION_STRING"`
	DBEngine           string `json:"DB_ENGINE" envconfig:"DB_ENGINE"`
	WebAddress         string `json:"API_ADDRESS" envconfig:"API_ADDRESS"`
	WebPort            int    `json:"API_PORT" envconfig:"API_PORT"`
}