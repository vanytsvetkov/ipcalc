package functions

import (
	"fmt"
	"net"
	"strings"
)

func IPsplit(prefix string, subnetSize int, sep string) (string, error) {
	var result string

	ip, ipNet, _ := net.ParseCIDR(prefix)

	var subnets []string

	for ipNet.Contains(ip) {
		subnetCIDR := &net.IPNet{IP: ip, Mask: net.CIDRMask(subnetSize, 32)}
		if subnetCIDR.IP.Equal(subnetCIDR.IP.Mask(subnetCIDR.Mask)) {
			subnets = append(subnets, subnetCIDR.String())
		}
		ip = incrementIP(ip)
	}

	result += fmt.Sprintf("Subnets /%d (%d pcs):\n\n", subnetSize, len(subnets))
	result += strings.Join(subnets, sep)

	return result, nil
}

func incrementIP(ip net.IP) net.IP {
	nextIP := make(net.IP, len(ip))
	copy(nextIP, ip)

	for j := len(nextIP) - 1; j >= 0; j-- {
		nextIP[j]++
		if nextIP[j] > 0 {
			break
		}
	}

	return nextIP
}
