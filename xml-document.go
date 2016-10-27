package main

import "fmt"

func XMLDocument() (*XMLElement, error) {
	return &XMLElement{}, nil
}

func SelectNodes(nodePath string) (XmlNodeSplice, error) {

	x := XmlNodeSplice{}
	fmt.Println(x)

	return nil, nil
}

func (node *XMLElement) search(nodePath string) (*XMLElement, error) {

	if node.Path == nodePath {
		return node, nil
	}
	for _, childNode := range node.Childs {
		nodesChild, err := childNode.search(nodePath)
		if err != nil {
			return nil, err
		}
		if nodesChild != nil {
			return nodesChild, nil
		}

	}
	return nil, nil
}
