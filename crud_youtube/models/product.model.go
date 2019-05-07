package models

type ProdutoModel struct{

	Db *sql.DB

}


func (produtoModel ProdutoModel) findAll() ([]entities.Produto, error){

	rows, err:= 
}