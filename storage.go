package main

type Storage[T any] struct {
	FileName string
}

// T is just a placeholder for the type you pass in

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}