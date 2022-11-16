package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/cumhurmesuttokalak/goweb/admin/helpers"
	"github.com/cumhurmesuttokalak/goweb/admin/models"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Category struct{}

func (category Category) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	tmp, err := template.ParseFiles(helpers.Include("categories/list")...)
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAllCategories()
	data["Alert"] = helpers.GetAlert(w, r)
	tmp.ExecuteTemplate(w, "index", data)
}
func (category Category) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	categoryTitle := r.FormValue("category-title")
	categorySlug := slug.Make(categoryTitle)

	models.Category{Title: categoryTitle, Slug: categorySlug}.CreateCategory()
	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi")
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}
func (category Category) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	kategori := models.Category{}.GetSingleCategory(params.ByName("id"))
	kategori.DeleteCategory()
	helpers.SetAlert(w, r, "Kayıt Silindi")
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}
