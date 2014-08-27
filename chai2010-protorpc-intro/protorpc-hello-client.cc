// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "./hello.pb/hello.pb.h" // HL

#include <google/protobuf/rpc/rpc_server.h>
#include <google/protobuf/rpc/rpc_client.h>

int main() {
	::google::protobuf::rpc::Client client("127.0.0.1", 9527); // HL

	hello::HelloService::Stub helloStub(&client); // HL

	::hello::HelloRequest helloArgs;
	::hello::HelloResponse helloReply;
	::google::protobuf::rpc::Error err;

	helloArgs.set_msg("Hello Protobuf-RPC!");
	err = helloStub.Hello(&helloArgs, &helloReply); // HL
	if(!err.IsNil()) {
		fprintf(stderr, "echoStub.Echo: %s\n", err.String().c_str());
		return -1;
	}
	printf("out = %s\n", helloReply.msg().c_str());
	return 0;
}
