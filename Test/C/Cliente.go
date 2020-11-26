package main

import (
	"fmt"
	"log"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	ipportDataNode1 = ":50051"
)

//EsperaChunks espera chunks provenientes de cliente y otro datanode.
func EsperaChunks(conn *grpc.ClientConn) *connection.Chunk {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.RecibeChunks(ctx, &connection.Message{Message: "En espera"})

	if err != nil {
		log.Fatalf("Error al llamar RecibeChunks: %s", err)
	}

	print(response.NombreLibro)
	return response
}

//EnviaPropuestaDistribuida envia propuesta distribuida
func EnviaPropuestaDistribuida(conn *grpc.ClientConn) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.EnviaPropuesta(ctx, &connection.Distribucion{})

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//EnviaPropuestaCentralizada envia propuesta centralizada
func EnviaPropuestaCentralizada(conn *grpc.ClientConn) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.EnviaPropuesta(ctx, &connection.Distribucion{})

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//EnviaChunks envia chunks con la distribucion que fue aceptada previamente
func EnviaChunks(conn *grpc.ClientConn) *connection.Message {
	print("EnviaChunks")
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	print("EnviaChunks")

	response, err := c.DistribuyeChunks(ctx, &connection.Chunk{})

	print("EnviaChunks")

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//CreaChunks crea chunks de un libro.
func CreaChunks() {

}

//ArmaChunks arma archivo desde un arreglo de chunks
func ArmaChunks() {

}

//Ejecucion de Clientes
func main() {
	fmt.Println("Hello there!")

	var connDN1 *grpc.ClientConn

	//Se crea la conexion con el servidor DN1
	connDN1, err := grpc.Dial(ipportDataNode1, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	EsperaChunks(connDN1)
}