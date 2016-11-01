// +build ignore

package main

import "time"

func main() {
	go println("你好, 并发!")
	time.Sleep(time.Second) // HL
}
