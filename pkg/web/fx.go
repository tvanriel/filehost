package web

import (
	"github.com/tvanriel/cloudsdk/http"
	"go.uber.org/fx"
)

//nolint:gochecknoglobals // Module.
var Module = fx.Module("web", fx.Provide(
	http.AsRouteGroup(New),
))
