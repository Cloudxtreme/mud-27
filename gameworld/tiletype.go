package gameworld

//TileType the type of a GameTile
type TileType int

const (
	//Moveable player can move onto
	Moveable TileType = iota
	//NotMoveable play can not move onto
	NotMoveable
)

var stringTileType = [...]string{"Movable", "NotMoveable"}

func (tileType TileType) String() string {
	return stringTileType[tileType]
}
