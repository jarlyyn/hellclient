package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"

	"golang.org/x/net/html/charset"
)

func main() {
	// flag.Parse()
	mcl := &Mcl{}
	if len(os.Args) != 2 {
		fmt.Println("usage: mclconvertor MCLFILENAME ")
	}
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	outpout := bytes.SplitN(data, []byte("\n"), 2)
	if len(outpout) == 2 {
		data = outpout[1]
	}
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&mcl)

	// err = xml.Unmarshal(data, &mcl)
	if err != nil {
		panic(err)
	}
	sd := mcl.ToScriptData()
	writer := bytes.NewBuffer(nil)
	err = toml.NewEncoder(writer).Encode(sd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(writer.Bytes()))
}
