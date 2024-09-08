package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listrooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_EXIT:
			s.exit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {

	log.Printf("Connection Established with %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string){
	c.nick = args[1]
	c.msg(fmt.Sprintf("Your Nickname has been set to %s",c.nick))
}

func (s *server) join(c *client, args []string){
	roomName := args[1]

	r,ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name: roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c
	s.quitCurrentRoom(c)
	c.room = r
	r.broadcast(c,fmt.Sprintf("%s has entered the room!",c.nick))
	c.msg(fmt.Sprintf("Welcome to %s!",r.name))
}

func (s *server) quitCurrentRoom (c *client){
	if c.room!=nil{
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c,fmt.Sprintf("%s has left the room",c.nick))
	}
}

func (s *server) listrooms (c *client, args []string){
	var rooms [] string
	for name:= range s.rooms{
		rooms = append(rooms,name)
	}
	c.msg(fmt.Sprintf("Here are the Available Rooms : \n %s",strings.Join(rooms,"\n")))
}

func (s *server) msg (c *client, args []string){
	 if c.room==nil{
		c.err(errors.New("You have to be participent of a room before sending a mesasge!,please give /rooms to check available rooms"))
		return
	}
	
	c.room.broadcast(c, c.nick+": "+strings.Join(args[1:len(args)]," "))
}

func (s *server) exit (c *client, args[] string){
	log.Printf("%s has been Disconnected!",c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg("Until next time! Sayonara...")
	c.conn.Close() 
}
















