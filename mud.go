package main

import (
	"fmt"

	"github.com/Norskan/mud/gameworld"
)

func main() {

	a := [][]string{}
	a = append(a, []string{"w", "w", "w", "w"})
	a = append(a, []string{"w", ".", ".", "w"})
	a = append(a, []string{"w", ".", ".", "w"})
	a = append(a, []string{"w", ".", ".", "w"})
	a = append(a, []string{"w", "w", "w", "w"})

	roomTemplate := gameworld.NewRoomTemplate(a)
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate(".", "Floor", gameworld.Moveable))
	roomTemplate.AddTileTemplate(gameworld.NewTileTemplate("w", "Wall", gameworld.NotMoveable))

	gw := gameworld.NewGameWorld(roomTemplate)
	character := gameworld.NewCharacter()
	gw.SetCharacter(character, gw.Room().Area()[1][1])

	gw.Room().PrintRoom()
	fmt.Println()

	fmt.Println("Move Right")
	character.Move(gameworld.Right)
	gw.Room().PrintRoom()
	fmt.Println()

	fmt.Println("Move Right")
	character.Move(gameworld.Right)
	gw.Room().PrintRoom()
}
