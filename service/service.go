package service

import "github.com/dpatsora/go-todo/app"

func NewApplication() (app.Application, func()) {
	return app.Application{}, func(){ return }
}