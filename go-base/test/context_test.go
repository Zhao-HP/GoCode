package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	go GoroutineCtx(ctx)
	time.Sleep(4 * time.Second)
	cancel()

	time.Sleep(10 * time.Second)
	fmt.Println("TestContext exec done")
}

func GoroutineCtx(ctx context.Context) {

	select {
	case v := <-ctx.Done():
		fmt.Println("v", v)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("GoroutineCtx exec done")

}
