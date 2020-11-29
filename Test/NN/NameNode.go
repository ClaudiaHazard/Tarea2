package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

type book struct {
	cantPar         int32
	chunkpormaquina []int32
}

//Server datos
type Server struct {
	id  int
	mux *sync.Mutex
	log map[string]book // string es el nombre del libro
	ipMaquinas map[int32]string
}


const (
	ipportListen    = ":50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
	probabilidad    = 0.8
)

//CreaRegistro en el que escribira el NameNode.
func CreaRegistro() *os.File {
	csvFile, err := os.Create("Log.csv")

	if err != nil {
		log.Fatalf("Fallo al crear csv file: %s", err)
	}

	return csvFile

}

//EditaResigtro agrega registro del NameNode al csv file.
func EditaResigtro(s *Server, NombreLibro string, csvFile *os.File) {

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	csvwriter.Write([]string{NombreLibro, strconv.Itoa(int(s.log[NombreLibro].cantPar))})
	var val string
	for index, element := range s.log[NombreLibro].chunkpormaquina {
	 	val = []string{"Parte_"+strconv.Itoa(len(s.log)+"_"+strconv.Itoa(i+1),s.ipMaquinas[}
		csvwriter.Write(val)
	}

}

//AceptaPropuesta acepta o rechaza la propuesta con cierta probablidad.
func AceptaPropuesta() string {
	n := rand.Float64()
	if n < 0.8 {
		return "SI"
	}
	return "NO"
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

//EnviaPropuesta en el caso de namenode recibe propuesta de distribucion rechaza o acepta y guarda dicha distribucion, en el caso que venga aceptada solo la guarda.
func (s *Server) EnviaPropuesta(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {
	m := AceptaPropuesta()
	return &connection.Message{Message: m}, nil
}

//EnviaDistribucion distribuye los chunks segun la propuesta aceptada
func (s *Server) EnviaDistribucion(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {
	defer s.mux.Unlock()
	s.mux.Lock()
	chunkpormaquina := []*connection.ListaNChunks{}
	chunkpormaquina = in.ListaDataNodesChunk
	s.log[in.NombreLibro] = book{cantPar: in.NumeroPar, chunkpormaquina: chunkpormaquina}

	return &connection.Message{Message: "Ok"}, nil
}

//Servidor ejecucion de servidor para NameNode
func main() {
	fmt.Println("Hello there!")

	// Escucha las conexiones grpc
	lis, err := net.Listen("tcp", ipportListen)

	if err != nil {
		log.Fatalf("Failed to listen on "+ipportListen+": %v", err)
	}

	s := Server{id: 1, mux: &sync.Mutex{}, log: map[string]book{}, ipMaquinas: map[int32]string{}}

	//Agrega el string ip de cada maquina
	s.ipMaquinas[1] = ipportDataNode1
	s.ipMaquinas[2] = ipportDataNode2
	s.ipMaquinas[3] = ipportDataNode3

	grpcServer := grpc.NewServer()

	//Crea el archivo de registro Log de NameNode
	//csvFile = CreaRegistro()

	fmt.Println("En espera de Informacion Chunks para servidor")

	//Inicia el servicio de mensajeria que contiene las funciones grpc
	connection.RegisterMensajeriaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over "+ipportListen+": %v", err)
	}
}
