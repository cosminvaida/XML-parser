package main

type XMLElement struct {
	Name       string
	InnerValue string
	Attributes *XmlNodeAttributes
	Childs     *XmlElements
	Path       string
	PathHash   uint64
}

func (s *XMLElement) GetHash() uint64 {
	hashString := GetHashString(s.Name) +
		GetHashString(s.InnerValue) +
		GetHashString(s.Path) +
		GetHashString(s.Attributes) +
		GetHashString(s.Childs) +
		GetHashString(s.Path)
	return hash(hashString)
}
