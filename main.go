package main

import (
	"time"
)

func schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func main() {
	schedule(updateAndApplyFromGithub, 5*time.Minute)

	for true {
		time.Sleep(500)
	}
}
