package models

import "github.com/DanielTrondoli/web_com_golang/db"

// Struct Produtos
type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
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

		p.Nome = name
		p.Descricao = description
		p.Preco = price
		p.Quantidade = qtd

		products = append(products, p)
	}

	return products
}
