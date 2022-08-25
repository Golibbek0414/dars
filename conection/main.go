package main
import (
	"io/ioutil"
	"fmt"
	"sync"
	"time"
	//"io"
	//"os"
)

func main() {
	file1, err := ioutil.ReadFile("text1.txt")
	if err !=nil {
		panic(err)
	}
	fmt.Println(string(file1))
	file2, err :=ioutil.ReadFile("text2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file2))
	file3, err :=ioutil.ReadFile("text3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file3))
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		sum(string(file1),time.Second*2)
	}()

	go func() {
		defer wg.Done()
		sum(string(file2),time.Second*2)
	}()
	
	go func() {
		defer wg.Done()
		sum(string(file3),time.Second*2)
	}()
	wg.Wait()
	fmt.Println("Successfully completed all goroutines")

}
func sum(x string,duration time.Duration) {
	time.Sleep(duration)
	sum := 0
	for i := 0; i <= len(x); i++{
		sum += i
	}
		fmt.Println(sum)
}
