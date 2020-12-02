package main

import (
	"fmt"
	"time"

	connection "github.com/ClaudiaHazard/Tarea2/Connection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//ChequeaCaido envia aviso para saber si los datanode estan disponibles
func ChequeaCaido(conn *grpc.ClientConn) *connection.Message {
	c := connection.NewMensajeriaServiceClient(conn)
	ctx := context.Background()

	res := make(chan string)
	timeout := make(chan bool, 1)

	go func() {
		response, err := c.ChequeoPing(ctx, &connection.Message{Message: "Disponible?"})
		if err != nil {
			fmt.Println("Error de conexion con el DataNode, puede que este caido")
			res <- "Caido"
		} else {
			res <- response.Message
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	select {
	case msg := <-res:
		return &connection.Message{Message: msg}

	case <-timeout:
		fmt.Println("Error de conexion con el DataNode, puede que este caido")
		return &connection.Message{Message: "Caido"}
	}
}

//ChequeaNodos chequea que el nodo no este caido
func ChequeaNodos(conn *grpc.ClientConn) string {

	return ChequeaCaido(conn).Message

}
