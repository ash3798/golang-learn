package main

import (
	"fmt"
)

func goroutineDemo() {
	//create a channel , will be used to send signal to go routine
	quit := make(chan bool)

	go func() {
		//add checks to see handle the signal to quit
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("still running")
				//we can add sleep here to do periodic check
			}
		}
	}()

	quit <- true // sending signal to channel to stop goroutine
}
