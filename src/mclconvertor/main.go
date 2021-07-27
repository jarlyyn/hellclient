package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"

	"golang.org/x/net/html/charset"
)

func ConvertWrold() {
	mcl := &Mcl{}
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("usage: mclconvertor MCLFILENAME ")
	}
	filename := args[0]
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
	sd := mcl.ToWorldData()
	writer := bytes.NewBuffer(nil)
	err = toml.NewEncoder(writer).Encode(sd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(writer.Bytes()))
}
func ConvertScript() {
	mcl := &Mcl{}
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("usage: mclconvertor MCLFILENAME ")
	}
	filename := args[0]
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
func main() {
	// flag.Parse()
	var world bool
	flag.BoolVar(&world, "world", false, "Convert mcl to world toml")
	flag.Parse()
	if world {
		ConvertWrold()
	} else {
		ConvertScript()
	}
}
