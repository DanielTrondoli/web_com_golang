package models

import (
	"strconv"
	"strings"
	"web_com_golang/db"
)

var DataBase = "FILE"
var produtos = []Produto{}

func InsertProductFile(name string, price float64, descricao string, qtd int) {
	id := nextIndex()

	p := Produto{id, name, descricao, price, qtd}

	produtos = append(produtos, p)
}

func UpdateProductFile(id int, name string, price float64, descricao string, qtd int) {

	for i, p := range produtos {
		if id == p.Id {
			p.Nome = name
			p.Descricao = descricao
			p.Preco = price
			p.Quantidade = qtd
			produtos[i] = p
		}
	}
}

func GetAllProductsFile() []Produto {
	return produtos
}

func GetProductsByIdFile(idProduct int) Produto {
	for _, p := range produtos {
		if idProduct == p.Id {
			return p
		}
	}
	return Produto{}
}
func DeleteProductFile(idProduct int) {

	for i, p := range produtos {
		if idProduct == p.Id {
			removeIndex(i)
		}
	}
}

func removeIndex(index int) {
	produtos = append(produtos[:index], produtos[index+1:]...)
}

func nextIndex() int {
	idMax := 0

	for _, p := range produtos {
		if p.Id > idMax {
			idMax = p.Id
		}
	}

	return idMax + 1
}

func CarregarProducts() {

	arq := db.CarregarDBFile()

	for _, linha := range arq {
		produtos = append(produtos, stripProduct(linha))
	}
}

func SaveProducts() {

	arq := []string{}
	for _, p := range produtos {

		arq = append(arq, joinProds(p))
	}

	db.SaveDBFile(arq)

}

func stripProduct(linha string) Produto {
	p := Produto{}

	linha = strings.TrimSpace(linha)

	linhaStriped := strings.Split(linha, ";")

	if len(linhaStriped) == 5 {
		id, errId := strconv.Atoi(linhaStriped[0])
		name := linhaStriped[1]
		description := linhaStriped[2]
		price, errPrice := strconv.ParseFloat(linhaStriped[3], 64)
		qtd, errQtd := strconv.Atoi(linhaStriped[4])

		if errId != nil {
			panic(errPrice.Error())
		}
		if errPrice != nil {
			panic(errPrice.Error())
		}
		if errQtd != nil {
			panic(errQtd.Error())
		}

		p.Id = id
		p.Nome = name
		p.Descricao = description
		p.Preco = price
		p.Quantidade = qtd

		/*
			fmt.Println(id)
			fmt.Println(name)
			fmt.Println(descricao)
			fmt.Println(price)
			fmt.Println(qtd)
		*/
	}
	return p
}

func joinProds(p Produto) string {

	id := strconv.Itoa(p.Id)
	name := p.Nome
	desc := p.Descricao
	price := strconv.FormatFloat(p.Preco, 'f', 2, 64)
	qtd := strconv.Itoa(p.Quantidade)

	return id + ";" + name + ";" + desc + ";" + price + ";" + qtd + "\n"

}
