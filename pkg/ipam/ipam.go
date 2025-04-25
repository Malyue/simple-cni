package ipam

import (
	"context"
	"fmt"
	"time"

	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/vishvananda/netlink"
	"k8s.io/klog/v2"
)

func ExecAdd(plugin string, netconf []byte) (types.Result, error) {
	return invoke.DelegateAdd(context.TODO(), plugin, netconf, nil)
}

func ExecCheck(plugin string, netconf []byte) error {
	return invoke.DelegateCheck(context.TODO(), plugin, netconf, nil)
}

func ExecDel(plugin string, netconf []byte) error {
	return invoke.DelegateDel(context.TODO(), plugin, netconf, nil)
}

func ExecStatus(plugin string, netconf []byte) error {
	return invoke.DelegateStatus(context.TODO(), plugin, netconf, nil)
}

const (
	// Note: use slash as separator so we can have dots in interface name (VLANs)
	DisableIPv6SysctlTemplate    = "net/ipv6/conf/%s/disable_ipv6"
	KeepAddrOnDownSysctlTemplate = "net/ipv6/conf/%s/keep_addr_on_down"

	dadSettleTimeout = 5 * time.Second
)

// ConfigureIface takes the result of IPAM plugin and
// applies to the ifName interface
func ConfigureIface(ifName string, res *current.Result) error {
	if len(res.Interfaces) == 0 {
		return fmt.Errorf("no interfaces to configure")
	}

	return nil
	// get the link
	//link,err := netlink.LinkByName(ifName)
	//if err != nil {
	//	return fmt.Errorf("failed to lookup %q: %v", ifName, err)
	//}
	//
	//for _,ipConfig := range res.IPs {
	//	if ipConfig.Interface == nil || *ipConfig.Interface !=
	//}

}
