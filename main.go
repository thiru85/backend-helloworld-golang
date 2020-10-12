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

    fmt.Fprintf(w, "%v", req.Header)
}

func helloworld(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello world\n")
}


func main() {

    http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
    http.HandleFunc("/thiru", thiru)
    http.HandleFunc("/helloworld", helloworld)

    http.ListenAndServe(":8080", nil)
}