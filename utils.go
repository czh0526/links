package main

import "github.com/libp2p/go-libp2p/core/host"

func shortID(h host.Host) string {
	id := h.ID().String()
	if len(id) > 10 {
		return id[:10]
	}
	return id
}
