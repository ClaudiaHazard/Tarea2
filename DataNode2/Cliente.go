package main

import (
	"fmt"
	"log"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//IP local 10.6.40.163
const (
	ipportNameNode  = "10.6.40.161:50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
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

//Cliente Ejecucion de Cliente para DataNode
func Cliente() {
	fmt.Println("Hello there!")

	var connNN *grpc.ClientConn
	var connDN2 *grpc.ClientConn
	var connDN3 *grpc.ClientConn

	//Se crea la conexion con el servidor Logistica
	connNN, err := grpc.Dial(ipportNameNode, grpc.WithInsecure(), grpc.WithBlock())
	connDN2, err2 := grpc.Dial(ipportDataNode2, grpc.WithInsecure(), grpc.WithBlock())
	connDN3, err3 := grpc.Dial(ipportDataNode2, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err2 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err3 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	EsperaChunks(connNN)
	EsperaChunks(connDN2)
	EsperaChunks(connDN3)

	wg.Done()
}
