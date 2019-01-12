package main

import "fmt"
import "time"

func doSleep() {
	fmt.Println("hello world!")
	time.Sleep(time.Duration(3) * time.Second)
}
func main() {

	bc := make(chan bool)
	go func() {
		select {
		case <-bc:
			fmt.Println("has value")
			close(bc)
		}
	}()
	time.Sleep(time.Second * 10)
	bc <- true
	<-bc
	fmt.Println("END")
}
