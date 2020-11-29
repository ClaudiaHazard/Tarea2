package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	ipport          = ":50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
)

//CreaPropuesta crea propuesta de distribucion en los datanodes.
func CreaPropuesta(Chunks []byte) []int32 {
	var l []int32
	for i := 0; i < len(Chunks); i++ {
		n := rand.Float64()
		if n < 0.33 {
			l = append(l, 1)
		}
		if n > 0.33 && n < 0.66 {
			l = append(l, 2)
		}
		if n > 0.66 {
			l = append(l, 3)
		}
	}
	return l
}

//EnviaDistribucionDistribuida envia distribucion utilizando algoritmo distribuido para utilizar el NameNode
func EnviaDistribucionDistribuida(conns []*grpc.ClientConn, conn *grpc.ClientConn, Distribucion *connection.Distribucion) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.EnviaDistribucion(ctx, Distribucion)

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//EnviaDistribucionCentralizada envia distribucion utilizando el NameNode
func EnviaDistribucionCentralizada(conn *grpc.ClientConn, Distribucion *connection.Distribucion) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.EnviaDistribucion(ctx, Distribucion)

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//EnviaPropuestaDistribuida envia propuesta distribuida
func EnviaPropuestaDistribuida(conns []*grpc.ClientConn, listaChunks []byte, nombreLibro string) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conns[0])
	ctx := context.Background()

	l := CreaPropuesta(listaChunks)

	Distribucion := &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l}

	response, err := c.EnviaPropuesta(ctx, Distribucion)

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//EnviaPropuestaCentralizada envia propuesta centralizada
func EnviaPropuestaCentralizada(conn *grpc.ClientConn, listaChunks []byte, nombreLibro string) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	l := CreaPropuesta(listaChunks)

	Distribucion := &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l}

	response, err := c.EnviaPropuesta(ctx, Distribucion)

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

//ChequeaCaido envia aviso para saber si los datanode estan disponibles
func ChequeaCaido(conn *grpc.ClientConn) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.ChequeoPing(ctx, &connection.Message{Message: "Disponible?"})

	if err != nil {
		fmt.Println("Error de conexion con el DataNode, puede que este caido")
		return &connection.Message{Message: "Caido"}
	}

	print(response.Message)
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

	ChequeaCaido(conn)
	time.Sleep(5 * time.Second)
	ChequeaCaido(conn)
	time.Sleep(5 * time.Second)
	ChequeaCaido(conn)
}
