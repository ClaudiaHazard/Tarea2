package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"strconv"
	"os"
	"io/ioutil"
    "bufio"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

//Server datos
type Server struct {
	id int
}

const (
	ipportListen = ":50051"
)

//EnviaChunks cliente/datanode envía chunk a datanode
func (s *Server) EnviaChunks(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {
	parts:=strings.Split(in.NombreLibro,".")
	fileName := parts[0]+strconv.Itoa(int(in.NChunk))
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

//ConsultaUbicacionArchivo consulta ubicacion al namenode de los chunks en los datanodes
func (s *Server) ConsultaUbicacionArchivo(ctx context.Context, in *connection.NombreLibro) (*connection.Distribucion, error) {

	return &connection.Distribucion{}, nil
}

//DescargaChunk cliente descarga un chunk de alguno de los datanodes
func (s *Server) DescargaChunk(ctx context.Context, in *connection.DivisionLibro) (*connection.Chunk, error) {
	pa:=strings.Split(in.NombreLibro,".")
	fileopen:=pa[0]+strconv.Itoa(int(in.NChunk))
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
	return &connection.Chunk{Chunk: chunkBufferBytes, NChunk:in.NChunk, NombreLibro: in.NombreLibro}, nil
}

//EnviaPropuesta en el caso de namenode recibe propuesta de distribucion rechaza o acepta y guarda dicha distribucion, en el caso que venga aceptada solo la guarda.
func (s *Server) EnviaPropuesta(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//EnviaDistribucion distribuye los chunks segun la propuesta aceptada
func (s *Server) EnviaDistribucion(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {

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

	s := Server{id: 1}

	grpcServer := grpc.NewServer()

	fmt.Println("En espera de Informacion Chunks para servidor")

	//Inicia el servicio de mensajeria que contiene las funciones grpc
	connection.RegisterMensajeriaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over "+ipportListen+": %v", err)
	}
}
