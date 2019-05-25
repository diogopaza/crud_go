package main


import "net/http"
import "text/template"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
	"encoding/json"
	"fmt"
	"os"
)


var tmpl = template.Must(template.ParseGlob("tmpl/*"))

type Names struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}


func dbConn()(db *sql.DB){

	db, err := sql.Open("mysql", "cursos09:lp231210@tcp(mysql13-farm70.uni5.net)/cursos09")
	if err != nil{
		fmt.Println("Erro: ", err)
		os.Exit(1)
		
	}

	return db
}

func Index(w http.ResponseWriter, r *http.Request){

	var id int
		var name string
		var email string
		n := Names{}
 		res := []Names{}
			
		db := dbConn()	
		
		selDB, err := db.Query("SELECT * FROM names ")
		if err != nil{
			fmt.Println("Erro no banco")
			os.Exit(1)
		}	

		for selDB.Next(){	
		
			err = selDB.Scan(&id, &name, &email)
			if err != nil{
				fmt.Println("Erro percorrer Names")
				os.Exit(1)
		}
		
		n.Id = id
		n.Name = name
		n.Email = email

		res = append(res, n)
		

	}
	
		js, _ := json.Marshal(res) 
		fmt.Println(string(js))
		//fmt.Println(res)
		tmpl.ExecuteTemplate(w,"Index.tmpl",res)
	
}

func jsonPage(w http.ResponseWriter, req *http.Request){

	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	var id int
		var name string
		var email string
		n := Names{}
 		res := []Names{}
			
		db := dbConn()	
		
		selDB, err := db.Query("SELECT * FROM names ")
		if err != nil{
			fmt.Println("Erro no banco")
			os.Exit(1)
		}	

		for selDB.Next(){	
		
			err = selDB.Scan(&id, &name, &email)
			if err != nil{
				fmt.Println("Erro percorrer Names")
				os.Exit(1)
		}
		
		n.Id = id
		n.Name = name
		n.Email = email

		res = append(res, n)
		

	}
		w.Header().Set("Content-Type", "application/json")
		js, _ := json.Marshal(res) 
		fmt.Println(string(js))
		w.Write([]byte(js))
	
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main(){		
			
	http.HandleFunc("/", Index)		
	http.HandleFunc("/json", jsonPage)
	fmt.Println("Server started on: http://localhost:9000")
	http.ListenAndServe(":9000", nil)

}





	
	
	



