package utils

import (
	"encoding/binary"
	"errors"
	"net"
)

var InvalidIPV4 = errors.New("Invalid IPv4")

func IPV4ToInt(ip net.IP) (uint32, error) {
	if ip == nil {
		return 0, InvalidIPV4
	}

	ip = ip.To4()
	if ip == nil {
		return 0, InvalidIPV4
	}

	return binary.BigEndian.Uint32(ip), nil
}

func IntToIPV4(ip uint32) net.IP {
	data := make(net.IP, 4)
	binary.BigEndian.PutUint32(data, ip)
	return data
}

func IPStrToInt(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, InvalidIPV4
	}
	return IPV4ToInt(ip)
}

func IntToIPStr(ip uint32) string {
	return IntToIPV4(ip).String()
}

func GetLastIPV4Int(cidr string) (lastIP4 uint32, err error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return
	}

	firstIP4, err := IPV4ToInt(ip.Mask(ipNet.Mask))
	if err != nil {
		return
	}
	ipNetMaskInt := binary.BigEndian.Uint32(ipNet.Mask)
	lastIP4 = firstIP4 | ^ipNetMaskInt
	return
}
