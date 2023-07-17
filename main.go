package main

import (
	"github.com/IvanLMItaca/max-inventory/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(),
	)
	app.Run()
}
