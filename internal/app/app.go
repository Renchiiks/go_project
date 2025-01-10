package app

type App struct {
	repo Repo
}

type Repo interface {
	UpdateCars() error
	DeleteCars() error
}

func New() *App {
	app := &App{
		repo: repo,
	}
	return app
}

func (app *App) Run() {

}
