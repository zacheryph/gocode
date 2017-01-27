package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

type Config struct {
	Status  bool
	Host    string
	Service Embedded

	File map[string]File
}

type Embedded struct {
	Name  string
	Limit int64
}

type File struct {
	// Name   string  // Would like to see id/name embedded
	Status string
	Limit  int64
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: hcltest file.conf")
		return
	}

	d, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	c := Config{}

	err = hcl.Decode(&c, string(d))
	if err != nil {
		fmt.Println("HCL", err)
		return
	}

	fmt.Println(c)
}
