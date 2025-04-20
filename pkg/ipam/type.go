package ipam

type Type string

const (
	Dummy     Type = "dummy"
	DHCP      Type = "dhcp"
	HostLocal Type = "host-local"
	Static    Type = "static"
	Etcd      Type = "etcd"
)

func ValidType(t Type) bool {
	switch t {
	case Dummy, DHCP, HostLocal, Static, Etcd:
		return true
	default:
		return false
	}
}
