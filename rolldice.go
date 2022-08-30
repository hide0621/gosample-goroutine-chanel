// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type dice struct {
// 	val int
// }

// func rollDice(d dice, c chan int) {
// 	rand.Seed(time.Now().UnixNano())
// 	for {
// 		time.Sleep(10 * time.Millisecond)
// 		v := rand.Intn(10)
// 		if d.val == v {
// 			fmt.Printf("マッチしました\n")
// 			break //一つのgoroutine（スレッド）の終了を意味する
// 		} else {
// 			fmt.Printf("%dか...%dではないな\n", v, d.val)
// 		}
// 	}
// 	c <- d.val
// }

// func main() {
// 	d1 := dice{2}
// 	d2 := dice{6}

// 	c := make(chan int)

// 間違いのアルゴリズム
// 	//x, y := <-c, <-c //goroutineが走ってチャネルが値を受信する前なのでデッドロックされる（空のチャネル）

// 	go rollDice(d1, c)
// 	go rollDice(d2, c) //上のgoroutineも走りつつ、このgoroutineも走る

// 正解のアルゴリズム
// 	x, y := <-c, <-c // チャネルの送受信を入れることでgoroutineが走る余地を与えている(goroutineが二つとも走り終わったら、値が送信される)

// 	fmt.Printf("%dが出ました\n", x)
// 	fmt.Printf("%dが出ました\n", y)
// }
