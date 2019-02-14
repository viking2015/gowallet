package main

import (
	"fmt"
	// "unsafe"
	"sync/atomic"
	// "time"
	// "errors"
	// "flag"
	"context"
	// "runtime"
	// "runtime/debug"
	"strings"
	"sync"
	"time"
)

func test3forstrings() {
	var builder1 strings.Builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	fmt.Println()

	// 示例2。
	fmt.Println("Grow the builder ...")
	builder1.Grow(10)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	fmt.Println()

	// 示例3。
	fmt.Println("Reset the builder ...")
	builder1.Reset()
	fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String())
}
func test3syncpool() {
	// defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}
	v1 := pool.Get()
	fmt.Println(v1)
	pool.Put(100)
	pool.Put(200)
	pool.Put(300)
	v2 := pool.Get()
	fmt.Println(v2)
	// runtime.GC()
	time.Sleep(3 * time.Minute)
	v3 := pool.Get()
	fmt.Println(v3)

	pool.New = nil
	v4 := pool.Get()
	fmt.Println(v4)
}
func someHandler() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doStuff(ctx)
	cancel()
	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	// cancel()
	node2 := context.WithValue(ctx, "xx", "yy")
	fmt.Println(node2.Value("xx"))
}
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
			fmt.Println("default")
		}
	}
}
