module github.com/moasq/go-b2b-starter/api/example_resource

go 1.25

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/moasq/go-b2b-starter/app/example_resource v0.0.0-00010101000000-000000000000
	github.com/moasq/go-b2b-starter/pkg/auth v0.0.0-00010101000000-000000000000
	github.com/moasq/go-b2b-starter/pkg/common v0.0.0-00010101000000-000000000000
	github.com/moasq/go-b2b-starter/server v0.0.0-00010101000000-000000000000
	go.uber.org/dig v1.18.0
)

replace github.com/moasq/go-b2b-starter/app/example_resource => ../../app/example_resource

replace github.com/moasq/go-b2b-starter/pkg/auth => ../../pkg/auth

replace github.com/moasq/go-b2b-starter/pkg/common => ../../pkg/common

replace github.com/moasq/go-b2b-starter/server => ../../server
