package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmp, err := template.ParseFiles("admin/views/dashboard/list/index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmp.Execute(w, nil)
}
