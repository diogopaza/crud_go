package main

import(

	"fmt"
	"html/template"
	"os"
)


var defineDemo = `
	{{ define "a" }} Template A {{ end }}
	{{define "b" }} Template B {{ end }}
	`

func main(){

	var err error

	t:=template.New("defineActionDemo")
	t, err = t.Parse(defineDemo)
	if err != nil{
		fmt.Println("parsing failed:", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "a", nil)
	if err != nil{
		fmt.Println("execution failed:", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "b", nil)
	if err != nil{
		fmt.Println("execution failed:", err)
	}






}