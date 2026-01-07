package discovery

import (
	"github.com/shirou/gopsutil/v3/net"
)

func PortsByPID(pid int32) []int {
	ports := []int{}

	conns, err := net.Connections("inet")
	if err != nil {
		return ports
	}

	for _, c := range conns {
		if c.Pid == pid && c.Laddr.Port != 0 {
			ports = append(ports, int(c.Laddr.Port))
		}
	}
	return ports
}

