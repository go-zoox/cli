# Spinner

The `Spinner` component displays a loading spinner with a message.

## Basic Usage

```go
package main

import (
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	bar := loading.Spinner("Loading...")
	
	// Simulate work
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

	bar.Finish()
}
```

## Methods

### Spinner Methods

| Method | Description |
|--------|-------------|
| `Add(n int)` | Add progress (increments the counter) |
| `Finish()` | Finish the spinner and show completion message |

## Examples

### Simple Spinner

```go
bar := loading.Spinner("Processing...")
// Do some work
time.Sleep(2 * time.Second)
bar.Finish()
```

### Spinner with Progress

```go
bar := loading.Spinner("Downloading...")
for i := 0; i < 100; i++ {
	// Simulate download progress
	bar.Add(1)
	time.Sleep(50 * time.Millisecond)
}
bar.Finish()
```

### Complete Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	// Simulate file processing
	bar := loading.Spinner("Processing files...")
	
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, file := range files {
		// Simulate processing each file
		time.Sleep(500 * time.Millisecond)
		bar.Add(1)
		fmt.Printf("Processed: %s\n", file)
	}
	
	bar.Finish()
	fmt.Println("All files processed!")
}
```
