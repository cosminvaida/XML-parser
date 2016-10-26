package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Node struct {
	Name       string
	Value      string
	Attr       []Attr
	ChildsNode []Node
}

type Attr struct {
	NameSpace string
	Key       string
	Value     string
}

func main() {
	strapsFilePath, err := filepath.Abs("file.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the straps.xml file
	file, err := os.Open(strapsFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	f, err23 := NewDocument(file)
	fmt.Println(f)
	fmt.Println(err23)

	//GetElements("HotelReference")

	Map(Envelope{})

}
