package di

import (
	"sync"
)

var once sync.Once
var instance *App

type App struct {
	Database *Database
}

func DI() *App {
	once.Do(func() {
		instance = &App{}
	})

	return instance
}

func (app *App) Bootstrap() {
	DI().Database = &Database{DB: nil}
}
