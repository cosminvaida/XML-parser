package main

import (
	"bytes"
	"errors"
	"io"
)

var dom = Stack()
var XMLMap = createXMLMap()

type Root *XMLElement

func XMLDocument(in interface{}) (*XMLElement, error) {
	var ioReader io.Reader
	switch in.(type) {
	case string:
		stringInput := in.(string)
		buffer := bytes.NewBufferString(stringInput)
		ioReader = buffer
	case io.Reader:
		ioReader = in.(io.Reader)
	default:
		return nil, errors.New("Cannot process the input type")
	}
	root, err := InitializeXML(ioReader)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func createXMLMap() map[uint64]XmlElements {
	return make(map[uint64]XmlElements)
}
