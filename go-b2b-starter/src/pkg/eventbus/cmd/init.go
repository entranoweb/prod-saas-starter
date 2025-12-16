package cmd

import "go.uber.org/dig"

// Init initializes the event bus dependencies
func Init(container *dig.Container) error {
	if err := ProvideEventBus(container); err != nil {
		return err
	}
	
	return nil
}