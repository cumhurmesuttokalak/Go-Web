package controllers

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"

	"github.com/cumhurmesuttokalak/goweb/admin/helpers"
	"github.com/cumhurmesuttokalak/goweb/admin/models"
	"github.com/julienschmidt/httprouter"
)

type User struct{}

func (userops User) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmp, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	tmp.ExecuteTemplate(w, "index", data)
}
func (userops User) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	user := models.User{}.GetSingleUser("username = ? AND password= ?", username, password)
	if user.Username == username && user.Password == password {
		helpers.SetUser(w, r, username, password)
		helpers.SetAlert(w, r, "Hoşgeldiniz")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		helpers.SetAlert(w, r, "Yanlış Kullanıcı Adı veya Şifre")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}
func (user User) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Hoşçakalın")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
