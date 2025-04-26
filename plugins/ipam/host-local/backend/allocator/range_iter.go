package allocator

import "net"

type RangeIter struct {
	rangeset *[]Range

	// The current range id
	rangeIdx int

	// Our current position
	cur net.IP

	// The IP where we started iterating; if we hit this again, we're done.
	startIP net.IP
}
