package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Millisecond)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()

			}
		}
	}
	return ""
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	var port int
	flag.IntVar(&port, "p", 25565, "number of the port")
	flag.Parse()

	ip, ipnet, err := net.ParseCIDR(GetLocalIP() + "/24")
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		open := scanPort("tcp", ip.String(), port)
		if open == true {
			fmt.Printf("IP %s Port %v Open: %t\n ", ip, port, open)
		}
	}

}
