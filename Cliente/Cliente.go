package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	//"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup

//Server datos
type Server struct {
	id int
}

//IP local 10.6.40.161
const (
	//ipport = "10.6.40.162:50051"
	ipport = ":50051"
)

//CreaChunks crea chunks de un libro.
func CreaChunks() {
	//var conn *grpc.ClientConn
	var ty string
	var name string
	//conn, err2 := grpc.Dial(ipport, grpc.WithInsecure(), grpc.WithBlock())
	//locks = "si"
	for true {
		fmt.Println("Ingrese acción a realizar (subir o descargar): ")
		fmt.Scanln(&ty)
		fmt.Println("Ingrese nombre de archivo (con formato, ej: ejemplo.pdf): ")
		fmt.Scanln(&name)
		if ty == "subir" {
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
			totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
			fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
			for i := uint64(0); i < totalPartsNum; i++ {
				partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
				partBuffer := make([]byte, partSize)
				file.Read(partBuffer)

				//ENVÍO A DATANODE USANDO GRPC

			}
		} else {
			_, err := os.Create(name)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			//set the newFileName file to APPEND MODE!!
			// open files r and w

			file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			//SOLICITAR UBICACIONES: SE ASUME VIENEN EN UNA LISTA!

			var writePosition int64 = 0
			parts := strings.Split(name, ".")
			for j := uint64(0); j < 20; j++ { //REEMPLAZAR CUANDO SE TENGA CONEXIÓN

				//read a chunk

				fileName := parts[0] + strconv.FormatUint(j, 10) //CAMBIAR CUANDO SE TENGA CONEXIÓN
				currentChunkFileName := fileName + strconv.FormatUint(j, 10)

				//CONECTARSE A DATENODE DE ACUERDO A POS DE ARREGLO, Y ABRIRLO
				newFileChunk, err := os.Open(currentChunkFileName)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				defer newFileChunk.Close()

				chunkInfo, err := newFileChunk.Stat()

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				// calculate the bytes size of each chunk
				// we are not going to rely on previous data and constant

				var chunkSize int64 = chunkInfo.Size()
				chunkBufferBytes := make([]byte, chunkSize)

				fmt.Println("Appending at position : [", writePosition, "] bytes")
				writePosition = writePosition + chunkSize

				// read into chunkBufferBytes
				reader := bufio.NewReader(newFileChunk)
				_, err = reader.Read(chunkBufferBytes)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				n, err := file.Write(chunkBufferBytes)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				file.Sync()
				chunkBufferBytes = nil
				fmt.Println("Written ", n, " bytes")
				fmt.Println("Recombining part [", j, "] into : ", name)
			}
			file.Close()
		}
	}
}

//ArmaChunks arma archivo desde un arreglo de chunks
func ArmaChunks() {

}

//CargaArchivo carga archivo a un datanode
func (s *Server) CargaArchivo(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//ConsultaUbicacionArchivo consulta ubicacion al namenode de los chunks en los datanodes
func (s *Server) ConsultaUbicacionArchivo(ctx context.Context, in *connection.IdArchivo) (*connection.Distribucion, error) {

	return &connection.Distribucion{}, nil
}

//DescargaChunk descarga un chunk de alguno de los datanodes
func (s *Server) DescargaChunk(ctx context.Context, in *connection.IdArchivo) (*connection.Chunk, error) {

	return &connection.Chunk{}, nil
}

//RecibeChunks Recibe propuesta de un namenode
func (s *Server) RecibeChunks(ctx context.Context, in *connection.Message) (*connection.Chunk, error) {
	print(in.Message)
	print("Se EnviaChunk")
	return &connection.Chunk{IdArchivo: 1}, nil
}

//EnviaPropuesta en el caso de namenode recibe propuesta de distribucion rechaza o acepta y guarda dicha distribucion, en el caso que venga aceptada solo la guarda.
func (s *Server) EnviaPropuesta(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//DistribuyeChunks distribuye los chunks segun la propuesta aceptada
func (s *Server) DistribuyeChunks(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

func main() {
	// Escucha las conexiones grpc
	lis, err := net.Listen("tcp", ipport)

	if err != nil {
		log.Fatalf("Failed to listen on "+ipport+": %v", err)
	}

	s := Server{id: 1}

	grpcServer := grpc.NewServer()

	fmt.Println("En espera de Informacion Chunks para servidor")

	//Inicia el servicio de mensajeria que contiene las funciones grpc
	connection.RegisterMensajeriaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over "+ipport+": %v", err)
	}

}
