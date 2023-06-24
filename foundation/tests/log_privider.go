package main

import (
	"github.com/go-fires/fires/foundation"
	"log"
)

type Log struct{}

func (l *Log) Report(v interface{}) {
	log.Println(v)
}

type LogProvider struct {
	app *Application
	*foundation.UnimplementedProvider
}

func NewLogProvider(app *Application) *LogProvider {
	return &LogProvider{
		app: app,
	}
}

func (p *LogProvider) Register() {
	p.app.Log = &Log{}
}
