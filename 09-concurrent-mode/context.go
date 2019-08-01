package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type key int

const outChannelKey = key(1)

func fetchResourceWithContext(ctx context.Context, server uint) {
	fmt.Printf("start fetching resource from server %d\n", server)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Printf("fetching from server %d has been cancelled\n", server)
	default:
		fmt.Printf("resource received from server %d\n", server)
		ctx.Value(outChannelKey).(chan<- uint) <- server
	}
}

func useContext(timeout time.Duration) {
	out := make(chan uint)
	ctx := context.WithValue(context.Background(), outChannelKey, chan<- uint(out))
	ctx, cancel := context.WithTimeout(ctx, timeout)

	go fetchResourceWithContext(ctx, 1)
	go fetchResourceWithContext(ctx, 2)

	select {
	case v := <-out:
		fmt.Printf("value returned: %d\n", v)
	case <-ctx.Done():
		fmt.Printf("both server timeout\n")
	}
	cancel()
	time.Sleep(3 * time.Second)
}
