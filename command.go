package main

type commandID int

// all possible current commands
const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client  // sender of the current command
	args   []string // args for the command
}
