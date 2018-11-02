package templates

// configTmpl : template of config
var configTmpl = `{
    "DBConnectionString"    :   "{{.DBConnectionString}}",
    "DBEngine"              :   "{{.DBEngine}}",
    "WebAddress"            :   "{{.WebAddress}}",
    "WebPort"               :   {{.WebPort}}
}
`
