package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/lib/pq"
)

func ConnectionDataBase() *sql.DB {
	conexao := "user=daniel dbname=alura_loja password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}

func CarregarDBFile() []string {

	produtos := []string{}

	dbFile, err := os.Open("dbfile.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return produtos
	}

	stream := bufio.NewReader(dbFile)

	for {

		linha, err := stream.ReadString('\n')
		if err == io.EOF {
			//produtos = append(produtos, linha)
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro:", err)
			break
		}
		produtos = append(produtos, linha)

	}

	return produtos

}

func SaveDBFile(produtos []string) {

	os.Remove("dbfile.txt")
	arquivo, err := os.OpenFile("dbfile.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range produtos {
		fmt.Println(p)
		arquivo.WriteString(p)

	}

}
