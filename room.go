package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		//we dont want to broadcast the same message to the sender
		//so we check
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
