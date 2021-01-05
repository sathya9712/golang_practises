package main

import (
"log"
"net/http"
)

func localhome(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
switch r.Method {
case "GET":
 w.WriteHeader(http.StatusOK)
 w.Write([]byte(`{"message": "hello sathya statusok"}`))
case "POST":
 w.WriteHeader(http.StatusCreated)
 w.Write([]byte(`{"message": "hello sathya status created"}`))
case "PUT":
 w.WriteHeader(http.StatusAccepted)
 w.Write([]byte(`{"message": "hello sathya status accepted"}`))
case "DELETE":
 w.WriteHeader(http.StatusOK)
 w.Write([]byte(`{"message": "deleted sathya"}`))
default:
 w.WriteHeader(http.StatusNotFound)
 w.Write([]byte(`{"message": "not found"}`))
}
}

func main() {
http.HandleFunc("/", localhome)
log.Fatal(http.ListenAndServe(":4640", nil))
}
