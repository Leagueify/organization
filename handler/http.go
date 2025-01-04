package handler

import (
	"embed"
	"net/http"

	"github.com/leagueify/organization/internal/config"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed docs/*
	doc   embed.FS
	docFS = echo.MustSubFS(doc, "docs")
)

type httpHandler struct {
	app *echo.Echo
	cfg *config.Config
}

func (h *httpHandler) Initialize() {
	group := h.app.Group("/organization")

	group.StaticFS("/docs", docFS)
	group.GET("/healthz", healthz)
}

func HTTP(app *echo.Echo, cfg *config.Config) Handler {
	return &httpHandler{
		app: app,
		cfg: cfg,
	}
}

func healthz(ctx echo.Context) error {
	type healthResponse struct {
		Service string `json:"service"`
		Status  string `json:"status"`
	}

	return ctx.JSON(http.StatusOK, healthResponse{
		Service: "organization",
		Status:  http.StatusText(http.StatusOK),
	})
}
