package main

import "errors"

func (node *XMLElement) SelectNodes(nodePath string) (*XmlElements, error) {
	hashKey := hash(nodePath)
	if value, ok := XMLMap[hashKey]; ok {
		return &value, nil
	} else {
		return nil, errors.New("The path was not found")
	}
}

func (node *XMLElement) GetElementsByPath(nodePath string) (*XmlElements, error) {
	return node.SelectNodes(nodePath)
}

func (node *XMLElement) GetElementsByName(name string) (*XmlElements, error) {

	fn := func(node *XMLElement) bool {
		return node.Name == name
	}
	nodes, err := node.search(fn)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (node *XMLElement) GetElementsByAttribute(attribute string) (*XmlElements, error) {
	fn := func(node *XMLElement) bool {
		for _, nodeAttribute := range *node.Attributes {
			if nodeAttribute.Key == attribute {
				return true
			}
		}
		return false
	}
	nodes, err := node.search(fn)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (node *XMLElement) search(fn func(*XMLElement) bool) (*XmlElements, error) {
	nodes := XmlElements{}
	if fn(node) {
		nodes.append(*node)
	}
	if node.Childs == nil {
		return &nodes, nil
	}
	for _, childNode := range *node.Childs {
		childResult, err := childNode.search(fn)
		if err != nil {
			return nil, err
		}
		if childResult != nil && len(*childResult) > 0 {
			nodes.append(childResult)
		}
	}
	return &nodes, nil
}
