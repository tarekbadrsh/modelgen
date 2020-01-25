package templates

// configjsonTmpl : template of config json file
var configjsonTmpl = `{
    "DB_CONNECTION_STRING"      :   "{{.DBConnectionString}}",
    "DB_ENGINE"                 :   "{{.DBEngine}}",
    "API_ADDRESS"               :   "{{.WebAddress}}",
    "API_PORT"                  :   {{.WebPort}}
}
`
