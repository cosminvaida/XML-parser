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

	xml, err := XMLDocument(file)
	res1, err := xml.SelectNodes("/Envelope/Body/MeetingAvailabilityResponse/AvailableProperties/FunctionSpace/SetupStyles")
	res, err := xml.GetElementsByPath("/Envelope/Body/MeetingAvailabilityResponse/AvailableProperties/FunctionSpace/SetupStyles")
	res2, err := xml.GetElementsByAttribute("avaliablityString")
	fmt.Println(xml)
	fmt.Println(res)
	fmt.Println(res1)
	fmt.Println(res2)

	fmt.Println(err)

}
