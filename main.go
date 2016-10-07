package main

import "net/http"
import "flag"

func main() {
    
    bind := flag.String("b", ":5656", "Server host:port to listen")
    panic(http.ListenAndServe(*bind, http.FileServer(http.Dir("."))))

}
