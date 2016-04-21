package main

import (
	"fmt"

	"github.com/Norskan/mud/gameworld"
)

func main() {

	a := [][]string{}
	a = append(a, []string{"w", "w", "w", "w"})
	a = append(a, []string{"w", ".", ".", "W"})
	a = append(a, []string{"w", ".", ".", "W"})
	a = append(a, []string{"w", ".", ".", "W"})
	a = append(a, []string{"w", "w", "w", "w"})

	gw := gameworld.NewGameWorld(a)
	character := gameworld.NewCharacter()
	gw.SetCharacter(character, gw.Room().Area()[1][1])

	gw.Room().PrintRoom()
	fmt.Println()
	character.Move(gameworld.Right)
	gw.Room().PrintRoom()
}
