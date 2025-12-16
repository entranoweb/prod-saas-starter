package cmd

import (
	"go.uber.org/dig"
	
	"github.com/moasq/go-b2b-starter/pkg/eventbus"
	"github.com/moasq/go-b2b-starter/pkg/logger/domain"
)

// ProvideEventBus creates and configures the event bus with middleware
func ProvideEventBus(container *dig.Container) error {
	return container.Provide(func(logger domain.Logger) eventbus.EventBus {
		middleware := []eventbus.EventMiddleware{
			eventbus.RecoveryMiddleware(logger),
			eventbus.LoggingMiddleware(logger),
			eventbus.MetricsMiddleware(),
		}
		
		return eventbus.NewInMemoryEventBus(middleware...)
	})
}