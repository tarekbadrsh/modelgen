package templates

// testConfigTmpl : template of Test Config.
var testConfigTmpl = `package server_test

type config struct {
	DBConnectionString string
	DBEngine           string
	WebAddress         string
	WebPort            int
}`
