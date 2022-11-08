package main

import (
	"net/http"

	"github.com/cumhurmesuttokalak/goweb/config"
)

func main() {
	http.ListenAndServe(":8080", config.Routes())
}
