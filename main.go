package main

import (
	"fmt"
	"time"
)

func runAtInterval(what func(), delay time.Duration) {
	for range time.Tick(delay) {
		fmt.Println("tick")
		what()
	}
}

func main() {
	fmt.Println("started")
	runAtInterval(updateAndApplyFromGithub, 5*time.Minute)
}
