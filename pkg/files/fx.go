package files

import "go.uber.org/fx"

//nolint:gochecknoglobals // Module.
var Module = fx.Module("files", fx.Provide(New))
