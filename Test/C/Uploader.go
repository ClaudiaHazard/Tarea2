package main

import (
	"fmt"
	"log"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"math"
	"os"
	"math/rand"
	//"bufio"
)

const (
	ipportNameNode  = "10.6.40.161:50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
)

//HaceChunk hace un chunk en base a una parte de un archivo
func HaceChunk(ch []byte, pos int32, n string,t int32) *connection.Chunk {
	return &connection.Chunk{Chunk: ch, NChunk: pos, NombreLibro: n,NumeroPar:t} //nombre con formato incluido
}

//CreaChunks crea chunks de un libro.
func CreaChunks(name string, conn1 *grpc.ClientConn, conn2 *grpc.ClientConn, conn3 *grpc.ClientConn) {

	//elección de a donde enviar
	seed:=rand.Intn(2)
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
	c1:=connection.NewMensajeriaServiceClient(conn1)
	c2:=connection.NewMensajeriaServiceClient(conn2)
	c3:= connection.NewMensajeriaServiceClient(conn3)
	for i := int32(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		ch := HaceChunk(partBuffer, i+1, name,totalPartsNum) //nombre con formato incluido)
		if (seed==0){
			response, err3 := c1.EnviaChunkCliente(context.Background(), ch)
			if err3 != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf(response.Message)

		} else if (seed==1){
			response, err3 := c2.EnviaChunkCliente(context.Background(), ch)
			if err3 != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf(response.Message)

		}else {
			response, err3 := c3.EnviaChunkCliente(context.Background(), ch)
			if err3 != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf(response.Message)

		}

	}
	fmt.Printf("Archivo subido: ")
}


//Ejecucion de Clientes
func main() {
	var connDN1 *grpc.ClientConn
	var connDN2 *grpc.ClientConn
	var connDN3 *grpc.ClientConn

	//Se crea la conexion con el servidor DN1
	connDN1, err := grpc.Dial(ipportDataNode1, grpc.WithInsecure(), grpc.WithBlock())
	connDN2, err2 := grpc.Dial(ipportDataNode2, grpc.WithInsecure(), grpc.WithBlock())
	connDN3, err3 := grpc.Dial(ipportDataNode3, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err2 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err3 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	var na string

	fmt.Println("Ingrese nombre de archivo a subir incluyendo formato (ej:ejemplo.pdf)")
	fmt.Scanln(&na)
	CreaChunks(na, connDN1,connDN2,connDN3)
	fmt.Println(na," subido")		

}
