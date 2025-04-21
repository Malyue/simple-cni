package main

import (
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	"github.com/containernetworking/plugins/pkg/utils/sysctl"
	"simple-cni/pkg/bridge"
	"simple-cni/pkg/ipam"
	"simple-cni/pkg/veth"
)

func cmdAdd(args *skel.CmdArgs) error {
	success := false
	conf, err := loadNetConf(args.StdinData, args.Args)
	if err != nil {
		return fmt.Errorf("failed to load netconf : %v", err)
	}

	// get ns
	netns, err := ns.GetNS(args.Netns)
	if err != nil {
		return fmt.Errorf("failed to get ns %q : %v ", args.Netns, err)
	}
	defer netns.Close()

	// create bridge if necessary
	br, brInterface, err := bridge.SetupBridge(conf.Bridge, conf.MTU)
	if err != nil {
		return fmt.Errorf("failed to setup bridge : %v", err)
	}

	hostVeth, contVeth, err := veth.SetupVeth(netns, br, args.IfName, conf.MTU, conf.Mac)
	if err != nil {
		return fmt.Errorf("failed to setup veth : %v", err)
	}

	return nil

	// allocate ip from ipam
	// valid ipam type
	valid := ipam.ValidType(ipam.Type(conf.IPAM.Type))
	if !valid {
		return fmt.Errorf("The type of ipam is invalid")
	}

	r, err := ipam.ExecAdd(conf.IPAM.Type, args.StdinData)
	if err != nil {
		return fmt.Errorf("failed to allocate ip from ipam : %v", err)
	}

	// release IP in case of failure
	defer func() {
		if !success {
			ipam.ExecDel(conf.IPAM.Type, args.StdinData)
		}
	}()

	ipamResult, err := current.NewResultFromResult(r)
	if err != nil {
		return err
	}

	result := &current.Result{
		CNIVersion: current.ImplementedSpecVersion,
		Interfaces: []*current.Interface{
			brInterface,
			hostVeth,
			contVeth,
		},
		IPs:    ipamResult.IPs,
		Routes: ipamResult.Routes,
		DNS:    ipamResult.DNS,
	}

	// Configure the container hardware address and IP address(es)
	if err := netns.Do(func(_ ns.NetNS) error {
		_, _ = sysctl.Sysctl(fmt.Sprintf("net/ipv4/conf/%s/arp_notify", args.IfName), "1")

		// Add the IP to the interface
		return ipam.ConfigureIface(args.IfName, result)
	}); err != nil {
		return err
	}

	return nil

}

func cmdDel(args *skel.CmdArgs) error {
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	return nil
}

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		//Status: cmdStatus,
		/* FIXME GC */
	}, version.All, bv.BuildString("bridge"))
}
