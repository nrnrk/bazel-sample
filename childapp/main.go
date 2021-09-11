package main

import (
	"fmt"
	"net/http"

	"github.com/nrnrk/bazel-sample/library/strlib"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strlib.Concat("I am", " child"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
