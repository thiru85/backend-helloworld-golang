package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func thiru(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "thiru:%v\n", req.Header)
}



func main() {

    http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/thiru", thiru)

    http.ListenAndServe(":8080", nil)
}