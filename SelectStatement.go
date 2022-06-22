package main

import "fmt"

func main() {
	c1 := make(chan int, 1024) //create a channel
	c2 := make(chan int, 1024)
	c3 := make(chan int, 1024)
	c1 <- 100 //sen data to a channel
	c2 <- 200
	c3 <- 300
	select { // randomly execute one case
	case <-c1: //receive data from a channel
		fmt.Println("Receive data from c1")
	case <-c2:
		fmt.Printf("Receive data from c2")
	case <-c3:
		fmt.Println("Receive data from c3")
	default:
		fmt.Println("Not receive data yet")

	}

}
