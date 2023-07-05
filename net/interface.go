//go:build !android

package manet

import (
	"net"

	ma "github.com/multiformats/go-multiaddr"
)

// InterfaceMultiaddrs will return the addresses matching net.InterfaceAddrs
func InterfaceMultiaddrs() ([]ma.Multiaddr, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
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
