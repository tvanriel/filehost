package app

import (
	"github.com/tvanriel/cloudsdk/http"
	"github.com/tvanriel/cloudsdk/logging"
	"github.com/tvanriel/cloudsdk/s3"
	"github.com/tvanriel/filehost/pkg/config"
	"github.com/tvanriel/filehost/pkg/files"
	"github.com/tvanriel/filehost/pkg/web"
	"go.uber.org/fx"
)

func Run() {
	app := fx.New(
		logging.Module,
		s3.Module,
		http.Module,
		web.Module,
		files.Module,
		fx.Provide(config.HclConfiguration),
	)

	app.Run()
}
