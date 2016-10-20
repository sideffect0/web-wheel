package main

import "net"
import "strings"

func GetAddrs() (addrs []string, err error) {
 addr_list, err := net.InterfaceAddrs()
 
 if err != nil{
    return nil, err
 }
 for _, addr := range(addr_list){
   split := strings.Split(addr.String(), "/")
   // removing ipv6 for now
   if net.ParseIP(split[0]).To4() != nil {
    addrs = append(addrs, split[0])
   }
 }
 return
}
