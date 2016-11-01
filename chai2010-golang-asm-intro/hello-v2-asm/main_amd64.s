// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h" // HL
#include "funcdata.h" // HL

// "Hello World!\n"
DATA  text<>+0(SB)/8,$"Hello Wo"
DATA  text<>+8(SB)/8,$"rld!\n"
GLOBL text<>(SB),NOPTR,$16 // HL

// func main()
TEXT ·main(SB), $16-0 // HL
	NO_LOCAL_POINTERS
	MOVQ $text<>+0(SB), AX // HL
	MOVQ AX, (SP)
	MOVQ $16, 8(SP)
	CALL fmt·Println(SB) // HL
	RET // HL
