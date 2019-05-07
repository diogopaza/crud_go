package main

import(
	
	"html/template"
	"fmt"
	"net/http"
	

	
)


type Todo struct{
	Title string
	Done  bool
}


type TodoPageData struct{

	PageTitle string
	Todos []Todo
}


func main(){	

	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request){

		data := TodoPageData{
			PageTitle: "Diogo Paza with Golang",
			Todos: []Todo{
				{ Title:"ffffff", Done: true},
				{ Title:"22222222", Done: true},
				{ Title:"333333", Done: false},
			},
		 }
		 
		 tmpl.Execute(response, data)
		 fmt.Println(data)
		
	})
	
	http.ListenAndServe(":8000", nil)		
	

}