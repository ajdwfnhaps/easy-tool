package app

import (
	"context"
)

//App Context
type App struct {
	context.Context
}

//New 创建应用
func New() *App {
	return &App{}
}
