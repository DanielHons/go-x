package cmdx

import (
	"fmt"
	"os"
)

// Must fatals with the optional message if err is not nil.
func Must(err error, message string, args ...interface{}) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	os.Exit(1)
}
