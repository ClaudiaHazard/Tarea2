package main

import (
	"fmt"
	"log"
	"strconv"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"bufio"
)

//IP local 10.6.40.161
const (
	ipportNameNode  = "10.6.40.161:50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
)

//LibrosDisponibles muestra libros disponibles en pantalla
func LibrosDisponibles(conn *grpc.ClientConn) {
	mes := &connection.Message{Message: "Consulta de libros"}
	c := connection.NewMensajeriaServiceClient(conn)
	li, err := c.ConsultaLibrosDisponibles(context.Background(), mes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Libros Disponibles:")
	for j := 0; j < len(li.LibrosDisponibles); j++ {
		fmt.Println(li.LibrosDisponibles[j])
	}
}

//ArmaChunks arma archivo, luego de solicitar ubicación y obtener donde está cada chunk
func ArmaChunks(name string, conn1 *grpc.ClientConn, conn2 *grpc.ClientConn, conn3 *grpc.ClientConn, connNN *grpc.ClientConn) {
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
	nl := &connection.NombreLibro{NombreLibro: name}
	c := connection.NewMensajeriaServiceClient(connNN)
	ubc, err2 := c.ConsultaUbicacionArchivo(context.Background(), nl)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c1 := connection.NewMensajeriaServiceClient(conn1)
	c2 := connection.NewMensajeriaServiceClient(conn2)
	c3 := connection.NewMensajeriaServiceClient(conn3)
	var writePosition int64 = 0
	for j := 0; j < int(ubc.NumeroPar); j++ { //original
		var newFileChunk *connection.Chunk
		dl := &connection.DivisionLibro{NombreLibro: name, NChunk: int32(j + 1)}
		if ubc.ListaDataNodesChunk[j] == 1 {
			newFileChunk, err2 = c1.DescargaChunk(context.Background(), dl)
		} else if ubc.ListaDataNodesChunk[j] == 2 {
			newFileChunk, err2 = c2.DescargaChunk(context.Background(), dl)
		} else {
			newFileChunk, err2 = c3.DescargaChunk(context.Background(), dl)
		}

		if err2 != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = int64(len(newFileChunk.Chunk))
		fmt.Println("Agregando " +strconv.Itoa(int(j + 1)) + " parte de " + strconv.Itoa(int(newFileChunk.NumeroPar))+ "para archivo " + name)
		writePosition = writePosition + chunkSize
		n, err := file.Write(newFileChunk.Chunk)
		n=n + 0
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file.Sync()
	}
	file.Close()
}

//Ejecucion de Downloader
func main() {
	var connDN1 *grpc.ClientConn
	var connDN2 *grpc.ClientConn
	var connDN3 *grpc.ClientConn
	var connNN *grpc.ClientConn
	var na string

	//Se crean conexiones
	connDN1, err := grpc.Dial(ipportDataNode1, grpc.WithInsecure(), grpc.WithBlock())
	connDN2, err2 := grpc.Dial(ipportDataNode2, grpc.WithInsecure(), grpc.WithBlock())
	connDN3, err3 := grpc.Dial(ipportDataNode3, grpc.WithInsecure(), grpc.WithBlock())
	connNN, err4 := grpc.Dial(ipportNameNode, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err2 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	if err3 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	if err4 != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}

	//solicitar archivos que existen
	LibrosDisponibles(connNN)
	fmt.Println("Ingrese nombre de archivo a descargar incluyendo formato (ej:ejemplo.pdf)")
	fmt.Scanln(&na)
	ArmaChunks(na, connDN1, connDN2, connDN3, connNN)
	fmt.Println(na, " descargado")

}
