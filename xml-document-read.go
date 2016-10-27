package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

func XMLDocumentFromString(in string) (*XMLElement, error) {
	root, err := ReadXML(in)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func XMLDocumentFromFile(r io.Reader) (*XMLElement, error) {
	root, err := ReadXML(r)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func ReadXML(in interface{}) (*XMLElement, error) {

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
	err := createDom(ioReader)
	if err != nil {
		return nil, err
	}
	Root, err := getRootElemet()
	if err != nil {
		return nil, err
	}
	return Root, nil
}

func createDom(in io.Reader) error {
	decoder := xml.NewDecoder(in)
	domTree := ""
	for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {

		if err != nil {
			return err
		}

		switch element := token.(type) {
		case xml.CharData:
			err := decodeCharData(element)
			if err != nil {
				return err
			}

		case xml.StartElement:
			_, err := decodeStartElement(element, &domTree)
			if err != nil {
				return err
			}

		case xml.EndElement:
			_, err := decodeEndElement(&domTree)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getAttributes(element xml.StartElement) []Attributes {
	attributes := []Attributes{}
	for _, attrEl := range element.Attr {
		attribute := Attributes{}
		attribute.NameSpace = attrEl.Name.Space
		attribute.Key = attrEl.Name.Local
		attribute.Value = attrEl.Value
		attributes = append(attributes, attribute)
	}
	return attributes
}

func decodeCharData(element xml.CharData) error {
	if str := strings.TrimSpace(string(element)); str != "" {
		dom.UpdateValueLastElemet(str)
	}
	return nil
}

func decodeStartElement(element xml.StartElement, domTree *string) (*XMLElement, error) {
	node := XMLElement{}
	node.Name = strings.TrimSpace(string(element.Name.Local))
	node.Attributes = getAttributes(element)

	*domTree = *domTree + "/" + node.Name
	node.Path = *domTree
	dom.Push(node)

	return &node, nil
}

func decodeEndElement(domTree *string) (*XMLElement, error) {
	if dom.Length() == 1 {
		return nil, nil
	}
	nodeChild, err := dom.Pop()
	if err != nil {
		return nil, err
	}
	nodeParent, err := dom.Pop()
	if err != nil {
		return nil, err
	}

	nodeParent.Childs = append(nodeParent.Childs, nodeChild)
	dom.Push(nodeParent)
	*domTree = strings.TrimSuffix(*domTree, "/"+nodeChild.Name)

	return &nodeChild, nil
}

func updateChildPath(rootPath string, node XMLElement) {
	node.Path = rootPath + "/" + node.Name
	for _, childNode := range node.Childs {
		updateChildPath(node.Path, childNode)
	}
}

func getRootElemet() (*XMLElement, error) {
	root, err := dom.Pop()
	if err != nil {
		return nil, err
	}
	return &root, nil
}
