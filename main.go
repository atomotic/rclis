package main

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer gzipReader.Close()

	decoder := xml.NewDecoder(gzipReader)

	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	for _, record := range response.ListRecords.Record {
		for _, item := range record.Metadata.DIDL.Item.Item {
			objectType := item.Descriptor.Statement.ObjectType
			if objectType == "info:eu-repo/semantics/humanStartPage" || objectType == "info:eu-repo/semantics/objectFile" {
				fmt.Println(item.Component.Resource.Ref)
			}
		}
	}
}
