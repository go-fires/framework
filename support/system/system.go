package system

import (
	"net"
	"os"
	"runtime"
)

// NumCPU Get the number of logical CPUs usable by the current process
func NumCPU() int {
	return runtime.NumCPU()
}

// Hostname Get the host name reported by the kernel
func Hostname() string {
	name, _ := os.Hostname()
	return name
}

// IsWindows Is windows system
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// LocalIPs Get all non-loopback IP addresses
func LocalIPs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	var ips []string
	for _, addr := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// skip ipv6
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return ips
}

// LocalIP Get the first non-loopback IP address
func LocalIP() string {
	ips := LocalIPs()
	if len(ips) > 0 {
		return ips[0]
	}

	return ""
}
