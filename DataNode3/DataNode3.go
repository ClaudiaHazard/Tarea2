package main

import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go Cliente()
	wg.Add(1)
	go Servidor()
	wg.Wait()
}
