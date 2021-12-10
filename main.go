package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Item struct {
	Id   int
	Name string
}

func main() {
	doMain()
}

func doMain() {
	// encode
	var e bytes.Buffer
	enc := gob.NewEncoder(&e)
	err := enc.Encode(NewItem(0, "aaa"))
	if err != nil {
		panic(err)
	}
	fmt.Println("== encode ==")
	fmt.Println(e)

	// decode
	var i Item
	dec := gob.NewDecoder(&e)
	err = dec.Decode(&i)
	if err != nil {
		panic(err)
	}
	fmt.Println("== decode ==")
	fmt.Println(i)
}

func NewItem(id int, name string) *Item {
	return &Item{
		Id:   id,
		Name: name,
	}
}
