package config

import (
	"net/http"

	admin "github.com/cumhurmesuttokalak/goweb/admin/controllers"
	site "github.com/cumhurmesuttokalak/goweb/site/controllers"
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

	//Categories routes
	r.GET("/admin/categories", admin.Category{}.Index)
	r.POST("/admin/categories/add", admin.Category{}.Add)
	r.GET("/admin/categories/delete/:id", admin.Category{}.Delete)

	//SITE
	//Homepage
	r.GET("/", site.Homepage{}.Index)
	r.GET("/yazilar/:slug", site.Homepage{}.Detail)

	//ServeFiles
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/site/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
