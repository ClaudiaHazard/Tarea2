package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"math/rand"

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

//CreaChunks crea chunks de un libro.
func CreaChunks(string name, *grpc.ClientConn conn1) {

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
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	for i := int32(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		ch:= connection.Chunk{Chunk:partBuffer,NChunk:i+1,IdArchivo:name}//nombre con formato incluido

		//ENVÍO A DATANODE USANDO GRPC

		"""if (seed==0){
			CargaArchivo(conn1,ch)
		}else if (seed==1){
			CargaArchivo(conn2,ch)
		}else {
			CargaArchivo(conn3,ch)
		}"""
		c := sm.NewMensajeriaServiceClient(conn1)
		response,err3:=c.CargaArchivo(context.Background(),ch)	
	}
	fmt.Println("Archivo subido: ")

}

//ArmaChunks arma archivo, luego de solicitar ubicación y obtener donde está cada chunk
func ArmaChunks(string name, *grpc.ClientConn conn1,  *grpc.ClientConn connNN) {
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

	//SOLICITAR UBICACIONES
	c := sm.NewMensajeriaServiceClient(connNN)
	ubc ,err2:= c.ConsultaUbicacionArchivo(context.Background(), name)

	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var writePosition int64 = 0
	parts := strings.Split(name, ".")
	//for j := uint64(0); j < (len(ubc.listaDatanode1)+ len(ubc.listaDatanode2) + len(ubc.listaDatanode3)); j++ { //original
	for j := uint64(0); j < (len(ubc.listaDatanode); j++ { //propuesta
		//read a chunk
		//CONECTARSE A DATENODE DE ACUERDO A POS DE ARREGLO, Y ABRIRLO

		"""if (ubc.listaDatanode[j])==1{

		}else if((ubc.listaDatanode[j])==2){

		}else {

		}"""

		c2 := sm.NewMensajeriaServiceClient(conn1)
		newFileChunk ,err3:= c2.DescargaChunk(context.Background(), name)

		if err2 != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//newFileChunk, err := os.Open(currentChunkFileName)

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

	var ty string
	var name string

}
