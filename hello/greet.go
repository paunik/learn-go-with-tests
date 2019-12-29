package hello

import (
	"fmt"
	"io"
)

// Greet hello function
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
