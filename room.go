package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast (sender *client, msg string){
	for addr,i:= range r.members{
		if addr!=sender.conn.RemoteAddr(){
			i.msg(msg)
		}
	}
}
