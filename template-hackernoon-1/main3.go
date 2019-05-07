package main

import(

	"html/template"
	"net/http"

)


func handler1(w http.ResponseWriter, r *http.Request){
	
	t, _ := template.ParseFiles("t1.html","t2.html")
	t.Execute(w, "Asit")

}

func handler2(w http.ResponseWriter, r *http.Request){
	
	t, _ := template.ParseFiles("t1.html","t2.html")
	t.ExecuteTemplate(w, "t2.html", "Golang")

}

func main(){

	http.HandleFunc("/t1", handler1)
	http.HandleFunc("/t2", handler2)

	http.ListenAndServe(":9000", nil)
}