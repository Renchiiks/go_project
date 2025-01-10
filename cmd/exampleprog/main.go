package main

import (
	"github.com/user/exampleprog/internal/app"
	"github.com/user/exampleprog/internal/reposotory/carrepo"
	"github.com/user1/mycoollog"
)

func main() {
	//log := mycoollo.NewWithName("general_log")
	conf := initConfig()

	carsRepo, err := carrepo.New(conf.Host, conf.Port)
	if err != nil {
		log
	}

	application := app.New(carsRepo)
	application.Run()
}
