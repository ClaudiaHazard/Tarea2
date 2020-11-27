package main

import (
	"fmt"
	"log"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"math"
	"os"
	//"math/rand"
	//"bufio"
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

//HaceChunk hace un chunk en base a una parte de un archivo
func HaceChunk(ch []byte, pos int32, n string) *connection.Chunk {
	return &connection.Chunk{Chunk: ch, NChunk: pos, NombreLibro: n} //nombre con formato incluido
}

//CreaChunks crea chunks de un libro.
func CreaChunks(name string, conn1 *grpc.ClientConn) {

	//elección de a donde enviar
	//seed:=rand.Intn(2)
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	const fileChunk = (1 * (1 << 20)) / 4 // 250 KB
	// num de partes en que se dividirá
	totalPartsNum := int32(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	for i := int32(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		ch := HaceChunk(partBuffer, i+1, name) //nombre con formato incluido)

		//ENVÍO A DATANODE USANDO GRPC

		c := connection.NewMensajeriaServiceClient(conn1)
		response, err3 := c.CargaArchivo(context.Background(), ch)
		if err3 != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf(response.Message)

	}
	fmt.Printf("Archivo subido: ")
}

//ArmaChunks arma archivo, luego de solicitar ubicación y obtener donde está cada chunk
func ArmaChunks(name string, conn1 *grpc.ClientConn, connNN *grpc.ClientConn) {
	_, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	nl := &connection.NombreLibro{NombreLibro: name}
	//SOLICITAR UBICACIONES
	c := connection.NewMensajeriaServiceClient(connNN)
	ubc, err2 := c.ConsultaUbicacionArchivo(context.Background(), nl)

	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var writePosition int64 = 0
	//parts := strings.Split(name, ".")
	//for j := uint64(0); j < (len(ubc.listaDatanode1)+ len(ubc.listaDatanode2) + len(ubc.listaDatanode3)); j++ { //original
	for j := int(0); j < len(ubc.ListaDataNode1); j++ { //propuesta
		//read a chunk
		//CONECTARSE A DATENODE DE ACUERDO A POS DE ARREGLO, Y ABRIRLO

		c2 := connection.NewMensajeriaServiceClient(conn1)
		newFileChunk, err3 := c2.DescargaChunk(context.Background(), nl)

		if err3 != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// calculate the bytes size of each chunk
		// we are not going to rely on previous data and constant

		var chunkSize int64 = int64(len(newFileChunk.Chunk))
		//chunkBufferBytes := make([]byte, chunkSize)
		fmt.Println("Appending at position : [", writePosition, "] bytes")
		writePosition = writePosition + chunkSize

		n, err := file.Write(newFileChunk.Chunk)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file.Sync()
		fmt.Println("Written ", n, " bytes")
		fmt.Println("Recombining part [", j, "] into : ", name)
	}
	file.Close()
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
