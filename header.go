package main

type Header1 struct {
	OGHeader  OGHeader `xml:"OGHeader"`
	MessageID string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing MessageID"`
	RelatesTo string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing RelatesTo"`
	To        string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing To"`
	Security  Security `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
}

type OGHeader struct {
	transactionID string `xml:"http://webservices.micros.com/og/4.3/Core/ transactionID,attr"`
	timeStamp     string `xml:"http://webservices.micros.com/og/4.3/Core/ timeStamp,attr"`
	primaryLangID string `xml:"http://webservices.micros.com/og/4.3/Core/ primaryLangID,attr"`
}

type Origin struct {
	entityID   string
	systemType string
}

type Destination struct {
	entityID   string
	systemType string
}

type Security struct {
	Id        string    `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Id,attr"`
	Timestamp Timestamp `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Timestamp"`
}

type Timestamp struct {
	Id      string `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Id,attr"`
	Created string
	Expires string
}
