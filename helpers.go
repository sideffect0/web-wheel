package main

import "net"
import "strings"
import "log"

func GetAddrs() (addrs []string, err error) {
	addr_list, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}
	for _, addr := range addr_list {
		split := strings.Split(addr.String(), "/")
		// removing ipv6 for now
		if net.ParseIP(split[0]).To4() != nil {
			addrs = append(addrs, split[0])
		}
	}
	return
}

func GetFreeBind() (free_bind string) {

	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		log.Fatalln(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	return l.Addr()
}

func BindAvailable(bind string) (status bool, err error) {

	tester, err := net.Listen("tcp", bind)
	// port not available with given bind
	if err != nil {
		return false, err
	}
	// free up for dev-server
	tester.Close()
	return true, nil
}

func Confirm() (status bool) {
	return
}
