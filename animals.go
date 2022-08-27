package main

import (
	"fmt"
	"time"
)

func count(name string, tlen int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(tlen) * time.Millisecond)
		fmt.Printf("%sが%d匹\n", name, i+1)
	}
}

func main() {
	go count("カエル", 200)
	go count("アヒル", 100)
	time.Sleep(3000 * time.Millisecond) // mainルーチンに暫く寝てもらい、gorourtineが走れるようにする
}
