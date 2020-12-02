package main

import (
	"fmt"
	"log"
	"time"
    "strings"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
    "io/ioutil"
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
	const fileChunk = (1 * (1 << 20)) / 4 
	totalPartsNum := int32(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Println("El archivo" + name + "Tiene " + strconv.Itoa(int(totalPartsNum)) + " partes")
	for i := int32(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		ch := HaceChunk(partBuffer, i+1, name, totalPartsNum)
		_, err := c.EnviaChunkCliente(context.Background(), ch)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Enviado chunk n° " + strconv.Itoa(int(i + 1)) + " de " + strconv.Itoa(int(ch.NumeroPar)))
	}
	fmt.Println(name, " terminó de enviarse")
	wg.Done()
}

//Ejecucion de Uploader
func main() {
	var connDN1 *grpc.ClientConn
	var connDN2 *grpc.ClientConn
	var connDN3 *grpc.ClientConn

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

	files, err := ioutil.ReadDir("./")
	if err != nil {
	    fmt.Println("Error obteniendo archivos: ",err)
	}

	var toread []string
	for _, f := range files {
		if ((f.IsDir()!=true) && (strings.Contains(f.Name(),".go")!=true) && (f.Name()!="Makefile")){	
			toread=append(toread,f.Name())
	    }
	}
	for i:=0;i<len(toread);i++ {
		wg.Add(1)
		go CreaChunks(toread[i], connDN1, connDN2, connDN3)
	    if (i==1){
			fmt.Println("Espera 10 segundos para botar un DataNode en caso de que se quiera probar")
			time.Sleep(10000 * time.Millisecond)	    	
		}
	}
}
