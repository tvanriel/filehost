package web

import (
	"fmt"
	nethttp "net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tvanriel/cloudsdk/http"
	"github.com/tvanriel/filehost/pkg/files"
	"github.com/tvanriel/filehost/pkg/views"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	S3PublicURLPrefix string `hcl:"s3_public_url_prefix"`
}
type Opts struct {
	fx.In

	Config  Config
	Logging *zap.Logger
	F       *files.Files
}
type Web struct {
	Config Config
	Logger *zap.Logger
	F      *files.Files
}

// ApiGroup implements [http.RouteGroup].
//
//nolint:stylecheck // implements interface.
func (w *Web) ApiGroup() string {
	return ""
}

// Handler implements http.RouteGroup.
func (w *Web) Handler(g *echo.Group) {
	g.POST("sign", w.SignHandler)
	g.GET("", w.IndexHandler)
}

type SignRequest struct {
	Filename string `json:"filename"`
}

type SignResponse struct {
	PresignedURL string `json:"presignedUrl"`
	FileURL      string `json:"fileUrl"`
}

func (w *Web) SignHandler(ctx echo.Context) error {
	req := SignRequest{}

	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	fileUUID := uuid.Must(uuid.NewRandom())

	presignedURL, err := w.F.PutFile(ctx.Request().Context(), req.Filename, fileUUID)
	if err != nil {
		w.Logger.Error("generate presign", zap.Error(err))

		return fmt.Errorf("putting file: %w", err)
	}

	err = ctx.JSON(nethttp.StatusOK, SignResponse{
		PresignedURL: presignedURL.String(),
		FileURL: strings.Join([]string{
			w.Config.S3PublicURLPrefix,
			"/",
			fileUUID.String(),
			"/",
			req.Filename,
		}, ""),
	})
	if err != nil {
		return fmt.Errorf("write json response: %w", err)
	}

	return nil
}

func (w *Web) IndexHandler(ctx echo.Context) error {
	err := views.Layout().Render(ctx.Request().Context(), ctx.Response().Writer)
	if err != nil {
		return fmt.Errorf("render index template: %w", err)
	}

	return nil
}

// Version implements http.RouteGroup.
func (w *Web) Version() string {
	return ""
}

func New(o Opts) *Web {
	return &Web{
		Config: o.Config,
		Logger: o.Logging.Named("web"),
		F:      o.F,
	}
}

var _ http.RouteGroup = (*Web)(nil)
