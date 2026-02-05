# Loading Examples

This page contains examples of using loading components.

## Spinner

```go
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
```

## Progress

```go
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
```
