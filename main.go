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
        fmt.Println("Will Catch SIGTERM")
        c := make(chan os.Signal, 1)
        signal.Notify(c, syscall.SIGTERM)
        <-c
        t := time.NewTicker(time.Second)
        for {
            fmt.Println("Caught SIGTERM")
            <-t.C
        }
    }()

    http.ListenAndServe(":8080", nil)
}
