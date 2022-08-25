package main
import "fmt"
func one(c1 chan string) {
	for i := 0; i < 5; i++ {
	  c1 <- "Ahvollar qalay"
	}
	close(c1)
   }
   func two(c2 chan string) {
	for i := 0; i < 5; i++ {
	  c2 <- "Nima gaplar ishlar zo'rmi"
	}
	close(c2)
   }
   func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go one(c1)
	go two(c2)
	for {
	  select {
		case msg, ok := <-c1:
		fmt.Println(msg)
		if !ok {
		  c1 = nil
		}
		case msg, ok := <-c2:
		fmt.Println(msg)
		if !ok {
		  c2 = nil
		}
	  }
	  if c1 == nil && c2 == nil {
		break
	  }
	}
}



// package main
// import (
//  "fmt"
//  "log"
//  "net/http"
//  "github.com/gorilla/mux"
// )
// func main() {
//  r := mux.NewRouter()
//  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Eyyy to'xta seniy")
//  })
//  log.Print("Server starting at localhost:4444")
//  http.ListenAndServe(":4444", r)
// }