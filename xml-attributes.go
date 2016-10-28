package main

type XmlNodeAttributes []*Attribute

func (s *XmlNodeAttributes) GetHash() uint64 {
	hashString := ""
	for _, item := range *s {
		hashString += GetHashString(item.GetHash())
	}
	return hash(hashString)
}
