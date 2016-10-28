package main

type Attribute struct {
	NameSpace string
	Key       string
	Value     string
}

func (s *Attribute) GetHash() uint64 {

	hashString := GetHashString(s.Key) +
		GetHashString(s.Value) +
		GetHashString(s.NameSpace)
	return hash(hashString)
}
