package config

import (
	"fmt"

	"github.com/tvanriel/cloudsdk/hclconfig"
	"github.com/tvanriel/cloudsdk/http"
	"github.com/tvanriel/cloudsdk/logging"
	"github.com/tvanriel/cloudsdk/s3"
	"github.com/tvanriel/filehost/pkg/files"
	"github.com/tvanriel/filehost/pkg/web"
	"go.uber.org/fx"
)

type Configuration struct {
	fx.Out

	HTTP    http.Configuration    `hcl:"http,block"`
	Logging logging.Configuration `hcl:"logging,block"`
	S3      *s3.Configuration     `hcl:"s3,block"`
	Web     web.Config            `hcl:"web,block"`
	Files   files.Config          `hcl:"files,block"`
}

func HclConfiguration() (Configuration, error) {
	config := Configuration{}

	err := hclconfig.HclConfiguration(&config, "filehost")
	if err != nil {
		return Configuration{}, fmt.Errorf("parse configuration: %w", err)
	}

	return config, nil
}
