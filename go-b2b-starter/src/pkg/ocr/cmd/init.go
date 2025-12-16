package cmd

import (
	"go.uber.org/dig"

	"github.com/moasq/go-b2b-starter/pkg/ocr/domain"
	"github.com/moasq/go-b2b-starter/pkg/ocr/infra"
	loggerDomain "github.com/moasq/go-b2b-starter/pkg/logger/domain"
)

func Init(container *dig.Container) error {
	return container.Provide(func(logger loggerDomain.Logger) (domain.OCRService, error) {
		config := infra.NewOCRConfig()
		return infra.NewMistralOCRClient(config, logger)
	})
}