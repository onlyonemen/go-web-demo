package controllers

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func FormController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/form.html")
	if err != nil {
		panic(err)
	}
	//1.
	err = t.Execute(w, "")
	//2.
	// err = t.Execute(w, &CommonResponse{"Hello World"})
	if err != nil {
		panic(err)
	}
}

func FormMultiController(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := r.FormValue("id")
	file, file_header, err := r.FormFile("file")
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("name:%s,id:%s,file_name:%s\n", name, id, file_header.Filename)

	if saveFile("files/"+file_header.Filename, file) != nil {
		w.Write([]byte("failed"))
		return
	}

	w.Write([]byte("success"))
}

func saveFile(fileNmae string, file multipart.File) error {
	defer file.Close()
	f, err := os.Create(fileNmae)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, file)

	return err

}
