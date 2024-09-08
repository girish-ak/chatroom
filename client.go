package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commands<-command{
				id: CMD_NICK,
				client: c,
				args: args,
			} 
		case "/join":
			c.commands<-command{
				id: CMD_JOIN,
				client: c,
				args: args,
			}
		case "/rooms":
			c.commands<-command{
				id: CMD_ROOMS,
				client: c,
				args: args,
			}
		case "/msg":
			c.commands<-command{
				id: CMD_MSG,
				client: c,
				args: args,
			}
		case "/exit":
			c.commands<-command{
				id: CMD_EXIT,
				client: c,
				args: args,
			}
		case "/help":
			c.commands<-command{
				id: CMD_HELP,
				client: c,
				args: args,
			}
		default:
			c.err(fmt.Errorf("Command Unknown -> %s , please give /help to list available commands\n", cmd))
		}
	}
}

func (c *client) err (err error){
	c.conn.Write([] byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg (msg string) {
	c.conn.Write([] byte("\n\n--------------------------------------------\n >> " + msg + "\n--------------------------------------------\n\n"))
}
