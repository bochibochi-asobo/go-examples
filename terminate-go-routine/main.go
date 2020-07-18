package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

// 非同期に実行するサンプル関数
func doSomethingWithChannel(queue chan string) {
	for {
		whatToSay, more := <-queue

		if more {
			fmt.Println("zzz...")
			time.Sleep(1)
			fmt.Println(whatToSay)
		} else {
			fmt.Println("worker exit")
			wg1.Done()
			return
		}
	}
}

// チャンネルを使って goroutine を停止させる例
func terminateWithChannel() {
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg1.Add(1)
		go doSomethingWithChannel(queue)
	}

	queue <- "(1) Good morning!"
	queue <- "(2) Can you hear me?"
	queue <- "(3) Hoge"
	queue <- "(4) Fuga"

	close(queue)
	wg1.Wait()
}

// 非同期に実行するサンプル関数
func doSomethingWithContext(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exit")
			wg2.Done()
			return
		case whatToSay := <-queue:
			fmt.Println("zzz...")
			time.Sleep(1)
			fmt.Println(whatToSay)
		}
	}
}

// Context を使って goroutine を停止させる例
func terminateWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go doSomethingWithContext(ctx, queue)
	}

	queue <- "(1) Good evening!"
	queue <- "(2) Hi, can you hear me?"
	queue <- "(3) Piyo"
	queue <- "(4) aaa"

	cancel()
	wg2.Wait()
}

func main() {
	fmt.Println("Start channel example....")
	terminateWithChannel()

	fmt.Println("\nStart context example....")
	terminateWithContext()
}
