package main

import (
	"net/http"

	admin_models "github.com/cumhurmesuttokalak/goweb/admin/models"
	"github.com/cumhurmesuttokalak/goweb/config"
)

func main() {
	admin_models.DBConn()
	admin_models.UpdatesPost(admin_models.Post{Title: "YepyeniTitle", Slug: "YepyeniSlug"})
	http.ListenAndServe(":8080", config.Routes())
}
