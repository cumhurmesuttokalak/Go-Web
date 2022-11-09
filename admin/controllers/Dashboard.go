package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/cumhurmesuttokalak/goweb/admin/helpers"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmp, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
	}
	tmp.ExecuteTemplate(w, "index", nil)
}
