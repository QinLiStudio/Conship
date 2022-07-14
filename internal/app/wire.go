package app

import (
	"github.com/QinLiStudio/Conship/internal/app/router"
	"github.com/google/wire"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		InitGinEngine,
		InjectorSet,
		router.RouterSet,
	)
	return new(Injector), nil, nil
}
