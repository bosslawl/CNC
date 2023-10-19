package Animations

import (
	"time"
	"io"
)

var Spinner bool

func NewSpinner(channel io.Writer) {

	Spinner = true

	Frames := []string{
		"|", "/", "-", "\\", "|", "/", "-", "\\",
	}

	for _, I := range Frames {
		channel.Write([]byte("\r" + I))

		if !Spinner {
			channel.Write([]byte("\r\n"))
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}