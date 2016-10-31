package main

import "net"
import "net/http"
import "flag"
import "log"

func withAvailableHost(bind string, http_serve func(free_bind string)) {
        available, _ := BindAvailable(bind) // discard err handle now
	if  available {
		http_serve(bind)
	} else if Confirm() {
		// call with free ports
	} else {
		log.Fatalln("Exiting not available : ", bind)
	}
}

func main() {

	host := flag.String("host", "0.0.0.0", "Server host to listen")
	port := flag.String("port", "5656", "Server port")
	dir := flag.String("dir", ".", "Directory location to files")
	flag.Parse()
	// i love anon :)
	greet := func(addr string) {
		log.Println("Listening at ", addr)
	}
	bind := net.JoinHostPort(*host, *port)
	if *host != "0.0.0.0" {
		greet(bind)
	} else {
		addr_list, err := GetAddrs()
		if err != nil {
			log.Println("error on getting system interface address")
		}
		for _, addr := range addr_list {
			greet(net.JoinHostPort(addr, *port))
		}
	}
	withAvailableHost(bind, func(bind string) {
		log.Fatalln(http.ListenAndServe(bind, http.FileServer(http.Dir(*dir))))
	})

}
