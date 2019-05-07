package main

import(

	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request){

	t, _ := template.ParseFiles("view.html")
	t.Execute(w, "Hello World")

}

func main(){

	http.HandleFunc("/view", handler)
	http.ListenAndServe(":9000", nil)

}