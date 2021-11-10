package Player

import (
	"ObjectOrientedCourse/src/Items"
	"ObjectOrientedCourse/src/Commands"
	"ObjectOrientedCourse/src/Rooms"
	"fmt"
)

var health int
var maxHealth int
var room *Rooms.Room
var alive bool
var items []Items.Item

func Initialize() {
	room = Rooms.DefaultRoom()
	health = 72
	maxHealth = 100
	alive = true
}

func RunCommand(command string, hasMoved *bool) {
	if command == "talk" {
		Commands.CommandTalk(room)
	} else if command == "fight" {
		if !Commands.CommandFight(room, &health, maxHealth) {
			alive = false
		}
	} else if command == "move" {
		room = Commands.CommandMove(room)
		*hasMoved = true
	} else {
		fmt.Println("Invalid command!")
	}
}

func GetAvailableCommandString() string {
	availableCommands := room.GetAvailableCommands()
	availableCommandString := ""
	for i, command := range availableCommands {
		availableCommandString += command
		if i+1 != len(availableCommands) {
			availableCommandString += ", "
		}
	}
	return availableCommandString
}

func Alive() bool {
	return alive
}

func DescribeRoom() {
	room.Describe()
}

func AddItem(item Items.Item) {
	items = append(items, item)
}