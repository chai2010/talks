// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main // OMIT

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/chai2010/talks/chai2010-golang-intro/arith.pb" // HL
	"log" // OMIT
)

type Arith int // HL

func (t *Arith) Multiply(args *arith.ArithRequest, reply *arith.ArithResponse) error { // HL
	defer log.Printf("Arith.Multiply: args = %v, reply = %v\n", args, reply)
	reply.Val = proto.Int32(args.GetA() * args.GetB())
	return nil
}
func (t *Arith) Divide(args *arith.ArithRequest, reply *arith.ArithResponse) error { // HL
	defer log.Printf("Arith.Multiply: args = %v, reply = %v\n", args, reply)
	reply.Quo = proto.Int32(args.GetA() / args.GetB())
	reply.Rem = proto.Int32(args.GetA() % args.GetB())
	return nil
}

func main() {
	log.Printf("ProtobufRPC Server running ...\n")
	arith.ListenAndServeArithService("tcp", ":1984", new(Arith)) // HL
}
