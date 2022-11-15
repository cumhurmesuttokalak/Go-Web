package config

import (
	"net/http"

	admin "github.com/cumhurmesuttokalak/goweb/admin/controllers"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()

	//Admin routes
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//Userops routes
	r.GET("/admin/login", admin.User{}.Index)
	r.POST("/admin/do_login", admin.User{}.Login)
	r.GET("/admin/logout", admin.User{}.Logout)

	//ServeFiles
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
