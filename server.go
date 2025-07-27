package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room //here string will be the name of the room
	commands chan command     //where all the messeages from the client will be sent to the server

}

// small helper function to initiliase new server
func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newServer() *server {
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
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)

		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client is connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("got it, so you are %s from now", c.nick))
}
func (s *server) join(c *client, args []string) {
	//checking if room exists
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		//creating a new room if that room not exists
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		//adding the newly created room to the server
		s.rooms[roomName] = r
	}
	//adding the client to the list of members
	r.members[c.conn.RemoteAddr()] = c

	//checking if the client was in any old room then removing it from there
	s.quitCurrentRoom(c)
	//adding the room name to the client
	c.room = r

	//broadcasting to other users that new guy has come
	r.broadcast(c, fmt.Sprintf("%s has joined the rooms, lets welcome him", c.nick))
	c.msg(fmt.Sprintf("welcome to %s", r.name)) //sending the a welocme message to the user who just joined

}
func (s *server) listRooms(c *client, args []string) {

}
func (s *server) msg(c *client, args []string) {

}
func (s *server) quit(c *client, args []string) {

}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		//broadcasting other members of the room
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}
