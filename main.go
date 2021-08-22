package main

import (
	"bufio"
	"os"
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
	schedule(updateAndApplyFromGithub, 1*time.Minute)

	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
}
