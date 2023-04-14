package main

import (
	"fmt"

	"github.com/b0gochort/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
