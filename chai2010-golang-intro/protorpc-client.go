// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main // OMIT

import (
	"fmt" // OMIT
	"log" // OMIT
	// OMIT
	"code.google.com/p/goprotobuf/proto"
	"github.com/chai2010/talks/chai2010-golang-intro/arith.pb" // HL
)

func main() {
	stub, client, err := arith.DialArithService("tcp", "127.0.0.1:1984") // HL
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close() // HL

	args := arith.ArithRequest{A: proto.Int32(7), B: proto.Int32(8)}
	reply := arith.ArithResponse{}
	if err = stub.Multiply(&args, &reply); err != nil { // HL
		log.Fatal(err)
	}
	fmt.Printf("stub.Multiply: 7*8 = %d\n", reply.GetVal())
}
