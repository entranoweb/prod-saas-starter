module github.com/moasq/go-b2b-starter/pkg/ocr

go 1.25


require (
	github.com/moasq/go-b2b-starter/pkg/logger v0.0.0
	go.uber.org/dig v1.19.0
)

replace github.com/moasq/go-b2b-starter/pkg/logger => ../logger
