package config

import(

	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	
)


func dbConn()(db *sql.DB, err error){

	db, err := sql.Open("mysql", "cursos09:lp231210@/cursos09")
	if err != nil{
		fmt.Println("Erro: ", err)
		//os.Exit(1)
		
	}

	return 
}
