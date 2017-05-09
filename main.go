package main

import (
    "log"
    "net/http"  
    "encoding/json"
)

type Response struct {
    Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

    message := Response {
        Message: "Google I/O 17",
    }

    response, err := json.Marshal(message)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func main() {

   http.HandleFunc("/", handler) 

   log.Println("Listening on port: 8080")

   err := http.ListenAndServe(":8080", nil) 
   if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}