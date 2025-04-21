package ipam

import (
	"errors"
	"github.com/containernetworking/cni/pkg/types"
	ipam2 "simple-cni/pkg/ipam"
)

type Config interface {
	Parse(b []byte) interface{}
}

type IPAM interface {
	ExecAdd(ipamType string, ipamConfig Config) (*types.Result, error)
	ExecDel(ipamType string, ipamConfig Config) error
}

func Get(t ipam2.Type) (*IPAM, error) {
	valid := ipam2.ValidType(t)
	if !valid {
		return nil, errors.New("The type is invalid")
	}

	return nil, nil
	//switch t {
	//case ipam2.HostLocal:
	//	return nil, nil
	//}
}
