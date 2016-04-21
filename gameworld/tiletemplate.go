package gameworld

//TileTemplate template for a tile
type TileTemplate struct {
	mark        string
	description string
	tileType    TileType
}

//NewTileTemplate creates a new TileTemplate
func NewTileTemplate(mark string, description string, tileType TileType) *TileTemplate {
	return &TileTemplate{mark: mark, description: description, tileType: tileType}
}
