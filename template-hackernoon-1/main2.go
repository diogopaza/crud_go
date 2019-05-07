package main

import(

	"html/template"
	"net/http"

)


var tmpl = `<html>
<head>
	<title>First Program</title>
</head>
<body>
	{{ . }}
</body>


</html>`

func handler(w http.ResponseWriter, r *http.Request){
	
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	t.Execute(w, "Hello")

}

func main(){

	http.HandleFunc("/view", handler)
	http.ListenAndServe(":9000", nil)

}


