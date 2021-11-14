package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if isIPv4Adress(IP) {
		return "IPv4"
	}
	if isIPv6Adress(IP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv6Adress(IP string) bool {
	ipArray := strings.Split(IP, ":")
	if len(ipArray) != 8 {
		return false
	}
	for _, ipPart := range ipArray {
		ipPartLen := len(ipPart)
		if ipPartLen == 0 || ipPartLen > 4 {
			return false
		}
		for i := 0; i < ipPartLen; i++ {
			if !(ipPart[i] >= '0' && ipPart[i] <= '9') && !(ipPart[i] >= 'A' && ipPart[i] <= 'F') && !(ipPart[i] >= 'a' && ipPart[i] <= 'f') {
				return false
			}
		}
	}
	return true
}

func isIPv4Adress(IP string) bool {
	ipArray := strings.Split(IP, ".")
	if len(ipArray) != 4 {
		return false
	}
	for _, ipPart := range ipArray {
		num, err := strconv.Atoi(ipPart)
		if err != nil || num < 0 || num > 255 {
			return false
		}
		if ipPart != fmt.Sprint(num) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("256.256.256.256"))
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334:"))
	fmt.Println(validIPAddress("1e1.4.5.6"))
}
