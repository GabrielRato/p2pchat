package model


type Command struct {
	Option int
	Display string
}

var CONNECT_PEER = Command{0, "connect to a peer"}
var SAVE_CONTACT = Command{1, "save contact to list"}
var EXIT = Command{2, "exit"}
