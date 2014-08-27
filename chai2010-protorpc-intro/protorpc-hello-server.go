// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log" // OMIT
	// OMIT
	"code.google.com/p/goprotobuf/proto"
	hello "github.com/chai2010/talks/chai2010-protorpc-intro/hello.pb"  // HL
)

type myHello struct{}

func (p *myHello) Hello(in *hello.HelloRequest, out *hello.HelloResponse) error { // HL
	out.Msg = proto.String("myHello.Hello: " + in.GetMsg())
	log.Println("myHello.Hello:", in.GetMsg())
	return nil
}

func main() {
	go hello.ListenAndServeHelloService("tcp", ":9527", new(myHello)) // HL
	log.Println("Start HelloService on :9527 ...")
	<-make(chan bool)
}
