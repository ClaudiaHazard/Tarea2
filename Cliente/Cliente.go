package main

import (
	"fmt"
  	"bufio"         
    "io/ioutil"
    "math"
    "os"
    "strconv"
     "trings"
)

func main() {

    //var conn *grpc.ClientConn
    var ty string
    var name string
    //conn, err2 := grpc.Dial(ipport, grpc.WithInsecure(), grpc.WithBlock())
    //locks = "si"
    for true {
	    fmt.Println("Ingrese acción a realizar (subir o descargar): ")
	    fmt.Scanln(&ty)
	    fmt.Println("Ingrese nombre de archivo (con formato, ej: ejemplo.pdf): ")
	    fmt.Scanln(&name)
	    if ty == "subir" {
            file, err := os.Open(name)
            if err != nil {
                    fmt.Println(err)
                    os.Exit(1)
            }
            defer file.Close()
            fileInfo, _ := file.Stat()
            var fileSize int64 = fileInfo.Size()
            const fileChunk = (1 * (1 << 20))/4 // 250 KB
            // num de partes en que se dividirá
            totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
            fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
             for i := uint64(0); i < totalPartsNum; i++ {
                partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
                partBuffer := make([]byte, partSize)
                file.Read(partBuffer)


                //ENVÍO A DATANODE USANDO GRPC

            }
	    } else {
		     _, err := os.Create(name)

		    if err != nil {
	            fmt.Println(err)
	            os.Exit(1)
		    }

		     //set the newFileName file to APPEND MODE!!
		     // open files r and w

		    file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		    if err != nil {
		            fmt.Println(err)
		            os.Exit(1)
		    }


		     //SOLICITAR UBICACIONES: SE ASUME VIENEN EN UNA LISTA!

		    var writePosition int64 = 0
		    parts:=strings.Split(name,".")
		    for j := uint64(0); j < len(ubic); j++ {

	             //read a chunk

            	fileName := parts[0] + strconv.FormatUint(i, 10)
	            currentChunkFileName := fileName + strconv.FormatUint(j, 10)

	            //CONECTARSE A DATENODE DE ACUERDO A POS DE ARREGLO, Y ABRIRLO
	            	newFileChunk, err := os.Open(currentChunkFileName)

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
	}
}