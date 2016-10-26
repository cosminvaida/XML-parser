package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"
)

var xmlDom Node

func NewDocument(r io.Reader) (*stack, error) {
	decoder := xml.NewDecoder(r)
	stack := Stack()
	for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {

		if err != nil {
			return nil, err
		}
		switch element := token.(type) {
		case xml.CharData:
			if str := strings.TrimSpace(string(element)); str != "" {
				if node, err := stack.Pop(); err != nil {
					return nil, err
				} else {
					node.Value = str
					stack.Push(node)
				}
			}
		case xml.StartElement:
			node := Node{}
			node.Name = strings.TrimSpace(string(element.Name.Local))
			node.Attr = GetAttrs(element)
			stack.Push(node)

		case xml.EndElement:
			if stack.Length() == 1 {
				continue
			}
			if nodeChild, err := stack.Pop(); err != nil {
				return nil, err
			} else {
				if nodeParent, err := stack.Pop(); err != nil {
					return nil, err
				} else {
					nodeParent.ChildsNode = append(nodeParent.ChildsNode, nodeChild)
					stack.Push(nodeParent)
				}
			}
		}
	}
	if nodeChild, err := stack.Pop(); err != nil {
		return nil, err
	} else {
		xmlDom = nodeChild
	}

	return nil, nil

}

func GetElements(elementName string) []Node {
	node := xmlDom
	nodes := Search(elementName, node)
	return nodes
}

func Search(elementName string, node Node) []Node {
	var nodes = []Node{}
	if node.Name == elementName {
		nodes = append(nodes, node)
	}
	for _, childNode := range node.ChildsNode {

		nodesChild := Search(elementName, childNode)
		for _, n := range nodesChild {
			nodes = append(nodes, n)
		}

	}

	return nodes
}

type Envelope struct {
	Header1 Header1
	Body    Body
}

type Header struct {
}

type Body struct {
}

func Map(in interface{}) {

	typeIn := reflect.TypeOf(in)
	MapType(typeIn)

}

func MapType(in reflect.Type) {
	name := in.Name()
	fmt.Println(name)
	node := GetElements(name)
	for index, _ := range make([]int, in.NumField()) {
		var field = in.Field(index)
		nameFieldChild := field.Type.Name()
		if nameFieldChild == "string" {
			if strings.Contains(string(field.Tag), "attr") {
				valueAttr := getValueAttr(node, field.Name)
				fmt.Println(valueAttr)
				fmt.Println(nameFieldChild)
				fmt.Println(node)

			}
			fmt.Println(nameFieldChild)
			continue
		}

		MapType(field.Type)

	}
}

func getValueAttr(nodes []Node, key string) string {
	var node = nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Value
		}
	}
	return ""
}

func GetAttrs(el xml.StartElement) []Attr {
	attrs := []Attr{}
	for _, attrEl := range el.Attr {
		attr := Attr{}
		attr.NameSpace = attrEl.Name.Space
		attr.Key = attrEl.Name.Local
		attr.Value = attrEl.Value
		attrs = append(attrs, attr)
	}
	return attrs
}
