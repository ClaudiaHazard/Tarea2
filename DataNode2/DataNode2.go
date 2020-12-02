package main

import (
	"fmt"
	"log"
	"sync"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup
var wgInf sync.WaitGroup
var s Server

var connNN *grpc.ClientConn
var connDN1 *grpc.ClientConn
var connDN3 *grpc.ClientConn

var err error
var err2 error
var err3 error

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
	fmt.Println("Se utilizara la distribucion " + distr)
	return distr
}

func main() {
	s = Server{id: 1, ChunksTemporal: map[string][]*connection.Chunk{}, distr: TipoDistr(), timestamp: ""}

	wgInf.Add(1)
	go Servidor()

	//Se crean las conexiones con NameNode y los DataNodes
	connNN, err = grpc.Dial(ipportNameNode, grpc.WithInsecure(), grpc.WithBlock())
	connDN1, err2 = grpc.Dial(ipportDataNode1, grpc.WithInsecure(), grpc.WithBlock())
	connDN3, err3 = grpc.Dial(ipportDataNode3, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err2 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err3 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	wgInf.Wait()
}
