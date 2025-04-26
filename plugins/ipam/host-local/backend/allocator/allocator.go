package allocator

import (
	current "github.com/containernetworking/cni/pkg/types/100"
	"net"
	"simple-cni/plugins/ipam/host-local/backend/store"
)

type IPAllocator struct {
	store   *store.Store
	RangeId string
	Ranges  *[]Range
}

func NewIPAllocator(store *store.Store, r *[]Range) *IPAllocator {
	return &IPAllocator{
		store:  store,
		Ranges: r,
	}
}

func (a *IPAllocator) Get(id string, ifname string, requestedIP net.IP) (*current.IPConfig, error) {
	a.store.Lock()
	defer a.store.Unlock()

	//if requestedIP != nil {
	//
	//}
	// check whether the id has ip
	a.store.GetByID(id, ifname)
	for _, r := range *a.Ranges {
		// check if the range is full
		// if it is full, use the next range
	}

	return nil, nil
}

func (a *IPAllocator) GetIter() (*RangeIter, error) {
	return nil, nil
}
