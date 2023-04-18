package utils

import "testing"

func TestIsDomain(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{
			arg:  "",
			want: false,
		},
		{
			arg:  ".",
			want: false,
		},
		{
			arg:  "-",
			want: false,
		},
		{
			arg:  ".a",
			want: false,
		},
		{
			arg:  "a..com",
			want: false,
		},
		{
			arg:  "123456",
			want: false,
		},
		{
			arg:  "123.456",
			want: false,
		},
		{
			arg:  "123_456.com",
			want: false,
		},
		{
			arg:  "-123456.com",
			want: false,
		},
		{
			arg:  "123456-.com",
			want: false,
		},
		{
			arg:  "123456.com",
			want: true,
		},
		{
			arg:  "123.456.com",
			want: true,
		},
		{
			arg:  "abc.123.com",
			want: true,
		},
		{
			arg:  "123-456.com",
			want: true,
		},
		{
			arg:  "baidu.com",
			want: true,
		},
		{
			arg:  "www.baidu.com",
			want: true,
		},
		{
			arg:  "123.baidu.com",
			want: true,
		},
		{
			arg:  "localhost",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := IsDomain(tt.arg); got != tt.want {
				t.Errorf("IsDomain(%s) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestParseAuthority(t *testing.T) {
	tests := []struct {
		arg      string
		wantHost string
		wantPort uint16
	}{
		{
			arg:      "",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "localhost",
			wantHost: "localhost",
			wantPort: defaultHTTPSPort,
		},
		{
			arg:      "localhost:80",
			wantHost: "localhost",
			wantPort: 80,
		},
		{
			arg:      "10.130.0.1",
			wantHost: "10.130.0.1",
			wantPort: defaultHTTPSPort,
		},
		{
			arg:      "10.130.0.1:8443",
			wantHost: "10.130.0.1",
			wantPort: 8443,
		},
		{
			arg:      "10.130.0.1:88443",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "[::1]",
			wantHost: "::1",
			wantPort: 443,
		},
		{
			arg:      "[::1]:8443",
			wantHost: "::1",
			wantPort: 8443,
		},
		{
			arg:      "[::1]:88443",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "[::1]:88:443",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "[::1:443",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "::1:443",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "::1]:443",
			wantHost: "",
			wantPort: 0,
		},
		{
			arg:      "::1",
			wantHost: "",
			wantPort: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if gotHost, gotPort := ParseAuthority(tt.arg); gotHost != tt.wantHost || gotPort != tt.wantPort {
				t.Errorf("ParseAuthority(%s) = %s %d, want %s %d", tt.arg, gotHost, gotPort, tt.wantHost, tt.wantPort)
			}
		})
	}
}
