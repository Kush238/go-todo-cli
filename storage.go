package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

// T is just a placeholder for the type you pass in

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}
// Functions with first letter capital will only be exported to other files 
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load (data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data) 
	// Unmarshal will convert JSON data and convert it to data argument we pass

}