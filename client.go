package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()
	res := make(chan string)
	defer cancel()

	for range 10 {
		wg.Add(1)
		go foo(wg, ctx, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for val := range res {
		fmt.Println(val)
	}

}

func foo(wg *sync.WaitGroup, ctx context.Context, res chan<- string) {
	defer wg.Done()
	t := rand.Intn(10) + 1
	time.Sleep(time.Second * time.Duration(t))
	res <- "foo() completed"
	
}
