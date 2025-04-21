package allocator

import (
	current "github.com/containernetworking/cni/pkg/types/100"
	"net"
	"simple-cni/plugins/ipam/backend"
)

type IPAllocator struct {
	store backend.Store
}

func NewIPAllocator(store backend.Store) *IPAllocator {
	return &IPAllocator{
		store: store,
	}
}

func (a *IPAllocator) Get(id string, ifname string, requestedIP net.IP) (*current.IPConfig, error) {
	a.store.Lock()
	defer a.store.Unlock()

	//if requestedIP != nil {
	//	// check if the ip is exists
	//
	//	a.store.Reserve(id, ifname, requestedIP, a.ra)
	//}

	return nil, nil
}
