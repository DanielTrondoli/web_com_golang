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

func UpdateProduct(id, name string, price float64, descricao string, qtd int) {

	myDb := db.ConnectionDataBase()
	defer myDb.Close()

	queryInsert, err := myDb.Prepare("update produtos set name = $1, description = $2, price = $3, qtd = $4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	queryInsert.Exec(name, descricao, price, qtd, id)

}

func GetAllProducts() []Produto {
	myDb := db.ConnectionDataBase()
	defer myDb.Close()

	allProducts, err := myDb.Query("select * from produtos order by id")
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

func GetProductsById(idProduct string) Produto {

	myDb := db.ConnectionDataBase()
	defer myDb.Close()

	query, err := myDb.Query("select * from produtos where id=" + idProduct)
	if err != nil {
		panic(err.Error())
	}

	var id, qtd int
	var name, description string
	var price float64

	p := Produto{}

	if query.Next() {
		err = query.Scan(&id, &name, &description, &price, &qtd)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = name
		p.Descricao = description
		p.Preco = price
		p.Quantidade = qtd
	}

	return p
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
