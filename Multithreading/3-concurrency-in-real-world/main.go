package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("You are the visitor number %d", number)))
	})

	http.ListenAndServe(":3000", nil)
}
