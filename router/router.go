package router

import (
	"demo/controllers"
	"net/http"
)

func init() {
	//public
	http.Handle("/public/", http.FileServer(http.Dir("public")))
	//controllers
	http.HandleFunc("/", controllers.IndexController)
	http.HandleFunc("/json", controllers.JsonController)
}
