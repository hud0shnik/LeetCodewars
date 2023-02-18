package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	ipArray := strings.Split(ip, ".")
	if len(ipArray) != 4 {
		return false
	}
	for _, ipPart := range ipArray {
		num, _ := strconv.Atoi(ipPart)
		if num < 0 || num > 255 {
			return false
		}
		if ipPart != fmt.Sprint(num) {
			return false
		}
	}
	return true
}
