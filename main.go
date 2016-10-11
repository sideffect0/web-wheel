package main

import "net"
import "net/http"
import "flag"
import "fmt"

func main() {
    
    host := flag.String("host", "0.0.0.0", "Server host to listen")
    port := flag.String("port", "5656", "Server port")
    dir  := flag.String("dir", ".", "Directory location to files")
    flag.Parse()
    bind := net.JoinHostPort(*host, *port)
    fmt.Println("Listening at ", bind)
    panic(http.ListenAndServe(bind, http.FileServer(http.Dir(*dir))))

}
