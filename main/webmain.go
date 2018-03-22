package main

import (
	"fmt"
	"net/http"
)

type SimpleHanle struct{}

func (*SimpleHanle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello goloader!"))
}

func main() {

	sh := &SimpleHanle{}
	mux := http.NewServeMux()
	mux.Handle("/", sh)

	err := http.ListenAndServe(":9090", mux)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("main end line ...")
}
