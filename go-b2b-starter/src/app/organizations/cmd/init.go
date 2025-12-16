package cmd

import (
	"go.uber.org/dig"

	"github.com/moasq/go-b2b-starter/app/organizations"
)

// Init initializes the organizations module
func Init(container *dig.Container) error {
	module := organizations.NewModule(container)
	return module.RegisterDependencies()
}