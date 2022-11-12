package controllers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/cumhurmesuttokalak/goweb/admin/helpers"
	"github.com/cumhurmesuttokalak/goweb/admin/models"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmp, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAllPosts()
	tmp.ExecuteTemplate(w, "index", data)
}
func (dashboad Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmp, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
	}
	tmp.ExecuteTemplate(w, "index", nil)
}

// formdan gelen verileri db ye kaydetme
func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	//upload
	r.ParseMultipartForm(10)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
	}
	//end upload
	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Picture_url: "uploads/" + header.Filename,
	}.CreatePost()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmp, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.GetSinglePost(params.ByName("id"))
	tmp.ExecuteTemplate(w, "index", data)
}
func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := models.Post{}.GetSinglePost(params.ByName("id"))
	post.DeletePost()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := models.Post{}.GetSinglePost(params.ByName("id"))
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")
	var picture_url string
	if is_selected == "1" {
		//Upload
		r.ParseMultipartForm(10)
		file, header, err := r.FormFile("blog-picture")
		if err != nil {
			fmt.Println(err)
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
		}
		//End upload
		picture_url = "uploads/" + header.Filename
		os.Remove(post.Picture_url)
	} else {
		picture_url = post.Picture_url
	}
	post.UpdatesPost(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
