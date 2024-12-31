package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/organization/handler"
	"github.com/leagueify/organization/internal/config"
	"github.com/leagueify/organization/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	organizationRouter := handler.OrganizationRouter()

	router.Handle("/organization/", http.StripPrefix("/organization", organizationRouter))

	// define server config
	organizationServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := organizationServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
