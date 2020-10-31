package di

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
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

func (app *App) SetupCloseHandler() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig

		// еще надо бы дождаться обработки открытых соединений
		app.Database.Close()

		os.Exit(0)
	}()
}
