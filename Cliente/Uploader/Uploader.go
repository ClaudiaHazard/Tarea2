package main

import (
	"fmt"
	"log"
	"time"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	//"bufio"
)

var wg sync.WaitGroup

//IP local 10.6.40.161
const (
	ipportNameNode  = "10.6.40.161:50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
)

//HaceChunk hace un chunk en base a una parte de un archivo
func HaceChunk(ch []byte, pos int32, n string, t int32) *connection.Chunk {
	return &connection.Chunk{Chunk: ch, NChunk: pos, NombreLibro: n, NumeroPar: t} //nombre con formato incluido
}

//CreaChunks crea chunks de un libro.
func CreaChunks(name string, conn1 *grpc.ClientConn, conn2 *grpc.ClientConn, conn3 *grpc.ClientConn) {

	work := false

	seed := 0

	xcon, _ := grpc.Dial(":50052")
	c := connection.NewMensajeriaServiceClient(xcon)

	for work != true {
		//elección de a donde enviar
		seed = rand.Intn(3)

		if seed == 0 {
			c = connection.NewMensajeriaServiceClient(conn1)
		}
		if seed == 1 {
			c = connection.NewMensajeriaServiceClient(conn2)
		}
		if seed == 2 {
			c = connection.NewMensajeriaServiceClient(conn3)
		}

		_, err := c.ChequeoPing(context.Background(), &connection.Message{Message: "Disponible?"})

		if err != nil {
			fmt.Println("Error de conexion con el DataNode" + strconv.Itoa(int(seed+1)) + ", puede que este caido")
			work = false
		} else {
			work = true
		}
	}

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
		ch := HaceChunk(partBuffer, i+1, name, totalPartsNum) //nombre con formato incluido

		fmt.Println("Enviando chunk n° " + strconv.Itoa(int(i)) + " de " + strconv.Itoa(int(ch.NumeroPar)) + "de Libro: " + name)

		_, err := c.EnviaChunkCliente(context.Background(), ch)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Enviado chunk n° " + strconv.Itoa(int(i)) + " de " + strconv.Itoa(int(ch.NumeroPar)))

	}
	fmt.Println(name, " terminó de enviarse")
	wg.Done()
}

//Ejecucion de Uploader
func main() {
	var connDN1 *grpc.ClientConn
	var connDN2 *grpc.ClientConn
	var connDN3 *grpc.ClientConn

	//Se crean conexiones
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
	nas := [5]string{"2.pdf", "Emma.pdf", "Dracula.pdf", "Test.pdf", "Lab1.pdf"}

	wg.Add(1)
	go CreaChunks(nas[0], connDN1, connDN2, connDN3)
	wg.Add(1)
	go CreaChunks(nas[1], connDN1, connDN2, connDN3)
	fmt.Println("Espera 8 segundos para botar un DataNode en caso de que se quiera probar")
	time.Sleep(8000 * time.Millisecond)
	wg.Add(1)
	go CreaChunks(nas[2], connDN1, connDN2, connDN3)
	wg.Add(1)
	go CreaChunks(nas[3], connDN1, connDN2, connDN3)
	wg.Add(1)
	go CreaChunks(nas[4], connDN1, connDN2, connDN3)

	wg.Wait()

}
