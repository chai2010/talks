// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Get(filename string) (data []byte, err error) {
	return ioutil.ReadFile(filename) // HL
}

func Call(method string, filename string) (data []byte, err error) { // HL
	switch method {
	case "Get":
		return Get(filename) // HL
	default:
		return nil, fmt.Errorf("Unknown method: %s", method)
	}
}

func main() {
	data, err := Call("GetX", "./chai2010-protorpc-intro/helloworld.go") // fixit // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
