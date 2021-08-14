package controllers

import (
	"github.com/olingern/golytics/pkg/db/clients"
)

type Controller struct {
	db *clients.SqliteClient
}

func NewController(client *clients.SqliteClient) *Controller {
	return &Controller{
		db: client,
	}
}
