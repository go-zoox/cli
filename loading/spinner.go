package loading

import (
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

// Spinner is a progress bar
func Spinner(description ...string) *progressbar.ProgressBar {
	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}

	return progressbar.NewOptions64(
		-1,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
	)
}
