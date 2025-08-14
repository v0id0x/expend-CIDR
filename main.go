package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <CIDR>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 185.14.168.0/22\n", os.Args[0])
		os.Exit(1)
	}
	cidr := os.Args[1]

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CIDR: %v\n", err)
		os.Exit(1)
	}

	currentIP := ip.To4()
	if currentIP == nil {
		fmt.Fprintf(os.Stderr, "Error: only IPv4 is supported by this tool.\n")
		os.Exit(1)
	}


	for ip := inc(currentIP); ipnet.Contains(ip); ip = inc(ip) {
		broadcast := make(net.IP, len(ipnet.IP))
		for i := range ipnet.IP {
			broadcast[i] = ipnet.IP[i] | ^ipnet.Mask[i]
		}
		
		if ip.Equal(broadcast) {
			break
		}

		fmt.Println(ip)
	}
}

func inc(ip net.IP) net.IP {
	newIP := make(net.IP, len(ip))
	copy(newIP, ip)

	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i]++
		if newIP[i] > 0 {
			break
		}
	}
	return newIP
}
