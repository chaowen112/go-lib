package link

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

func GetInterfaceIpAddr(interfaceName string) (net.IP, error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)

	ief, err := net.InterfaceByName(interfaceName)

	if err != nil { // get interface
		return nil, errors.Wrap(err, "not found")
	}

	addrs, err = ief.Addrs()

	if err != nil {
		return nil, err
	}

	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return nil, errors.New(fmt.Sprintf("interface %s doesn't have an ipv4 address\n", interfaceName))
	}

	return ipv4Addr, nil
}
