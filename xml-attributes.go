package main

type XmlNodeAttributes []*Attribute

func (s *XmlNodeAttributes) GetHash() uint64 {
	hashString := ""
	for _, item := range *s {
		hashString += GetHashString(item.GetHash())
	}
	return hash(hashString)
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
