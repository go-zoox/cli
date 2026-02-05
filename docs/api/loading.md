# Loading API

Loading components for progress indication.

## Spinner

Create a loading spinner.

```go
func Spinner(message string) *Spinner
```

### Spinner Methods

#### Add

Add progress.

```go
func (s *Spinner) Add(n int)
```

#### Finish

Finish the spinner.

```go
func (s *Spinner) Finish()
```

## Progress

Create a progress bar.

```go
func Progress(total int) *Progress
```

### Progress Methods

#### Add

Add progress.

```go
func (p *Progress) Add(n int)
```

#### Set

Set progress value.

```go
func (p *Progress) Set(n int)
```

#### Finish

Finish the progress bar.

```go
func (p *Progress) Finish()
```
