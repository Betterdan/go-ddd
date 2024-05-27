package service

import (
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	container.Provide(NewUserAppService)
	container.Provide(NewOrderAppService)
}
