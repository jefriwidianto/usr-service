package Routes

import (
	"fmt"
	"github.com/labstack/echo"
	"usr-service/Controller"
	logger "usr-service/Logger"
)

type Routes struct {
	Controller Controller.Controller
	Log        *logger.Logger
}

func (app *Routes) CollectRoutes(e *echo.Echo, host, port string) {
	appRoutes := e

	appRoutes.POST("/", app.Controller.Login)

	app.Log.Fatal(appRoutes.Start(fmt.Sprintf("%s:%s", host, port)).Error())
}
