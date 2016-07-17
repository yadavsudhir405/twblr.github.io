package main

import "fmt"

type Resource struct {
	name   string
	closed bool
}

func Open(name string) (*Resource, error) {
	return &Resource{name: name}, nil
}

func PanicOpen(name string) (*Resource, error) {
	panic("failed to open resource")
}

func (res *Resource) Close() error {
	fmt.Println("Closing resource %s", res.name)
	res.closed = true
	return nil
}

func allocateResource() *Resource {
	res, err := Open("Hello")
	if err != nil {
		return nil
	}
	return res
}

func handlePanicOpen() error {
	return nil
}
