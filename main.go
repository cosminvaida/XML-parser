package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	strapsFilePath, err := filepath.Abs("file.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the straps.xml file
	file, err := os.Open(strapsFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	xml, err := XMLDocumentFromFile(file)
	res, err := xml.search("/Envelope/Header/RelatesTo")
	fmt.Println(xml)
	fmt.Println(res)
	fmt.Println(err)

	//XmlNodeList xnList = xml.SelectNodes("/Names/Name");
	// f, err23 := NewDocument(file)
	// fmt.Println(f)
	// fmt.Println(err23)

	// //GetElements("HotelReference")

	// Map(Envelope{})

}
