package main

import (
	"time"
)

func runAtInterval(what func(), delay time.Duration) {
	ticker := time.NewTicker(delay)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				what()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func main() {
	runAtInterval(updateAndApplyFromGithub, 5*time.Minute)
}
