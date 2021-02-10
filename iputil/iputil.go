package iputil

import (
	"fmt"
	"log"
	"net"
)

func ParseIPOrCIDR(ipOrCidr string) *net.IPNet {
	// First try parsing it as an IP
	ip := net.ParseIP(ipOrCidr)
	if ip != nil && ip.To4() != nil {
		// IPv4
		_, cidr, _ := net.ParseCIDR(ipOrCidr + "/32")
		return cidr
	}
	if ip != nil && ip.To4() == nil {
		// IPv6
		_, cidr, _ := net.ParseCIDR(ipOrCidr + "/128")
		return cidr
	}

	// CIDR
	_, cidr, _ := net.ParseCIDR(ipOrCidr)
	return cidr
}

func FromIPBeginAndEndToIPNet(begin net.IP, end net.IP) *net.IPNet {
	if begin.Equal(net.ParseIP("0.0.0.0")) || begin.Equal(net.ParseIP("127.0.0.0")) {
		return nil
	}

	var mask = 32
	for mask > 0 {
		cidr := begin.String() + fmt.Sprintf("/%d", mask)
		_, ipnet, _ := net.ParseCIDR(cidr)

		if ipnet.Contains(end) {
			break
		}
		mask--
	}

	if mask == 0 {
		log.Printf("Range %s-%s unparsable", begin, end)
		return nil
	}

	cidr := begin.String() + fmt.Sprintf("/%d", mask)
	_, ipnet, _ := net.ParseCIDR(cidr)
	return ipnet
}
