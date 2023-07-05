// Package manet provides Multiaddr
// (https://github.com/multiformats/go-multiaddr) specific versions of common
// functions in Go's standard `net` package. This means wrappers of standard
// net symbols like `net.Dial` and `net.Listen`, as well as conversion to
// and from `net.Addr`.
package manet

import (
	"net"

	"github.com/functionland/anet"
	ma "github.com/multiformats/go-multiaddr"
)

// InterfaceMultiaddrs will return the addresses matching net.InterfaceAddrs
func InterfaceMultiaddrs() ([]ma.Multiaddr, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		addrs, err = anet.InterfaceAddrs()
		if err != nil {
			return nil, err
		}
	}

	maddrs := make([]ma.Multiaddr, len(addrs))
	for i, a := range addrs {
		maddrs[i], err = FromNetAddr(a)
		if err != nil {
			return nil, err
		}
	}
	return maddrs, nil
}
