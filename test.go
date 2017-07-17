package main

import "fmt"
import "time"

func doSleep() {
	fmt.Println("hello world!")
	time.Sleep(time.Duration(3) * time.Second)
}
func main() {
	/*
		t := time.NewTicker(time.Duration(10) * time.Millisecond)
		for {
			//t := time.After(time.Duration(10) * time.Millisecond)
			select {
			case <-t.C:
				doSleep()
			}
		}
	*/

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
