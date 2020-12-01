package main

import (
	"fmt"
	"sync"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
)

var wg sync.WaitGroup

var s Server

//TipoDistr para seleccionar el tipo de distribucion
func TipoDistr() string {
	var num int32
	distr := ""

	fmt.Println("Ingrese el numero correspondiente a la distribucion a utilizar: ")
	fmt.Println("1. Distribuida")
	fmt.Println("2. Centralizada")

	fmt.Scanln(&num)

	if num == 1 {
		distr = "Distribuida"
	} else {
		distr = "Centralizada"
	}

	return distr
}

func main() {
	s = Server{id: 1, ChunksTemporal: map[string][]*connection.Chunk{}, distr: TipoDistr()}
	Servidor()
}
