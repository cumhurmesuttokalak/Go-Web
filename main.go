package main

import (
	"net/http"

	admin_models "github.com/cumhurmesuttokalak/goweb/admin/models"
	"github.com/cumhurmesuttokalak/goweb/config"
)

func main() {
	admin_models.DBConn()
	http.ListenAndServe(":8080", config.Routes())
}
