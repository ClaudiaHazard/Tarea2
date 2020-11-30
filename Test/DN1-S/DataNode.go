package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

//Server datos
type Server struct {
	id             int
	ChunksTemporal map[string][]*connection.Chunk //string es el nombre del libro
	usandoLog      bool
}

const (
	ipportListen    = ":50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
	probabilidad    = 0.8
)

var wg sync.WaitGroup

//AceptaPropuesta acepta o rechaza la propuesta con cierta probablidad.
func AceptaPropuesta() string {
	n := rand.Float64()
	if n < 0.8 {
		return "SI"
	}
	return "NO"
}

//GuardaChunk guarda el chunk en archivo.
func GuardaChunk(chunk *connection.Chunk) {

}

//EnviaChunkCliente no es necesaria en el NameNode
func (s *Server) EnviaChunkCliente(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {
	parts := strings.Split(in.NombreLibro, ".")
	fileName := parts[0] + strconv.Itoa(int(in.NChunk))
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write/save buffer to disk
	ioutil.WriteFile(fileName, in.Chunk, os.ModeAppend)

	fmt.Println("Downloaded : ", fileName)
	return &connection.Message{Message: "Descargada\n"}, nil
}

//EnviaChunkDataNode no es necesaria en el NameNode
func (s *Server) EnviaChunkDataNode(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {
	GuardaChunk(in)
	return &connection.Message{Message: "Guardado"}, nil
}

//ConsultaUbicacionArchivo consulta ubicacion al namenode de los chunks en los datanodes
func (s *Server) ConsultaUbicacionArchivo(ctx context.Context, in *connection.NombreLibro) (*connection.Distribucion, error) {

	return &connection.Distribucion{}, nil
}

//DescargaChunk cliente descarga un chunk de alguno de los datanodes
func (s *Server) DescargaChunk(ctx context.Context, in *connection.DivisionLibro) (*connection.Chunk, error) {
	pa := strings.Split(in.NombreLibro, ".")
	fileopen := pa[0] + strconv.Itoa(int(in.NChunk))
	newFileChunk, err := os.Open(fileopen)

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
	var chunkSize int64 = chunkInfo.Size()
	chunkBufferBytes := make([]byte, chunkSize)
	reader := bufio.NewReader(newFileChunk)
	_, err = reader.Read(chunkBufferBytes)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &connection.Chunk{Chunk: chunkBufferBytes, NChunk: in.NChunk, NombreLibro: in.NombreLibro}, nil
}

//EnviaPropuesta en el caso de namenode recibe propuesta de distribucion rechaza o acepta y guarda dicha distribucion, en el caso que venga aceptada solo la guarda.
func (s *Server) EnviaPropuesta(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {
	m := AceptaPropuesta()
	return &connection.Message{Message: m}, nil
}

//EnviaDistribucion distribuye los chunks segun la propuesta aceptada
func (s *Server) EnviaDistribucion(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//ConsultaLibrosDisponibles no es utilizada por DataNode
func (s *Server) ConsultaLibrosDisponibles(ctx context.Context, in *connection.Message) (*connection.Libros, error) {

	return &connection.Libros{}, nil
}

//ChequeoPing chequea que un nodo no este caido
func (s *Server) ChequeoPing(ctx context.Context, in *connection.Message) (*connection.Message, error) {

	return &connection.Message{Message: "Disponible"}, nil
}

//ConsultaUsoLog chequea que un nodo no este caido
func (s *Server) ConsultaUsoLog(ctx context.Context, in *connection.Message) (*connection.Message, error) {
	wg.Wait()
	return &connection.Message{Message: "Ok"}, nil
}

//Servidor ejecucion de servidor para DataNode
func main() {
	fmt.Println("Hello there!")

	// Escucha las conexiones grpc
	lis, err := net.Listen("tcp", ipportListen)

	if err != nil {
		log.Fatalf("Failed to listen on "+ipportListen+": %v", err)
	}

	s := Server{id: 1, ChunksTemporal: map[string][]*connection.Chunk{}, usandoLog: false}

	grpcServer := grpc.NewServer()

	fmt.Println("En espera de Informacion Chunks para servidor")

	//Inicia el servicio de mensajeria que contiene las funciones grpc
	connection.RegisterMensajeriaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over "+ipportListen+": %v", err)
	}
}
