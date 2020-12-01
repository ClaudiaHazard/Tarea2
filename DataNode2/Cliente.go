package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//IP local 10.6.40.163
const (
	ipportNameNode  = "10.6.40.161:50051"
	ipportDataNode1 = "10.6.40.162:50051"
	ipportDataNode2 = "10.6.40.163:50051"
	ipportDataNode3 = "10.6.40.164:50051"
)

var wg2 sync.WaitGroup

//CreaPropuesta crea propuesta de distribucion en los datanodes.
func CreaPropuesta(Chunks []*connection.Chunk, Nodelist []int32) []int32 {
	var l []int32

	if len(Nodelist) == 3 {
		l = append(l, 1)
		if len(Chunks) > 1 {
			l = append(l, 2)
		}
		if len(Chunks) > 2 {
			l = append(l, 3)
		}
		if len(Chunks) > 3 {
			for i := 3; i < len(Chunks); i++ {
				n := rand.Float64()
				if n < 0.33 {
					l = append(l, 1)
				}
				if n > 0.33 && n < 0.66 {
					l = append(l, 2)
				}
				if n > 0.66 {
					l = append(l, 3)
				}
			}
		}
		return l
	}
	if len(Nodelist) == 2 {
		l = append(l, Nodelist[0])
		if len(Chunks) > 1 {
			l = append(l, Nodelist[1])
		}
		if len(Chunks) > 2 {
			for i := 2; i < len(Chunks); i++ {
				n := rand.Float64()
				if n < 0.5 {
					l = append(l, Nodelist[0])
				}
				if n > 0.5 {
					l = append(l, Nodelist[1])
				}
			}
		}
		return l
	}
	if len(Nodelist) == 1 {
		for i := 0; i < len(Chunks); i++ {
			l = append(l, Nodelist[0])
		}
		return l
	}
	return l
}

//EnviaDistribucionDistribuida envia distribucion utilizando algoritmo distribuido para utilizar el NameNode
func EnviaDistribucionDistribuida(conns []*grpc.ClientConn, conn *grpc.ClientConn, Distribucion *connection.Distribucion) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	//Consulta a los otros datanode por el uso del log
	wg.Add(1)
	go ConsultaUsoLogDistribuido(conns[0])
	wg.Add(1)
	go ConsultaUsoLogDistribuido(conns[1])
	wg.Wait()

	wg.Add(1)
	response, err := c.EnviaDistribucion(ctx, Distribucion)
	wg.Done()

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//EnviaDistribucionCentralizada envia distribucion utilizando el NameNode
func EnviaDistribucionCentralizada(conn *grpc.ClientConn, Distribucion *connection.Distribucion) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	ConsultaUsoLogCentralizado(conn)

	response, err := c.EnviaDistribucion(ctx, Distribucion)

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	return response
}

//BorrarElemento borra el elemento en la posicion pos.
func BorrarElemento(arr []int32, pos int) []int32 {
	copy(arr[pos:], arr[pos+1:])
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	return arr
}

//EnviaPropuestaDistribuida envia propuesta distribuida
func EnviaPropuestaDistribuida(conns []*grpc.ClientConn, listaChunks []*connection.Chunk, nombreLibro string) *connection.Distribucion {
	//c := connection.NewMensajeriaServiceClient(conns[0])
	//ctx := context.Background()

	listaNodos := []int32{1, 2, 3}

	l := CreaPropuesta(listaChunks, listaNodos)

	Distribucion := &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

	//Chequea que los otros 2 nodos acepten la propuesta
	respuesta1 := ChequeaCaido(conns[0]).Message
	respuesta2 := ChequeaCaido(conns[1]).Message

	if respuesta1 == "Caido" && respuesta2 != "Caido" {
		//listaNodos = []int32{1, 3}
		listaNodos = []int32{2, 3}
		//listaNodos = []int32{3,2}

		l = CreaPropuesta(listaChunks, listaNodos)

		Distribucion = &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

		//Chequea que el otro nodo acepte la propuesta
		respuesta1 = ChequeaCaido(conns[1]).Message

		return Distribucion
	}

	if respuesta2 == "Caido" && respuesta1 != "Caido" {
		//listaNodos = []int32{1, 2}
		listaNodos = []int32{2, 1}
		//listaNodos = []int32{3,1}

		l = CreaPropuesta(listaChunks, listaNodos)

		Distribucion = &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

		//Chequea que el otro nodo acepte la propuesta
		respuesta1 = ChequeaCaido(conns[0]).Message

		return Distribucion
	}

	if respuesta1 == "Caido" && respuesta2 == "Caido" {
		//listaNodos = []int32{1}
		listaNodos = []int32{2}
		//listaNodos = []int32{3}

		l = CreaPropuesta(listaChunks, listaNodos)

		Distribucion = &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

		return Distribucion
	}

	return Distribucion
}

//EnviaPropuestaCentralizada envia propuesta centralizada
func EnviaPropuestaCentralizada(conn *grpc.ClientConn, listaChunks []*connection.Chunk, nombreLibro string) *connection.Distribucion {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	//Envia propuesta con todos los nodos

	listaNodos := []int32{1, 2, 3}

	l := CreaPropuesta(listaChunks, listaNodos)

	Distribucion := &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

	response, err := c.EnviaPropuesta(ctx, Distribucion)

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}

	if response.Message == "SI" {
		return Distribucion
	}

	if response.Message != "NO" {
		//Elimina el nodo caido de la propuesta y crea nueva propuesta.

		if response.Message == "1" {
			listaNodos = BorrarElemento(listaNodos, 0)
		}

		if response.Message == "2" {
			listaNodos = BorrarElemento(listaNodos, 1)
		}

		if response.Message == "3" {
			listaNodos = BorrarElemento(listaNodos, 2)
		}

		l = CreaPropuesta(listaChunks, listaNodos)

		Distribucion = &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

		response, err = c.EnviaPropuesta(ctx, Distribucion)

		if err != nil {
			log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
		}
		if response.Message == "SI" {
			return Distribucion
		}
	}

	//En el caso de que los otros 2 nodos esten caidos.

	if response.Message == "NO" {
		//Nodo1
		//listaNodos = []int32{1}
		//Nodo2
		listaNodos = []int32{2}
		//Nodo3
		//listaNodos = []int32{3}
	}

	l = CreaPropuesta(listaChunks, listaNodos)

	Distribucion = &connection.Distribucion{NombreLibro: nombreLibro, ListaDataNodesChunk: l, NumeroPar: int32(len(l))}

	return Distribucion
}

//EnviaChunks envia chunks con la distribucion que fue aceptada previamente
func EnviaChunks(conn *grpc.ClientConn, chunk *connection.Chunk) *connection.Message {
	defer wg2.Done()
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.EnviaChunkDataNode(ctx, chunk)

	if err != nil {
		log.Fatalf("Error al llamar EnviaPropuesta: %s", err)
	}
	return response
}

//ReparteChunks envia los chunks a los datanodes correspondientes.
func ReparteChunks(conns []*grpc.ClientConn, nombreLibro string, Distribucion *connection.Distribucion) {
	for index, element := range Distribucion.ListaDataNodesChunk {
		fmt.Println("Envia Chunks a DataNode " + strconv.Itoa(int(element)))
		if element == 1 {
			wg2.Add(1)
			go EnviaChunks(conns[0], s.ChunksTemporal[nombreLibro][index])
		}
		if element == 2 {
			wg2.Add(1)
			GuardaChunk(s.ChunksTemporal[nombreLibro][index])
			wg2.Done()
		}
		if element == 3 {
			wg2.Add(1)
			go EnviaChunks(conns[1], s.ChunksTemporal[nombreLibro][index])
		}

	}
	wg2.Wait()
	wg.Done()
}

//ChequeaCaido envia aviso para saber si los datanode estan disponibles
func ChequeaCaido(conn *grpc.ClientConn) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	response, err := c.ChequeoPing(ctx, &connection.Message{Message: "Disponible?"})

	if err != nil {
		fmt.Println("Error de conexion con el DataNode, puede que este caido")
		return &connection.Message{Message: "Caido"}
	}

	fmt.Println(response.Message)
	return response
}

//ConsultaUsoLogDistribuido envia aviso para saber si se puede utilizar el log
func ConsultaUsoLogDistribuido(conn *grpc.ClientConn) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()
	s.timestamp = time.Now().Format("02/01/2006 03:04:05.000000 PM")
	fmt.Println("Consulta por el log")
	response, err := c.ConsultaUsoLog(ctx, &connection.Message{Message: s.timestamp})
	fmt.Println("Recibio respuesta del uso del log")
	s.timestamp = ""
	if err != nil {
		log.Fatalf("Error al llamar ConsultaUsoLog: %s", err)
	}
	defer wg.Done()
	return response
}

//ConsultaUsoLogCentralizado envia aviso para saber si se puede utilizar el log
func ConsultaUsoLogCentralizado(conn *grpc.ClientConn) *connection.Message {

	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()
	fmt.Println("Consulta por el log")
	response, err := c.ConsultaUsoLog(ctx, &connection.Message{Message: s.timestamp})
	fmt.Println("Recibio respuesta del uso del log")
	if err != nil {
		log.Fatalf("Error al llamar ConsultaUsoLog: %s", err)
	}
	return response
}

//EjecutaCliente ejecuta el cliente para guardar los chunks entre los datanodes
func EjecutaCliente(conn *grpc.ClientConn, connDN1 *grpc.ClientConn, connDN2 *grpc.ClientConn, nombreLibro string, distr string) string {
	conns := []*grpc.ClientConn{connDN1, connDN2}
	if distr == "Distribuida" {
		fmt.Println("Envia Propuesta de distribucion para el libro: " + nombreLibro)
		Distribucion := EnviaPropuestaDistribuida(conns, s.ChunksTemporal[nombreLibro], nombreLibro)
		fmt.Println("Envia Chunks")
		wg.Add(1)
		ReparteChunks(conns, nombreLibro, Distribucion)
		wg.Wait()
		fmt.Println("Envia distribucion para el libro: " + nombreLibro + ", tiempo: " + time.Now().Format("02/01/2006 03:04:05.000000 PM"))
		ok := EnviaDistribucionDistribuida(conns, conn, Distribucion)
		fmt.Println("Escribio distribucion del libro: " + nombreLibro + ", tiempo: " + time.Now().Format("02/01/2006 03:04:05.000000 PM"))
		delete(s.ChunksTemporal, nombreLibro)
		return ok.Message
	}
	if distr == "Centralizada" {
		fmt.Println("Envia Propuesta de distribucion para el libro: " + nombreLibro)
		Distribucion := EnviaPropuestaCentralizada(conn, s.ChunksTemporal[nombreLibro], nombreLibro)
		fmt.Println("Envia Chunks")
		ReparteChunks(conns, nombreLibro, Distribucion)
		fmt.Println("Envia distribucion para el libro: " + nombreLibro + ", tiempo: " + time.Now().Format("02/01/2006 03:04:05.000000 PM"))
		ok := EnviaDistribucionCentralizada(conn, Distribucion)
		fmt.Println("Escribio distribucion del libro: " + nombreLibro + ", tiempo: " + time.Now().Format("02/01/2006 03:04:05.000000 PM"))
		delete(s.ChunksTemporal, nombreLibro)
		return ok.Message
	}
	return ""
}

//Cliente Ejecucion de DataNode Cliente
func Cliente(nombreLibro string, distr string) {

	var connNN *grpc.ClientConn
	var connDN1 *grpc.ClientConn
	var connDN3 *grpc.ClientConn

	//Se crean las conexiones con NameNode y los DataNodes
	connNN, err := grpc.Dial(ipportNameNode, grpc.WithInsecure(), grpc.WithBlock())
	connDN1, err2 := grpc.Dial(ipportDataNode1, grpc.WithInsecure(), grpc.WithBlock())
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

	EjecutaCliente(connNN, connDN1, connDN3, nombreLibro, distr)
}
