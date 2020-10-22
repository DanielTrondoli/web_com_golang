package models

import "web_com_golang/db"

// Struct Produtos
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func InsertProduct(name string, price float64, descricao string, qtd int) {

	myDb := db.ConnectionDataBase()
	defer myDb.Close()

	queryInsert, err := myDb.Prepare("insert into produtos(name, description, price, qtd) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	queryInsert.Exec(name, descricao, price, qtd)
}

func GetAllProducts() []Produto {
	myDb := db.ConnectionDataBase()
	defer myDb.Close()

	allProducts, err := myDb.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	products := []Produto{}

	for allProducts.Next() {
		var id, qtd int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &qtd)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = name
		p.Descricao = description
		p.Preco = price
		p.Quantidade = qtd

		products = append(products, p)
	}

	return products
}

func DeleteProduct(idProduct string) {
	myDb := db.ConnectionDataBase()
	defer myDb.Close()

	queryDelete, err := myDb.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	queryDelete.Exec(idProduct)

}
