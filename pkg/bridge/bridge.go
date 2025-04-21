package bridge

import (
	"errors"
	"fmt"
	current "github.com/containernetworking/cni/pkg/types/040"
	"github.com/containernetworking/plugins/pkg/utils/sysctl"
	"github.com/vishvananda/netlink"
	"syscall"
)

func SetupBridge(brName string, mtu int) (*netlink.Bridge, *current.Interface, error) {
	linkAttrs := netlink.NewLinkAttrs()
	linkAttrs.Name = brName
	linkAttrs.MTU = mtu

	br := &netlink.Bridge{
		LinkAttrs: linkAttrs,
	}

	// it equals to `ip link add name <brName> type bridge mtu <mtu>`
	err := netlink.LinkAdd(br)
	if err != nil && !errors.Is(err, syscall.EEXIST) {
		return nil, nil, fmt.Errorf("could not add :%q: %v", brName, err)
	}

	//br, err = bridgeByName(brName)
	// we want to own the routes for this interface
	_, _ = sysctl.Sysctl(fmt.Sprintf("net/ipv6/conf/%s/accept_ra", brName), "0")

	if err := netlink.LinkSetUp(br); err != nil {
		return nil, nil, err
	}

	return br, &current.Interface{
		Name: brName,
	}, nil
}

//func bridgeByName(name string) (*netlink.Bridge, error) {
//	l, err := netlinksafe.LinkByName(name)
//	if err != nil {
//		return nil, fmt.Errorf("could not lookup %q: %v", name, err)
//	}
//	br, ok := l.(*netlink.Bridge)
//	if !ok {
//		return nil, fmt.Errorf("%q already exists but is not a bridge", name)
//	}
//	return br, nil
//}
