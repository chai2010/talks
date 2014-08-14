// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdio.h>

#include "./arith.pb.h" // HL

int main() {
	google::protobuf::rpc::Client client("127.0.0.1", 1984); // HL
	arith::ArithService::Stub stub(&client); // HL

	arith::ArithRequest args;
	arith::ArithResponse reply;
	args.set_a(7);
	args.set_b(8);

	auto err = stub.multiply(&args, &reply); // HL
	if(!err.IsNil()) { // HL
		fprintf(stderr, "stub.multiply: %s\n", err.String().c_str());
		return -1;
	}
	printf("7*8 = %d\n", reply.val());
	return 0;
}
