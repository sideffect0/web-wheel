package main

import "net/http"
import "flag"
import "fmt"

func main() {
    
    host := flag.String("host", "localhost", "Server host to listen")
    port := flag.String("port", "5656", "Server port")
    dir  := flag.String("dir", ".", "Directory location to files")
    flag.Parse()
    bind := fmt.Sprintf("%s:%s", *host, *port)
    fmt.Println("Listening at ", bind)
    panic(http.ListenAndServe(bind, http.FileServer(http.Dir(*dir))))

}