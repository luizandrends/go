package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:4]
	// arr[2] = 15
	// slice[0] = 123
	// fmt.Println(slice)

	// slice := []int{1, 2, 3}

	// fmt.Println(slice)

	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[0:]
	// fmt.Println(slice)

	// slice := []int{1, 2, 3, 4, 5}
	// slice2 := slice[:0]

	// fmt.Println(len(slice2), cap(slice2))

	// slice := []int{}
	// fmt.Println(slice == nil)

	// var filmesNoDB = []string{
	// 	"O Poderoso Chefão",
	// 	"Titanic",
	// 	"O Senhor dos Anéis: O Retorno do Rei",
	// 	"Matrix",
	// 	"Forrest Gump",
	// 	"O Rei Leão",
	// 	"Harry Potter e a PF",
	// 	"Gladiador",
	// 	"O Sexto Sendito",
	// 	"O Curioso Caso de Benjamin Button",
	// 	"Pulp Fiction: Tempo de Violencia",
	// 	"Esqueceram de Mim",
	// 	"O Exterminador do futuro 2: O Julgamento Final",
	// 	"O Fabuloso Destino de Amélie Poulain",
	// 	"O Labirinto do Fauno",
	// }

	// resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// matrix := [][]int{}
	// matrix3D := [][][]int{}

	// filmes := make([]string, 0, 10)

	// for _, id := range resultsFromApi {
	// 	filme := filmesNoDB[id]
		// filmes = append(filmes, filme)
	// }

	// fmt.Println(filmes)

	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[0:2:2]
	// fmt.Println(slice, cap(slice))

	// slice := []int{1, 2, 3, 4}

	// foo(slice)

	// m := map[string]string{
	// 	"Pedro": "Pessoa",
	// 	"Joaquim": "Pedro",
	// }
	// fmt.Println(m)

	// m := map[string][]int{
	// 	"Pedro": {1, 2, 3},
	// }
	// fmt.Println(m)

	// m := make(map[string]string)
	// m["Pedro"] = "Pessoa"
	// valor, ok := m["Pedro"]
	// valor2, ok := m["foo"]
	// fmt.Println(m)
	// fmt.Println(valor2, ok)

	// delete(m, "Pedro")

	// fmt.Println(valor, ok)

	m := map[string]string{
		"Pedro": "Pessoa",
		"Joaquim": "Pedro",
	}

	clear(m)
	fmt.Println(m["NaN"])
}