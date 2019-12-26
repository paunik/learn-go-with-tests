package main

import (
	"net/http"

	"github.com/paunik/learn-go-with-test/hello"
)

func main() {
	// fmt.Println(hello.Hello("From Main", "English"))
	// hello.Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingHandler))
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	hello.Greet(w, "World!")
}
