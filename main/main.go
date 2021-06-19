package main

import (
    "os"
    "fmt"
    "go-united/router"
    "log"
    "net/http"
)

func main() {
    r := router.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)

    port := os.Getenv("PORT")
    if port == "" {
            port = "3000"
    }

    fmt.Println("Starting server on the ******** GO UNITED!!!! ******* port..."+port)

    log.Fatal(http.ListenAndServe(":"+port, r))
}