package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })
	// http.HandleFunc("/increment", incrementCounter)
	http.Handle("/", http.FileServer(http.Dir("./html5up-identity")))
        fmt.Println("Listening on 0.0.0.0:80")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))

}
