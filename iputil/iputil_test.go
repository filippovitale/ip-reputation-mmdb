package iputil

import (
	"fmt"
	"net"
	"testing"
)

func TestFromIPBeginAndEndToIPNet(t *testing.T) {
	if FromIPBeginAndEndToIPNet(net.ParseIP("1.0.66.0"), net.ParseIP("1.0.66.127")) == nil {
		t.Error("No cidr block found.")
	}

	if FromIPBeginAndEndToIPNet(net.ParseIP("1.0.70.0"), net.ParseIP("1.0.70.255")) == nil {
		t.Error("No cidr block found.")
	}

	if FromIPBeginAndEndToIPNet(net.ParseIP("254.50.53.0"), net.ParseIP("254.50.53.255")) == nil {
		t.Error("No cidr block found.")
	}
}

func TestParseIPOrCIDR(t *testing.T) {
	if !(fmt.Sprintf("%s", ParseIPOrCIDR("1.0.66.127")) == "1.0.66.127/32") {
		t.Error("Error creating the cidr block.")
	}

	if !(fmt.Sprintf("%s", ParseIPOrCIDR("1.0.66.127/25")) == "1.0.66.0/25") {
		t.Error("Error creating the cidr block.")
	}

	if !(fmt.Sprintf("%s", ParseIPOrCIDR("2001:db8::")) == "2001:db8::/128") {
		t.Error("Error creating the cidr block.")
	}

	if ParseIPOrCIDR("# random string #") != nil {
		t.Error("Error creating a cidr block from a random string.")
	}
}
