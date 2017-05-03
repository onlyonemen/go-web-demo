package router

import (
	"demo/controllers"
	"net/http"
)

func init() {
	//public
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//controllers
	http.HandleFunc("/", controllers.IndexController)
	http.HandleFunc("/json", controllers.JsonController)
	http.HandleFunc("/file/add", controllers.FileAddController)
	http.HandleFunc("/form", controllers.FormController)
	http.HandleFunc("/form/multi", controllers.FormMultiController)

}
