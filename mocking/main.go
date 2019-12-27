package main

import "github.com/paunik/learn-go-with-test/mocking/countdown"

import "os"

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
