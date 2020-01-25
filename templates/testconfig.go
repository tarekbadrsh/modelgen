package templates

// testConfigTmpl : template of Test Config.
var testConfigTmpl = `package server_test

type config struct {
	DBConnectionString string {{backQuote}}json:"DB_CONNECTION_STRING" envconfig:"DB_CONNECTION_STRING"{{backQuote}}
	DBEngine           string {{backQuote}}json:"DB_ENGINE" envconfig:"DB_ENGINE"{{backQuote}}
	WebAddress         string {{backQuote}}json:"API_ADDRESS" envconfig:"API_ADDRESS"{{backQuote}}
	WebPort            int    {{backQuote}}json:"API_PORT" envconfig:"API_PORT"{{backQuote}}
}`
