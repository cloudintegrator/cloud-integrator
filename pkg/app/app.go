package app

import (
	"fmt"
	"github.com/cloudintegrator/cloud-integrator/pkg/app/spec"
)

type App struct {
	AppSpec *spec.AppSpec
}

func (app *App) NewApp() {
	if len(app.AppSpec.Listener.Type) != 0 {
		fmt.Printf("Type:%s", app.AppSpec.Listener.Type)
	} else {
		panic("Error")
	}

}
