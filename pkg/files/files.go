package files

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	Bucket    string `hcl:"bucket"`
	Directory string `hcl:"directory"`
}

type Files struct {
	Logging *zap.Logger
	Config  Config
	S3      *minio.Client
}

type Opts struct {
	fx.In

	Logging *zap.Logger
	Config  Config
	S3      *minio.Client
}

func New(o Opts) *Files {
	return &Files{
		Logging: o.Logging.Named("files"),
		Config:  o.Config,
		S3:      o.S3,
	}
}

func (f *Files) PutFile(ctx context.Context, filename string, u uuid.UUID) (*url.URL, error) {
	f.Logging.Debug("write file to bucket", zap.String("filename", filename), zap.String("uuid", u.String()))

	presigned, err := f.S3.PresignedPutObject(ctx, f.Config.Bucket, f.ObjectName(u, filename), 1*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("generate presigned url: %w", err)
	}

	return presigned, nil
}

func (f *Files) ObjectName(u uuid.UUID, filename string) string {
	return strings.Join([]string{f.Config.Directory, "/", u.String(), "/", filename}, "")
}
