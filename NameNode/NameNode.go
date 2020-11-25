package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

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
	ipport = "10.6.40.161:50051"
	//ipport = ":50051"
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
