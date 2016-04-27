package main

import (
	"fmt"

	"github.com/Norskan/mud/gameworld"
	"github.com/Norskan/mud/gameworld/character"
	"github.com/Norskan/mud/gameworld/room"
)

func main() {

	a := [][]string{}
	a = append(a, []string{"w", "w", "w", "w"})
	a = append(a, []string{"w", ".", ".", "w"})
	a = append(a, []string{"w", ".", ".", "w"})
	a = append(a, []string{"w", ".", ".", "w"})
	a = append(a, []string{"w", "w", "w", "w"})

	roomTemplate := room.NewRoomTemplate(a)
	roomTemplate.AddTileTemplate(room.NewTileTemplate(".", "Floor", room.Moveable))
	roomTemplate.AddTileTemplate(room.NewTileTemplate("w", "Wall", room.NotMoveable))

	gw := gameworld.NewGameWorld(roomTemplate)
	character := character.NewCharacter()
	gw.SetCharacter(character, gw.Room().Area()[1][1])

	gw.Room().PrintRoom()
	fmt.Println()

	fmt.Println("Move Right")
	character.Move(room.Right)
	gw.Room().PrintRoom()
	fmt.Println()

	fmt.Println("Move Right")
	character.Move(room.Right)
	gw.Room().PrintRoom()
}
