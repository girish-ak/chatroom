package chatroom

type CommandID int

const (
	CMD_NICK CommandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_EXIT
)

type command struct {
	id CommandID 
	client *client
}
