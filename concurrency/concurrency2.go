package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums :=sendRandomNumbers()
	for num := range nums {
		fmt.Println(num)
	}
}

func sendRandomNumbers() chan int {
	rand.Seed(time.Now().UnixNano())
	n :=rand.Intn(10) +1
	nums :=make(chan int)
	go func() {
		defer close(nums)
		for i :=0; i<n; i++ {
			nums <- rand.Intn(80)

		}
	}()
	return nums
}