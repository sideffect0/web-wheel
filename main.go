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
    // i love anon :)
    greet:= func(addr string) {
     fmt.Println("Listening at ", addr)
    }
    bind := net.JoinHostPort(*host, *port)
    if *host != "0.0.0.0" {
       greet(bind)
    }else{
      addr_list, err := GetAddrs()
      if err != nil{
        fmt.Println("error on getting system interface address")
      }
      for _, addr := range(addr_list) {
         greet(net.JoinHostPort(addr, *port))
      }
    }
    panic(http.ListenAndServe(bind, http.FileServer(http.Dir(*dir))))

}
