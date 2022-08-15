package loading

import "github.com/schollz/progressbar/v3"

// Progress is a progress bar
func Progress(total int64, description ...string) *progressbar.ProgressBar {
	return progressbar.Default(total, description...)
}

// ProgressBytes is a progress bar for bytes
func ProgressBytes(total int64, description ...string) *progressbar.ProgressBar {
	return progressbar.DefaultBytes(total, description...)
}
