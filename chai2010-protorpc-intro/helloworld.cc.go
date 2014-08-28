// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main // HL

/*
#define main myMain // OMIT

// CC START // OMIT
#include <stdio.h> // HL

int main() {
	printf("Hello, World!\n"); // HL
	return 0;
}
// CC END // OMIT
*/
import "C"

func main() {
	C.myMain()
}
