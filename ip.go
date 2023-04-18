package utils

import (
	"encoding/binary"
	"net"
)

func IPV4ToInt(ip net.IP) *uint32 {
	if ip == nil {
		return nil
	}

	ip = ip.To4()
	if ip == nil {
		return nil
	}

	i := binary.BigEndian.Uint32(ip)
	return &i
}

func IntToIPV4(ip uint32) net.IP {
	data := make(net.IP, 4)
	binary.BigEndian.PutUint32(data, ip)
	return data
}

func IPStrToInt(ipStr string) *uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil
	}
	return IPV4ToInt(ip)
}

func IntToIPStr(ip uint32) string {
	return IntToIPV4(ip).String()
}
