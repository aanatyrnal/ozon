package main

import (
	"os"

	"ozon/config"
	"ozon/internal/server"
	"ozon/pkg/db/postgres"
	"ozon/pkg/utils"
)

func main() {
	println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	println(utils.GetConfigPath(os.Getenv("config")))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		println("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		println("ParseConfig: %v", err)
	}

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		println("Postgresql init: %s", err)
	}
	defer psqlDB.Close()

	s := server.NewServer(cfg, psqlDB)

	if err = s.Run(); err != nil {
		println(err.Error())
	}
}
