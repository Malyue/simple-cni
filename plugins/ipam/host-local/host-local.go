package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	current "github.com/containernetworking/cni/pkg/types/040"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	"simple-cni/plugins/ipam/host-local/backend/allocator"
	"simple-cni/plugins/ipam/host-local/backend/store"
)

func cmdAdd(args *skel.CmdArgs) error {
	ipamConf, confVersion, err := allocator.LoadIPAMConfig(args.StdinData, args.Args)
	if err != nil {
		return err
	}

	result := &current.Result{CNIVersion: current.ImplementedSpecVersion}

	s, err := store.New(ipamConf.Name, ipamConf.DataDir)

	allocator := allocator.NewIPAllocator(s, &ipamConf.Ranges)
	ipResp, err := allocator.Get(args.ContainerID, args.IfName, nil)
	if err != nil {
		return err
	}

}

func cmdCheck(args *skel.CmdArgs) error {
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	return nil
}

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.All, bv.BuildString("host-local"))
}
