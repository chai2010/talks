// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "./hello.pb/hello.pb.h" // HL

#include <google/protobuf/rpc/rpc_server.h>
#include <google/protobuf/rpc/rpc_client.h>

struct HelloService: public hello::HelloService { // HL
	virtual const ::google::protobuf::rpc::Error Hello(
		const ::service::HelloRequest* in, ::service::HelloResponse* out
	) {
		out->set_msg(in->msg());
		printf("HelloService.Hello: in:\"%s\", out:\"%s\"\n",
			in->msg().c_str(), out->msg().c_str()
		);
		return ::google::protobuf::rpc::Error::Nil(); // HL
	}
};

int main() {
	::google::protobuf::rpc::Server server;
	server.AddService(new HelloService, true); // HL
	server.BindAndServe(9527);// HL
	return 0;
}
