package main

import (
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	bar := loading.Spinner("Loading...")
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

	bar.Finish()
}
