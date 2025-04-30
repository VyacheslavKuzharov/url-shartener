package infile

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type FileStorage struct {
	mutex   sync.RWMutex
	file    *os.File
	encoder *json.Encoder
}

func NewStorage(fileName string) *FileStorage {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return &FileStorage{
		file:    file,
		encoder: json.NewEncoder(file),
	}
}
