package main

import (
	"fmt"
    "net/http"
    "io/ioutil"
)

func main() {

    fmt.Println("Starting the frontend application...")

    servicea := "no results"
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        response, err := http.Get("http://localhost:8080/")
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            servicea = string(data)
        
        }
        
        fmt.Fprintf(w, "first service: %s ",servicea)
	})
	http.ListenAndServe(":9090", nil)
}
