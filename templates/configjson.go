package templates

// configjsonTmpl : template of config json file
var configjsonTmpl = `{
    "DBConnectionString"    :   "{{.DBConnectionString}}",
    "DBEngine"              :   "{{.DBEngine}}",
    "WebAddress"            :   "{{.WebAddress}}",
    "WebPort"               :   {{.WebPort}}
}
`
