// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Get(filename string) (data []byte, err error) { // HL
	return ioutil.ReadFile(filename) // HL
}

func main() {
	data, err := Get("./chai2010-protorpc-intro/local-get.go") // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
