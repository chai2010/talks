// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdio.h>
#include <string>

bool Get(const std::string& filename, std::string* data) { // HL
	FILE* fp = fopen(name, "rb");
	if(!fp) return false;

	fseek(fp, 0, SEEK_END);
	data->resize(ftell(fp));

	rewind(fp);
	int n = fread((void*)data->data(), 1, data->size(), fp); // HL
	fclose(fp);
	return (n == data->size());
}

int main() {
	std::string data;
	if(!Get("./chai2010-protorpc-intro/helloworld.go")) {  // HL
		fprintf(stderr, "open file failed!\n");
		return -1;
	}
	printf("%s", data.c_str());
	return 0;
}
