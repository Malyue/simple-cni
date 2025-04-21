package veth

import (
	"fmt"
	current "github.com/containernetworking/cni/pkg/types/040"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

func SetupVeth(netNs ns.NetNS, br *netlink.Bridge, ifName string, mtu int, mac string) (*current.Interface, *current.Interface, error) {
	contIface := &current.Interface{}
	hostIface := &current.Interface{}

	// create veth pair in the container and move the other end into host netns
	err := netNs.Do(func(hostNs ns.NetNS) error {
		hostVeth, containerVeth, err := ip.SetupVeth(ifName, mtu, mac, hostNs)
		if err != nil {
			return err
		}
		contIface.Name = containerVeth.Name
		contIface.Mac = containerVeth.HardwareAddr.String()
		contIface.Sandbox = netNs.Path()
		hostIface.Name = hostVeth.Name
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	// get host veth
	// it euqals to `ip link show <hostIface.Name>`
	hostVeth, err := netlink.LinkByName(hostIface.Name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to lookup %q: %v", hostIface.Name, err)
	}

	// connect host veth end to the bridge
	// it euqals to `ip link set <hostVeth.Name> master <br.Name> `
	if err := netlink.LinkSetMaster(hostVeth, br); err != nil {
		return nil, nil, fmt.Errorf("failed to connect %q to bridge %v: %v", hostVeth.Attrs().Name, br.Attrs().Name, err)
	}

	return hostIface, contIface, nil

}
