module github.com/northvolt/kin-openapi

go 1.16

require (
	github.com/getkin/kin-openapi v0.94.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/jsonpointer v0.19.5
	github.com/gorilla/mux v1.8.0
	github.com/mitchellh/copystructure v1.2.0
	github.com/stretchr/testify v1.5.1
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/getkin/kin-openapi => ./
