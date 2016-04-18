package gameworld

import (
	"fmt"
)

//GameWorld hold information about the current world
type GameWorld struct {
	// TODO: refactor to level
	gameArea [][]*GameTile
}

//NewGameWorld creates a new gameWorld instance
func NewGameWorld(gameWorldTemplate [][]string) *GameWorld {
	gw := GameWorld{}

	//fill gameArea with defautlvalue
	count := 0
	gw.gameArea = make([][]*GameTile, len(gameWorldTemplate))
	for i := 0; i < len(gw.gameArea); i++ {
		gw.gameArea[i] = make([]*GameTile, len(gameWorldTemplate[i]))

		for j := 0; j < len(gw.gameArea[i]); j++ {
			//create tile
			gw.gameArea[i][j] = NewDefaultTile(count, gameWorldTemplate[i][j])
			count++
		}
	}

	//set connections
	for i := 0; i < len(gw.gameArea); i++ {
		for j := 0; j < len(gw.gameArea[i]); j++ {
			//set connections
			currentTile := gw.gameArea[i][j]

			//set connection up
			if (i - 1) > 0 {
				currentTile.SetConnetionTile(gw.gameArea[i-1][j], Up)
				gw.GameArea()[i-1][j].SetConnetionTile(currentTile, Down)
			}

			//set connection right
			if (j + 1) < len(gw.gameArea[i]) {
				currentTile.SetConnetionTile(gw.gameArea[i][j+1], Right)
				gw.gameArea[i][j+1].SetConnetionTile(currentTile, Left)
			}

			//set connection down
			if (i + 1) < len(gw.gameArea) {
				currentTile.SetConnetionTile(gw.gameArea[i+1][j], Down)
				gw.gameArea[i+1][j].SetConnetionTile(currentTile, Up)
			}

			//set connection left
			if (j - 1) > 0 {
				currentTile.SetConnetionTile(gw.gameArea[i][j-1], Left)
				gw.gameArea[i][j-1].SetConnetionTile(currentTile, Right)
			}
		}
	}

	return &gw
}

//PrintWorldMap prints the current world map
func (gw *GameWorld) PrintWorldMap() {
	for i := 0; i < len(gw.gameArea); i++ {
		for j := 0; j < len(gw.gameArea[i]); j++ {
			fmt.Printf("[%v]", gw.gameArea[i][j].Mark())
		}
		fmt.Println()
	}
}

//PrintDebug prints the current world map with debug information
func (gw *GameWorld) PrintDebug() {
	fmt.Println("Gameworld debug information!")
	fmt.Printf("Game Map: rows=%d\n", len(gw.gameArea))
	for i := 0; i < len(gw.gameArea); i++ {
		fmt.Printf("row %d ", i)
		for j := 0; j < len(gw.gameArea[i]); j++ {
			gw.gameArea[i][j].PrintDebug()
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

//GameArea return the game are of the game world
func (gw *GameWorld) GameArea() [][]*GameTile {
	return gw.gameArea
}
