package main

import (
	"hash/fnv"
	"strconv"
)

func mapNode(key string, node *XMLElement) {
	hashKey := hash(key)
	if value, ok := XMLMap[hashKey]; ok {
		value = append(value, node)
		XMLMap[hashKey] = value
	} else {
		newValue := XmlElements{}
		newValue = append(newValue, node)
		XMLMap[hashKey] = newValue
	}
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func (s *XmlElements) append(items interface{}) {
	if s == nil {
		*s = *oo()
	}
	switch items.(type) {
	case *XmlElements:
		XmlElements, _ := items.(*XmlElements)
		for _, item := range *XmlElements {
			*s = append(*s, item)
		}
	case XMLElement:
		xmlElement, _ := items.(XMLElement)
		*s = append(*s, &xmlElement)
	}
}

func oo() *XmlElements {
	XmlElements := XmlElements{}
	return &XmlElements
}

func (s *XmlNodeAttributes) append(items interface{}) {
	switch items.(type) {
	case *XmlNodeAttributes:
		xmlNodeAttributes, _ := items.(*XmlNodeAttributes)
		for _, item := range *xmlNodeAttributes {
			*s = append(*s, item)
		}
	case Attribute:
		attribute, _ := items.(Attribute)
		*s = append(*s, &attribute)
	}
}

func (s *XmlNodeAttributes) contains(item Attribute) bool {
	for _, attribute := range *s {
		if attribute.GetHash() == item.GetHash() {
			return true
		}
	}
	return false
}

func GetHash(s string) uint64 {
	return hash(s)
}

func ToString(s uint64) string {
	return strconv.FormatUint(s, 10)
}

func GetHashString(in interface{}) string {
	switch in.(type) {
	case string:
		item := in.(string)
		return ToString(hash(item))
	case *XmlNodeAttributes:
		items, _ := in.(*XmlNodeAttributes)
		return ToString(items.GetHash())
	case *XmlElements:
		items, _ := in.(*XmlElements)
		return ToString(items.GetHash())
	}
	return ""
}
