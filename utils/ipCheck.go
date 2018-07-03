package utils

import (
	"encoding/binary"
	"net"

	"strings"
)

func IpCheck(thisip string, ips []string) bool {
	for _, ip := range ips {
		ip = strings.TrimRight(ip, "/")
		if strings.Contains(ip, "/") {
			if ipCheckMask(thisip, ip) {
				return true
			}
		} else if strings.Contains(ip, "-") {
			ipRange := strings.SplitN(ip, "-", 2)
			if ipCheckRange(thisip, ipRange[0], ipRange[1]) {
				return true
			}
		} else {
			if thisip == ip {
				return true
			}
		}
	}
	return false
}

func ipCheckRange(ip, ipStart, ipEnd string) bool {
	thisIP := net.ParseIP(ip)
	firstIP := net.ParseIP(ipStart)
	endIP := net.ParseIP(ipEnd)
	if thisIP.To4() == nil || firstIP.To4() == nil || endIP.To4() == nil {
		return false
	}
	firstIPNum := ipToInt(firstIP.To4())
	endIPNum := ipToInt(endIP.To4())
	thisIpNum := ipToInt(thisIP.To4())
	if thisIpNum >= firstIPNum && thisIpNum <= endIPNum {
		return true
	}
	return false
}

func ipCheckMask(ip, ipMask string) bool {
	_, subnet, _ := net.ParseCIDR(ipMask)

	thisIP := net.ParseIP(ip)
	if subnet.Contains(thisIP) {
		return true
	}
	return false
}

func ipToInt(ip net.IP) int32 {
	return int32(binary.BigEndian.Uint32(ip.To4()))
}
