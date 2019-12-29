package countdown

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// DefaultSleeper struct
type DefaultSleeper struct{}

// Sleep sleeps time
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown counts down and writes output to io.Writer
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}
