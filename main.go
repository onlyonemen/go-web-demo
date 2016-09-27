package main

import (
	_ "demo/router"
	"net/http"
)

const SERVER_PORT string = ":8088"

func main() {
	setPort()
}
func setPort() {
	println("start server at" + SERVER_PORT)
	//port
	e := http.ListenAndServe(SERVER_PORT, nil)
	if e != nil {
		panic(e)
	}

}
