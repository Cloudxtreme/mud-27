package gameworld

import (
	"fmt"
)

//GameWorld hold information about the current world
type GameWorld struct {
	//GameArea area of the GameWorld
	// TODO: refactor to level
	GameArea [][]*GameTile
}

//NewGameWorld creates a new gameWorld instance
func NewGameWorld(gameWorldTemplate [][]string) *GameWorld {
	gw := GameWorld{}

	//fill gameArea with defautlvalue
	count := 0
	gw.GameArea = make([][]*GameTile, len(gameWorldTemplate))
	for i := 0; i < len(gw.GameArea); i++ {
		gw.GameArea[i] = make([]*GameTile, len(gameWorldTemplate[i]))

		for j := 0; j < len(gw.GameArea[i]); j++ {
			//create tile
			gw.GameArea[i][j] = NewDefaultTile(count, gameWorldTemplate[i][j])
			count++
		}
	}

	//set connections
	for i := 0; i < len(gw.GameArea); i++ {
		for j := 0; j < len(gw.GameArea[i]); j++ {
			//set connections
			currentTile := gw.GameArea[i][j]

			//set connection up
			if (i - 1) > 0 {
				currentTile.SetConnetionTile(gw.GameArea[i-1][j], Up)
			}

			//set connection right
			if (j + 1) < len(gw.GameArea[i]) {
				currentTile.SetConnetionTile(gw.GameArea[i][j+1], Right)
			}

			//set connection down
			if (i + 1) < len(gw.GameArea) {
				currentTile.SetConnetionTile(gw.GameArea[i+1][j], Down)
			}

			//set connection left
			if (j - 1) > 0 {
				currentTile.SetConnetionTile(gw.GameArea[i][j-1], Left)
			}
		}
	}

	return &gw
}

//PrintWorldMap prints the current world map
func (gw *GameWorld) PrintWorldMap() {
	for i := 0; i < len(gw.GameArea); i++ {
		for j := 0; j < len(gw.GameArea[i]); j++ {
			fmt.Printf("[%v]", gw.GameArea[i][j].Mark())
		}
		fmt.Println()
	}
}

//PrintDebug prints the current world map with debug information
func (gw *GameWorld) PrintDebug() {
	fmt.Println("Gameworld debug information!")
	fmt.Printf("Game Map: rows=%d\n", len(gw.GameArea))
	for i := 0; i < len(gw.GameArea); i++ {
		fmt.Printf("row %d ", i)
		for j := 0; j < len(gw.GameArea[i]); j++ {
			gw.GameArea[i][j].PrintDebug()
		}
		fmt.Println()
	}
}

//SetCharacter sets a character on a tile
func (gw *GameWorld) SetCharacter(character *Character, gameTile *GameTile) {
	gameTile.SetCharacter(character)
	character.tilePosition = gameTile
}

//MoveCharacter moves a character in the game world
func (gw *GameWorld) MoveCharacter(character *Character, direction Direction) {
	character.tilePosition.GetConnetionTile(direction).MoveCharacter(character)
}
