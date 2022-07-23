package main

import (
	"time"
)

func runAtInterval(what func(), delay time.Duration) {
	for range time.Tick(delay) {
		what()
	}
}

func main() {
	runAtInterval(updateAndApplyFromGithub, 5*time.Minute)
}
