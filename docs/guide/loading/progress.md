# Progress

The `Progress` component displays a progress bar with a total count.

## Basic Usage

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

## Methods

### Progress Methods

| Method | Description |
|--------|-------------|
| `Add(n int)` | Add progress by n units |
| `Set(n int)` | Set progress to n |
| `Finish()` | Finish the progress bar |

## Examples

### Simple Progress Bar

```go
bar := loading.Progress(100)
for i := 0; i < 100; i++ {
	bar.Add(1)
	time.Sleep(10 * time.Millisecond)
}
```

### Progress with Custom Steps

```go
total := 50
bar := loading.Progress(total)

for i := 0; i < total; i++ {
	// Do work
	time.Sleep(100 * time.Millisecond)
	bar.Add(1)
}
```

### Progress with Set

```go
bar := loading.Progress(100)

// Set progress to 50%
bar.Set(50)

// Add 25 more
bar.Add(25)

// Finish
bar.Finish()
```

### Complete Example

```go
package main

import (
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	// Simulate downloading 100 files
	totalFiles := 100
	bar := loading.Progress(totalFiles)

	for i := 0; i < totalFiles; i++ {
		// Simulate download
		time.Sleep(50 * time.Millisecond)
		bar.Add(1)
	}

	fmt.Println("\nDownload complete!")
}
```
