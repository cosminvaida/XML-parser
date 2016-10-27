package main

type XMLElement struct {
	Name       string
	InnerValue string
	Attributes []Attributes
	Childs     []XMLElement
	Path       string
}

type Attributes struct {
	NameSpace string
	Key       string
	Value     string
}

type XmlNodeSplice []XMLElement

type Root *XMLElement

var dom = Stack()
