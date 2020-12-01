package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"
	"google.golang.org/grpc"
)

type book struct {
	cantPar         int32
	chunkpormaquina []int32
}

var wg sync.WaitGroup

//Server datos
type Server struct {
	id         int
	mux        *sync.Mutex
	log        map[string]book // string es el nombre del libro
	ipMaquinas map[int32]string
	librosDisp []string
	distr      string
}

var csvFile *os.File

//IP local 10.6.40.161
const (
	ipportListen    = "10.6.40.161:50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
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
	var val []string
	for index, element := range s.log[NombreLibro].chunkpormaquina {
		val = []string{"Parte_" + strconv.Itoa(len(s.log)) + "_" + strconv.Itoa(index+1), s.ipMaquinas[element]}
		csvwriter.Write(val)
	}

}

//IntIn chequea que un entero este en una lista
func IntIn(l []int32, n int32) bool {
	for _, element := range l {
		if element == n {
			return true
		}
	}
	return false
}

//NodosEnPropuesta entrega cuales nodos estan en la propuesta.
func NodosEnPropuesta(prop *connection.Distribucion) []int32 {
	l := []int32{}
	if IntIn(prop.ListaDataNodesChunk, 1) {
		l = append(l, 1)
	}
	if IntIn(prop.ListaDataNodesChunk, 2) {
		l = append(l, 2)
	}
	if IntIn(prop.ListaDataNodesChunk, 3) {
		l = append(l, 3)
	}
	return l
}

//AceptaPropuesta acepta o rechaza la propuesta con cierta probablidad.
func AceptaPropuesta(prop *connection.Distribucion) string {
	l := NodosEnPropuesta(prop)
	p1 := false
	p2 := false
	p3 := false
	for _, element := range l {
		if element == 1 {
			if ChequeaNodos(ipportDataNode1) != "Caido" {
				p1 = true
			}
		}
		if element == 2 {
			if ChequeaNodos(ipportDataNode2) != "Caido" {
				p2 = true
			}
		}
		if element == 3 {
			if ChequeaNodos(ipportDataNode3) != "Caido" {
				p3 = true
			}
		}
	}
	if p1 && p2 && p3 {
		return "SI"
	}
	if p1 && p2 {
		return "3"
	}
	if p1 && p3 {
		return "2"
	}
	if p2 && p3 {
		return "1"
	}
	return "NO" // If no, solo el nodo que envio el mensaje esta disponible.
}

//EnviaChunkCliente no es necesaria en el NameNode
func (s *Server) EnviaChunkCliente(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {

	return &connection.Message{}, nil
}

//EnviaChunkDataNode no es necesaria en el NameNode
func (s *Server) EnviaChunkDataNode(ctx context.Context, in *connection.Chunk) (*connection.Message, error) {

	return &connection.Message{}, nil
}

//ConsultaUbicacionArchivo consulta ubicacion al namenode de los chunks en los datanodes
func (s *Server) ConsultaUbicacionArchivo(ctx context.Context, in *connection.NombreLibro) (*connection.Distribucion, error) {

	return &connection.Distribucion{NombreLibro: in.NombreLibro, NumeroPar: s.log[in.NombreLibro].cantPar, ListaDataNodesChunk: s.log[in.NombreLibro].chunkpormaquina}, nil
}

//DescargaChunk no es necesaria en el NameNode
func (s *Server) DescargaChunk(ctx context.Context, in *connection.DivisionLibro) (*connection.Chunk, error) {

	return &connection.Chunk{}, nil
}

//EnviaPropuesta en el caso de namenode recibe propuesta de distribucion rechaza o acepta y guarda dicha distribucion, en el caso que venga aceptada solo la guarda.
func (s *Server) EnviaPropuesta(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {
	m := AceptaPropuesta(in)
	return &connection.Message{Message: m}, nil
}

//EnviaDistribucion distribuye los chunks segun la propuesta aceptada en el caso de que sea centralizada se encarga de que nadie utilice el log al mismo tiempo.
func (s *Server) EnviaDistribucion(ctx context.Context, in *connection.Distribucion) (*connection.Message, error) {
	if s.distr == "Centralizada" {
		defer s.mux.Unlock()
		s.mux.Lock()
	}
	s.log[in.NombreLibro] = book{cantPar: in.NumeroPar, chunkpormaquina: in.ListaDataNodesChunk}
	EditaResigtro(s, in.NombreLibro, csvFile)
	s.librosDisp = append(s.librosDisp, in.NombreLibro)
	return &connection.Message{Message: "Ok"}, nil
}

//ConsultaLibrosDisponibles Cliente downloader consulta libros disponibles
func (s *Server) ConsultaLibrosDisponibles(ctx context.Context, in *connection.Message) (*connection.Libros, error) {

	return &connection.Libros{LibrosDisponibles: s.librosDisp}, nil
}

//ChequeoPing chequea que un nodo no este caido
func (s *Server) ChequeoPing(ctx context.Context, in *connection.Message) (*connection.Message, error) {
	wg.Add(1)
	time.Sleep(5 * time.Second)
	wg.Done()
	return &connection.Message{Message: "Disponible"}, nil
}

//ConsultaUsoLog no es necesaria en el NameNode
func (s *Server) ConsultaUsoLog(ctx context.Context, in *connection.Message) (*connection.Message, error) {
	println("Entro a la consulta: " + time.Now().Format("2006-01-02 15:04:05"))
	wg.Wait()
	println("Termino: " + time.Now().Format("2006-01-02 15:04:05"))
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

	s := Server{id: 1, mux: &sync.Mutex{}, log: map[string]book{}, ipMaquinas: map[int32]string{}, librosDisp: []string{}, distr: ""}

	//Agrega el string ip de cada maquina
	s.ipMaquinas[1] = ipportDataNode1
	s.ipMaquinas[2] = ipportDataNode2
	s.ipMaquinas[3] = ipportDataNode3

	grpcServer := grpc.NewServer()

	//Crea el archivo de registro Log de NameNode
	csvFile = CreaRegistro()

	fmt.Println("En espera de Informacion Chunks para servidor")

	//Inicia el servicio de mensajeria que contiene las funciones grpc
	connection.RegisterMensajeriaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over "+ipportListen+": %v", err)
	}
}
