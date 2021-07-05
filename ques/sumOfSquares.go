package main

import "fmt"

func SumOfSquares(c, quit chan int) {

	// your code here
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}

}

func FindSumOfSquares() int {

	mychannel := make(chan int)

	quitchannel := make(chan int)

	sum := 0

	go func() {
		for i := 1; i < 6; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
	return sum
}
