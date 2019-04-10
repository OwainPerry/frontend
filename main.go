package main

import (
	"fmt"
    "net/http"
    "io/ioutil"
)

func main() {

    fmt.Println("Starting the frontend application... 1.0")

    servicea := "no results"
    serviceb := "no results"
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        response, err := http.Get("http://servicea:8080/")
        if err != nil {
            fmt.Printf("The HTTP request failed servicea with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            servicea = string(data)
        }
        response1, err1 := http.Get("http://serviceb:8080/")
        if err1 != nil {
            fmt.Printf("The HTTP request failed serviceb with error %s\n", err)
        } else {
            data1, _ := ioutil.ReadAll(response1.Body)
            serviceb = string(data1)
        }
        fmt.Fprintf(w, "<html><body>")
        fmt.Fprintf(w, "first service: %s <br />",servicea)
        fmt.Fprintf(w, "second service: %s <br />",serviceb)
        fmt.Fprintf(w, "</body></html>")
	})
	http.ListenAndServe(":9090", nil)
}
