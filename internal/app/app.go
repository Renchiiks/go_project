package app

import "github.com/Renchiiks/go_project/internal/repository/carrepo"

type App struct {
	repo carrepo.Car
}

type Repo interface {
	UpdateCars() error
	DeleteCars() error
}

func New() *App {

	app := &App{
		repo: carrepo.Car{},
	}
	return app
}

func (app *App) Run() {

}
