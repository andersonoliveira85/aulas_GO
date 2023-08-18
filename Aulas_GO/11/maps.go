package main

import "fmt"

func main() {
	//criando um map
	m := make(map[string]int)

	//atribuindo chave/valore
	m["k1"] = 15
	m["k2"] = 44

	//imprime o map m
	fmt.Println("map: ", m)

	//pega conteúdo
	v1 := m["k1"]
	fmt.Println(v1)

	//retorna o tamanho do map
	fmt.Println("Tamanho: ", len(m))

	//remove campo chave/valor
	delete(m, "k2")
	fmt.Println("map: ", m)
}
