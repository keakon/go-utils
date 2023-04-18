package utils

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

const defaultHTTPSPort uint16 = 443

var domainPattern = regexp.MustCompile(`^([a-zA-Z0-9]([-a-zA-Z0-9]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z0-9]([-a-zA-Z0-9]{0,61}[a-zA-Z0-9])?$`)

func IsDigit(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func IsDomain(domain string) bool {
	length := len(domain)
	if length >= 3 && length <= 253 && domainPattern.MatchString(domain) {
		sepIndex := strings.LastIndexByte(domain, '.')
		if sepIndex > 0 {
			label := domain[sepIndex+1:]
			if label != "" {
				return !IsDigit(label) // 顶级域名不能全是数字，例如 abc.12345
			}
		} else if sepIndex == -1 { // 不包含 . 的域名只能是顶级域名或内部域名，例如 localhost
			return !IsDigit(domain) // 顶级域名不能全是数字
		}
	}
	return false
}

func ParseAuthority(authority string) (host string, port uint16) {
	// 支持下列格式：
	// test
	// test:443
	// test.com
	// test.com:443
	// 127.0.0.1
	// 127.0.0.1:443
	// [::1]
	// [::1]:443

	if authority != "" {
		if authority[0] == '[' { // IPv6
			if strings.Contains(authority, "::") {
				lastIndex := len(authority) - 1
				if authority[lastIndex] == ']' {
					host := authority[1:lastIndex]
					if net.ParseIP(host) != nil { // 前面已经判断了包含 ::，因此只要能解析出来，一定是 IPv6
						return host, defaultHTTPSPort
					}
				} else {
					sepIndex := strings.LastIndexByte(authority, ':') // 因为前面已经判断包含 :: 了，下面就不检查 sepIndex 是否为负数了
					if authority[sepIndex-1] == ']' {
						p := authority[sepIndex+1:]
						pt, err := strconv.ParseUint(p, 10, 0)
						if err == nil && pt > 0 && pt <= 65535 {
							return authority[1 : sepIndex-1], uint16(pt)
						}
					}
				}
			}
		} else { // IPv4 或域名
			parts := strings.SplitN(authority, ":", 2)
			if len(parts) == 1 {
				host = parts[0]
				port = defaultHTTPSPort
			} else {
				p := parts[1]
				pt, err := strconv.ParseUint(p, 10, 0)
				if err != nil || pt == 0 || pt > 65535 {
					return
				}
				host = parts[0]
				port = uint16(pt)
			}
			if net.ParseIP(host) != nil || IsDomain(host) {
				return
			}
			return "", 0
		}
	}
	return
}
