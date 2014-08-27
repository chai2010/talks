// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt" // OMIT
	"log" // OMIT
	// OMIT
	"code.google.com/p/goprotobuf/proto"
	hello "github.com/chai2010/talks/chai2010-protorpc-intro/hello.pb" // HL
)

func main() {
	client, _, err := hello.DialHelloService("tcp", "127.0.0.1:9527") // HL
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close() // HL

	in := &hello.HelloRequest{Msg: proto.String("Gopher")}
	out := &hello.HelloResponse{}
	if err = client.Hello(in, out); err != nil { // HL
		log.Fatal(err)
	}
	fmt.Println(out.GetMsg())
}
