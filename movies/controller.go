package movie

import (
	"github.com/sirupsen/logrus"
)

type Controller struct {
	log *logrus.Logger
}

func NewController(log *logrus.Logger) *Controller {
	return &Controller{log: log}
}
