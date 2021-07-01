package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
     "github.com/gorilla/mux"
)

func main() {
   // r := router.Router()
var r *mux.Router
    r = mux.NewRouter()
    //router(r)
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)

    port := os.Getenv("PORT")
    if port == "" {
            port = "3000"
    }

    fmt.Println("Starting server on the ******** GO UNITED!!!! ******* port..."+port)

    log.Fatal(http.ListenAndServe(":"+port, r))
}