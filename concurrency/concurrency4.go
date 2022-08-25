package main
import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(done <-chan stuct{} ) <-chan int {
	rand.Seed(time.Now().UnixNano())
	n :=rand.Intn(20)+5
	fmt.Println("n=",n)
	c :=make(chan int)
	go func() {
		defer close(c)
		for i:=0;i<n; i++ {
			time.Sleep(time.Second)
			fmt.Println("attempting to send random number to c")
			select {
			case c <- rand.Intn(25):
			case <-done:
				fmt.Println("goroutine exited")
				return
			}
		}
	}()

	return c
}

func main() {
	done :=make(chan struct{})
	for num := range generateNumbers(done) {
		fmt.Println(num)
		if num > 10 {
			close(done)
			break
		}
	}
	time.Sleep(time.Second)
	fmt.Println("main finished")
}




