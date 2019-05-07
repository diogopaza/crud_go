package main

import(
	"fmt"
	"net/http"
	"html/template"
)




type Pessoa struct{

	Nome string
	Idade int
}

type TodasPessoas struct{

	Pessoas []Pessoa
}

func main(){

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func( w http.ResponseWriter, r *http.Request){
		
		
		 data := TodasPessoas{
			 Pessoas: []Pessoa{
				 {Nome:"Diogo",Idade:30},
				 {Nome:"Beto", Idade:40},
			 },
		 }		
			
		
		//http.ServeFile(w,r,"index.html")
		fmt.Println(data)
		tmpl.Execute(w, data)

	})

	http.ListenAndServe(":8000", nil)
}