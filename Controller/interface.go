package Controller

import logger "usr-service/Logger"

type ControllerInterface interface {
	UserInterface
}

type Controller struct {
	ControllerInterface
	Log *logger.Logger
}
