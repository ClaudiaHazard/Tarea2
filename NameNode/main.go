package main

import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go Clientes()
	wg.Add(1)
	go NameNode()
	wg.Wait()
}
