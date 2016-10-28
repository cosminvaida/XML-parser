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
