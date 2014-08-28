// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main // HL

/*
#define main myMain // OMIT

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* Get(const char* filename, int* size) { // HL
	FILE* fp = fopen(filename, "rb");
	if(!fp) return NULL;

	fseek(fp, 0, SEEK_END);
	*size = ftell(fp) + 1;
	char* data = (char*)calloc(*size, 1);

	rewind(fp);
	fread((void*)data, 1, *size, fp); // HL
	fclose(fp);
	return data;
}

// CC START // OMIT
int Call(const char* method, const char* filename, char** data, int* size) { // HL
	if(strcmp(method, "Get") == 0) {
		*data = Get(filename, size); // HL
		return 1;
	}
	return 0;
}

int main() {
	int size;
	char* data;
	if(!Call("Get", "./chai2010-protorpc-intro/local-get-proxy.cc.go", &data, &size)) {  // HL
		fprintf(stderr, "open file failed!\n");
		return -1;
	}
	printf("%s", data);
	return 0;
}
// CC END // OMIT
*/
import "C"

func main() {
	C.myMain()
}
