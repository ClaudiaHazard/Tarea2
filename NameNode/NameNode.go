package main

import (
	"context"
	"fmt"
	"log"
	"net"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

//Server datos
type Server struct {
	id int
}

//IP local 10.6.40.161
const (
	ipportListen = "10.6.40.161:50051"
	//ipportListen = ":50051"
)

//CargaArchivo carga archivo a un datanode
func (s *Server) CargaArchivo(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//EnviaDistribucion distribuye los chunks segun la propuesta aceptada
func (s *Server) EnviaDistribucion(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//EnviaChunks Recibe propuesta de un namenode
func (s *Server) EnviaChunks(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {
	print("Se EnviaChunk")
	return &connection.Message{Message: "hola"}, nil
}

//ConsultaUbicacionArchivo consulta ubicacion al namenode de los chunks en los datanodes
func (s *Server) ConsultaUbicacionArchivo(ctx context.Context, in *connection.NombreLibro) (*connection.Distribucion, error) {

	return &connection.Distribucion{}, nil
}

//DescargaChunk descarga un chunk de alguno de los datanodes
func (s *Server) DescargaChunk(ctx context.Context, in *connection.DivisionLibro) (*connection.Chunk, error) {

	return &connection.Chunk{}, nil
}

//RecibeChunks Recibe propuesta de un namenode
func (s *Server) RecibeChunks(ctx context.Context, in *connection.Message) (*connection.Chunk, error) {
	print(in.Message)
	print("Se EnviaChunk")
	return &connection.Chunk{NombreLibro: "2"}, nil
}

//EnviaPropuesta en el caso de namenode recibe propuesta de distribucion rechaza o acepta y guarda dicha distribucion, en el caso que venga aceptada solo la guarda.
func (s *Server) EnviaPropuesta(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//DistribuyeChunks distribuye los chunks segun la propuesta aceptada
func (s *Server) DistribuyeChunks(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {

	return &connection.Message{Message: "Ok"}, nil
}

//ConsultaLibrosDisponibles Cliente downloader consulta libros disponibles
func (s *Server) ConsultaLibrosDisponibles(ctx context.Context, in *connection.Message) (*connection.Libros, error) {

	return &connection.Libros{}, nil
}

//ChequeoPing chequea que un nodo no este caido
func (s *Server) ChequeoPing(ctx context.Context, in *connection.Message) (*connection.Message, error) {

	return &connection.Message{Message: "Disponible?"}, nil
}

//NameNode ejecucion de servidor para NameNode
func NameNode() {
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

	wg.Done()
}
