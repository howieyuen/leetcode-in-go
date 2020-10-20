package _468_validate_ip_address

import (
	"net"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	const (
		V4      = "IPv4"
		V6      = "IPv6"
		Neither = "Neither"
	)
	if strings.Contains(IP, ".") && strings.Contains(IP, ":") {
		return Neither
	} else if strings.Contains(IP, ".") {
		strs := strings.Split(IP, ".")
		if len(strs) != 4 {
			return Neither
		}
		for _, str := range strs {
			if len(str) == 0 || len(str) > 3 {
				return Neither
			}
			if str[0] == '0' && len(str) != 1 {
				return Neither
			}
			num, err := strconv.Atoi(str)
			if err != nil || num < 0 || num > 255 {
				return Neither
			}
		}
		return V4
	} else if strings.Contains(IP, ":") {
		strs := strings.Split(IP, ":")
		if len(strs) != 8 {
			return Neither
		}
		for _, str := range strs {
			if len(str) == 0 || len(str) > 4 {
				return Neither
			}
			for i := range str {
				if str[i] >= '0' && str[i] <= '9' || str[i] >= 'a' && str[i] <= 'f' || str[i] >= 'A' && str[i] <= 'F' {
					continue
				}
				return Neither
			}
		}
		return V6
	} else {
		return Neither
	}
}

func validIPAddress1(IP string) string {
	ip := net.ParseIP(IP)
	if ip != nil {
		if ip.To4() != nil {
			if ip.String() != IP {
				return "Neither"
			}
			return "IPv4"
		}
		temp := strings.Split(IP, ":")
		for i := range temp {
			if len(temp[i]) > 4 || len(temp[i]) == 0 {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
