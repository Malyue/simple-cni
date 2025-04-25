module simple-cni

go 1.24

require (
	github.com/alexflint/go-filemutex v1.3.0
	github.com/containernetworking/cni v1.3.0
	github.com/containernetworking/plugins v1.6.2
	github.com/stretchr/testify v1.4.0
	github.com/vishvananda/netlink v1.3.0
	github.com/vishvananda/netns v0.0.4
	k8s.io/klog/v2 v2.130.1
)

require (
	github.com/coreos/go-iptables v0.8.0 // indirect
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/safchain/ethtool v0.5.9 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	sigs.k8s.io/knftables v0.0.18 // indirect
)
