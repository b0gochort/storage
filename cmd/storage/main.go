package main

import (
	"fmt"
	"log"

	"github.com/b0gochort/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	trstortedFile, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it uploaded", trstortedFile)
}
