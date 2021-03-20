package tcpsyn

import (
	"fmt"

	"github.com/v-byte-cpu/sx/pkg/scan"
)

// Set to typical maximum Ethernet frame size = MTU (1500 bytes)
// + Ethernet header (14 bytes) + FCS (4 bytes)
const MaxPacketLength = 1518

func BPFFilter(r *scan.Range) (filter string, maxPacketLength int) {
	if r.DstSubnet == nil {
		return "tcp", MaxPacketLength
	}
	return fmt.Sprintf("tcp and ip src net %s", r.DstSubnet.String()), MaxPacketLength
}
