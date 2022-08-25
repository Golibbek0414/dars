package main

import (
	"fmt"
)

func generator(words []string) <-chan string {
	c :=make(chan string)
	go func() {
		defer close(c)
		for _,word := range words {
			c <-word
		}
	}()
	return c
}

func main() {
	words :=[]string{
		"word","word","hello","world","word","ball","apple","apple",
	}
	for word :=range removeRepeated(generator(words)) {
		fmt.Println(word)
	}
	
}

func removeRepeated(inputStream <-chan string) <-chan string {
	outputStream :=make(chan string)
	go func() {
		defer close(outputStream)
		var oldVal string

		for val := range inputStream {
			if val !=oldVal {
				outputStream <-val
			}
			oldVal=val
		}
	}()
	return outputStream
}