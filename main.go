package main

import (
	"os"
	"fmt"
	"bufio"
	"net"
	"strings"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func cidrToHosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips[1 : len(ips)-1], nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		cidr := strings.ToLower(sc.Text())	
		hosts, _ := cidrToHosts(cidr)
		for _, host := range hosts {
			fmt.Println(host)
		}
	}
}
