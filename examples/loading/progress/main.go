package main

import (
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	bar := loading.Progress(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}
