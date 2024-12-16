package main

import (
	"product-worker/worker"
)

func main() {
	worker := worker.New()
	worker.Run()
	worker.Close()
}