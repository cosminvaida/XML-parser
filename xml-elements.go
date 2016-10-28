package main

type XmlElements []*XMLElement

func (s *XmlElements) GetHash() uint64 {
	hashString := ""
	for _, item := range *s {
		hashString += GetHashString(item.GetHash())
	}
	return hash(hashString)
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
