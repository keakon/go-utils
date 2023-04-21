package utils

import (
	"bytes"
	"net"
	"testing"
)

const localhost = 127<<24 + 1

func TestIPV4ToInt(t *testing.T) {
	tests := []struct {
		ip      net.IP
		want    uint32
		wantErr bool
	}{
		{
			ip:      []byte{},
			want:    0,
			wantErr: true,
		},
		{
			ip:      []byte{1, 2},
			want:    0,
			wantErr: true,
		},
		{
			ip:      net.ParseIP("::1"),
			want:    0,
			wantErr: true,
		},
		{
			ip:      net.ParseIP("0.0.0.0"),
			want:    0,
			wantErr: false,
		},
		{
			ip:      net.ParseIP("127.0.0.1"),
			want:    localhost,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.ip.String(), func(t *testing.T) {
			got, err := IPV4ToInt(test.ip)
			if (err != nil) != test.wantErr {
				t.Errorf("got err %v, want %v", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
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
		ip      string
		want    uint32
		wantErr bool
	}{
		{
			ip:      "",
			want:    0,
			wantErr: true,
		},
		{
			ip:      "1.2",
			want:    0,
			wantErr: true,
		},
		{
			ip:      "::1",
			want:    0,
			wantErr: true,
		},
		{
			ip:      "0.0.0.0",
			want:    0,
			wantErr: false,
		},
		{
			ip:      "127.0.0.1",
			want:    localhost,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.ip, func(t *testing.T) {
			got, err := IPStrToInt(test.ip)
			if (err != nil) != test.wantErr {
				t.Errorf("got err %v, want %v", err, test.wantErr)
			}
			if test.want != got {
				t.Errorf("got %d, want %d", got, test.want)
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

func TestGetLastIPV4Int(t *testing.T) {
	tests := []struct {
		cidr    string
		want    uint32
		wantErr bool
	}{
		{
			want:    0,
			wantErr: true,
		},
		{
			cidr:    "127.0.0.1/8",
			want:    128<<24 - 1,
			wantErr: false,
		},
		{
			cidr:    "::1/8",
			want:    0,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.cidr, func(t *testing.T) {
			got, err := GetLastIPV4Int(test.cidr)
			if (err != nil) != test.wantErr {
				t.Errorf("got err %v, want %v", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
