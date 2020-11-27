package main

import (
	"fmt"
	"log"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	ipport = ":50051"
)

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

	response, err := c.EnviaChunks(ctx, &connection.Chunk{})

	print("EnviaChunks")

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//Ejecucion de DataNode Cliente
func main() {
	fmt.Println("Hello there!")

	var conn *grpc.ClientConn

	//Se crea la conexion con el servidor Logistica
	conn, err := grpc.Dial(ipport, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	EnviaChunks(conn)

}
