package controllers

import (
	"html/template"
	"net/http"
)

type CommonResponse struct {
	Msg string
}

func IndexController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html", "view/temp/body.temp", "view/temp/head.temp")
	if err != nil {
		panic(err)
	}
	//1.
	err = t.Execute(w, map[string]string{"Msg": "Hello World"})
	//2.
	// err = t.Execute(w, &CommonResponse{"Hello World"})
	if err != nil {
		panic(err)
	}
}
