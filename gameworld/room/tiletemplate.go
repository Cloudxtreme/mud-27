package room

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

//Mark gets the mark of the template
func (tileTemplate *TileTemplate) Mark() string {
	return tileTemplate.mark
}

//Description gets the description of the template
func (tileTemplate *TileTemplate) Description() string {
	return tileTemplate.description
}

//TileType gets the tileType of the template
func (tileTemplate *TileTemplate) TileType() TileType {
	return tileTemplate.tileType
}
