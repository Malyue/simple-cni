package allocator

import (
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/plugins/pkg/ip"
	"net"
)

type RangeSet []Range

type Range struct {
	RangeStart net.IP      `json:"rangeStart,omitempty"` // The first ip, inclusive
	RangeEnd   net.IP      `json:"rangeEnd,omitempty"`   // The last ip, inclusive
	Subnet     types.IPNet `json:"subnet"`
	Gateway    net.IP      `json:"gateway,omitempty"`
}

// Contains Check the addr is in the range
func (r *Range) Contains(addr net.IP) bool {
	if len(addr) != len(r.Subnet.IP) {
		return false
	}
	subnet := (net.IPNet)(r.Subnet)
	if !subnet.Contains(addr) {
		return false
	}

	if r.RangeStart != nil {
		if ip.Cmp(addr, r.RangeStart) < 0 {
			return false
		}
	}

	if r.RangeEnd != nil {
		if ip.Cmp(addr, r.RangeEnd) > 0 {
			return false
		}
	}

	return true
}
