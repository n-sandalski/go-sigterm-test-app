package main

import (
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on port 8080...")

    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, syscall.SIGTERM)
        <-c
        fmt.Println("Caught SIGTERM")
        os.Exit(0)
    }()

    http.ListenAndServe(":8080", nil)
}
