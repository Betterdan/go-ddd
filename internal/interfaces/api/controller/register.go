package api

import (
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	container.Provide(NewUserController)
	container.Provide(NewOrderController)
	container.Provide(NewRouter)
	container.Provide(NewServer)
}
