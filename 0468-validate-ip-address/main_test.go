package _468_validate_ip_address

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{IP: "172.16.254.1"}, want: "IPv4"},
		{args: args{IP: "2001:0db8:85a3:0:0:8A2E:0370:7334"}, want: "IPv6"},
		{args: args{IP: "256.256.256.256"}, want: "Neither"},
		{args: args{IP: "2001:0db8:85a3:0:0:8A2E:0370:7334:"}, want: "Neither"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
			if got := validIPAddress1(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress1() = %v, want %v", got, tt.want)
			}
		})
	}
}
