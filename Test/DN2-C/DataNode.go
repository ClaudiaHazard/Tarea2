package main

import (
	"fmt"
	"log"
	"math/rand"

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
func CreaPropuesta(Chunks []byte) ([]int32, []int32, []int32) {
	var l1 []int32
	var l2 []int32
	var l3 []int32

	for i := 0; i < len(Chunks); i++ {
		n := rand.Float64()
		if n < 0.33 {
			l1 = append(l1, int32(i))
		}
		if n > 0.33 && n < 0.66 {
			l2 = append(l2, int32(i))
		}
		if n > 0.66 {
			l3 = append(l3, int32(i))
		}
	}
	return l1, l2, l3
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

	l1, l2, l3 := CreaPropuesta(listaChunks)

	Distribucion := &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNode1: l1, ListaDataNode2: l2, ListaDataNode3: l3}

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

	l1, l2, l3 := CreaPropuesta(listaChunks)

	Distribucion := &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNode1: l1, ListaDataNode2: l2, ListaDataNode3: l3}

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
