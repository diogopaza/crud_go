package main

import(

	"fmt"
	"html/template"
	"os"
)


var defineDemo = 
	{{ define "a" }} Template A {{ end }}
	{{define b }} Template B {{ end }}