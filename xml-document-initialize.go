package main

import (
	"encoding/xml"
	"io"
	"strings"
)

func InitializeXML(in io.Reader) (*XMLElement, error) {
	err := createDom(in)
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

func getRootElemet() (*XMLElement, error) {
	root, err := dom.Pop()
	if err != nil {
		return nil, err
	}
	return &root, nil
}

func getAttributes(element xml.StartElement) *XmlNodeAttributes {
	attributes := XmlNodeAttributes{}
	for _, attrEl := range element.Attr {
		attribute := Attribute{}
		attribute.NameSpace = attrEl.Name.Space
		attribute.Key = attrEl.Name.Local
		attribute.Value = attrEl.Value
		attributes.append(attribute)
	}
	return &attributes
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
	if nodeParent.Childs == nil {
		nodeParent.Childs = &XmlElements{}
	}
	nodeChild.PathHash = hash(nodeChild.Path)

	nodeParent.Childs.append(nodeChild)

	dom.Push(nodeParent)
	mapNode(*domTree, &nodeChild)
	*domTree = strings.TrimSuffix(*domTree, "/"+nodeChild.Name)

	return &nodeChild, nil
}
