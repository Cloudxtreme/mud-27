package room

import "fmt"

//Room represents a area of the game
type Room struct {
	gameArea [][]*GameTile
}

//NewRoom creates a new Room from a template
func NewRoom(roomTemplate *Template) *Room {
	room := Room{}

	//fill gameArea with defautlvalue
	count := 0
	room.gameArea = make([][]*GameTile, roomTemplate.Height())
	for i := 0; i < len(room.gameArea); i++ {
		room.gameArea[i] = make([]*GameTile, roomTemplate.Width())

		for j := 0; j < roomTemplate.Width(); j++ {
			//create tile
			tileTemplate := roomTemplate.TileTemplates()[roomTemplate.TileTemplate(i, j)]

			if tileTemplate == nil {
				fmt.Printf("No template found for mark %s", roomTemplate.TileTemplate(i, j))
				return nil
			}

			room.gameArea[i][j] = NewTile(count, tileTemplate)
			count++
		}
	}

	//set connections
	for i := 0; i < len(room.gameArea); i++ {
		for j := 0; j < len(room.gameArea[i]); j++ {
			//set connections
			currentTile := room.gameArea[i][j]

			//set connection up
			if (i - 1) > 0 {
				currentTile.SetConnetionTile(room.gameArea[i-1][j], Up)
				room.gameArea[i-1][j].SetConnetionTile(currentTile, Down)
			}

			//set connection right
			if (j + 1) < len(room.gameArea[i]) {
				currentTile.SetConnetionTile(room.gameArea[i][j+1], Right)
				room.gameArea[i][j+1].SetConnetionTile(currentTile, Left)
			}

			//set connection down
			if (i + 1) < len(room.gameArea) {
				currentTile.SetConnetionTile(room.gameArea[i+1][j], Down)
				room.gameArea[i+1][j].SetConnetionTile(currentTile, Up)
			}

			//set connection left
			if (j - 1) > 0 {
				currentTile.SetConnetionTile(room.gameArea[i][j-1], Left)
				room.gameArea[i][j-1].SetConnetionTile(currentTile, Right)
			}
		}
	}

	return &room
}

//PrintRoom prints the current world map
func (room *Room) PrintRoom() {
	for i := 0; i < len(room.gameArea); i++ {
		for j := 0; j < len(room.gameArea[i]); j++ {
			fmt.Printf("[%v]", room.gameArea[i][j].Mark())
		}
		fmt.Println()
	}
}

//PrintRoomDebug prints the current world map with debug information
func (room *Room) PrintRoomDebug() {
	fmt.Println("Gameworld debug information!")
	fmt.Printf("Game Map: rows=%d\n", len(room.gameArea))
	for i := 0; i < len(room.gameArea); i++ {
		fmt.Printf("row %d ", i)
		for j := 0; j < len(room.gameArea[i]); j++ {
			room.gameArea[i][j].PrintDebug()
		}
		fmt.Println()
	}
}

//Area returns the GameTiles of the room
func (room *Room) Area() [][]*GameTile {
	return room.gameArea
}
