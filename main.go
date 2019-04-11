package main

import (
	"fmt"
    "net/http"
    "io/ioutil"
    "github.com/tkanos/gonfig"
)

type Configuration struct {
    Port              int
    ServiceA   string
    ServiceB string
}

func main() {

    configuration := Configuration{}
    conferr := gonfig.GetConf("config/config.json", &configuration)
    if conferr != nil {
        fmt.Println("Failed to read config file")
    }

    fmt.Printf("Starting the frontend application... 1.3 port : %d",configuration.Port)

    servicea := "no results"
    serviceb := "no results"
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        response, err := http.Get(configuration.ServiceA)
        if err != nil {
            fmt.Printf("The HTTP request failed servicea with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            servicea = string(data)
        }
        response1, err1 := http.Get(configuration.ServiceB)
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
	http.ListenAndServe(fmt.Sprintf(":%d",configuration.Port), nil)
}
