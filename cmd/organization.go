package main

import (
	"github.com/leagueify/organization/internal/config"
	"github.com/leagueify/organization/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewEchoServer(cfg).Start()
}
