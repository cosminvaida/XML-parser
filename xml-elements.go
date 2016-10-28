package main

type XmlElements []*XMLElement

func (s *XmlElements) GetHash() uint64 {
	hashString := ""
	for _, item := range *s {
		hashString += GetHashString(item.GetHash())
	}
	return hash(hashString)
}
