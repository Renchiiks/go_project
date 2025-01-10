package main

import (
	"github.com/Renchiiks/go_project/internal/app"
	"github.com/Renchiiks/go_project/internal/repository/carrepo"
	//"log"

	//"github.com/user1/mycoollog"

	"log/slog"
)

func main() {
	//log := mycoollo.NewWithName("general_log")
	conf := initConfig()

	carsRepo, err := carrepo.New(conf.Host, conf.Port)
	if err != nil {
		slog.Error("Could not initialize car repo", "error", err)
	}

	application := app.New(carsRepo)
	application.Run()
}
