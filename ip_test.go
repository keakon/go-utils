package utils

import (
	"bytes"
	"net"
	"testing"
)

var localhost = uint32(127<<24 + 1)

func TestIPV4ToInt(t *testing.T) {
	tests := []struct {
		ip   net.IP
		want *uint32
	}{
		{
			ip:   []byte{},
			want: nil,
		},
		{
			ip:   []byte{1, 2},
			want: nil,
		},
		{
			ip:   net.ParseIP("::1"),
			want: nil,
		},
		{
			ip:   net.ParseIP("0.0.0.0"),
			want: new(uint32),
		},
		{
			ip:   net.ParseIP("127.0.0.1"),
			want: &localhost,
		},
	}

	for _, test := range tests {
		t.Run(test.ip.String(), func(t *testing.T) {
			got := IPV4ToInt(test.ip)
			if test.want == nil {
				if got != nil {
					t.Errorf("got %v, want nil", got)
				}
			} else if *test.want != *got {
				t.Errorf("got %v, want %v", *got, *test.want)
			}
		})
	}
}

func TestIntToIPV4(t *testing.T) {
	tests := []struct {
		i    uint32
		want net.IP
	}{
		{
			i:    0,
			want: make([]byte, 4),
		},
		{
			i:    localhost,
			want: []byte{127, 0, 0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.want.String(), func(t *testing.T) {
			got := IntToIPV4(test.i)
			if !bytes.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestIPStrToInt(t *testing.T) {
	tests := []struct {
		ip   string
		want *uint32
	}{
		{
			ip:   "",
			want: nil,
		},
		{
			ip:   "1.2",
			want: nil,
		},
		{
			ip:   "::1",
			want: nil,
		},
		{
			ip:   "0.0.0.0",
			want: new(uint32),
		},
		{
			ip:   "127.0.0.1",
			want: &localhost,
		},
	}

	for _, test := range tests {
		t.Run(test.ip, func(t *testing.T) {
			got := IPStrToInt(test.ip)
			if test.want == nil {
				if got != nil {
					t.Errorf("got %v, want nil", got)
				}
			} else if *test.want != *got {
				t.Errorf("got %v, want %v", *got, *test.want)
			}
		})
	}
}

func TestIntToIPStr(t *testing.T) {
	tests := []struct {
		i    uint32
		want string
	}{
		{
			i:    0,
			want: "0.0.0.0",
		},
		{
			i:    localhost,
			want: "127.0.0.1",
		},
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			got := IntToIPStr(test.i)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
