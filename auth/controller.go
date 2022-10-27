package auth

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Controller struct {
	log *logrus.Logger
	db  gorm.DB
}

func NewController(log *logrus.Logger, db gorm.DB) *Controller {
	return &Controller{
		log: log,
		db:  db,
	}
}
