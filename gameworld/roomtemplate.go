package gameworld

//RoomTemplate template for a room
type RoomTemplate struct {
	roomMapTemplate [][]string
	tileTemplates   map[string]*TileTemplate
}

//NewRoomTemplate creates a new room template
func NewRoomTemplate(roomMapTemplate [][]string) *RoomTemplate {
	roomTemplate := &RoomTemplate{roomMapTemplate: roomMapTemplate}
	roomTemplate.tileTemplates = make(map[string]*TileTemplate)
	return roomTemplate
}

//TileTemplate gets the string represents of the tile at the position x, y
func (roomTemplate *RoomTemplate) TileTemplate(x int, y int) string {
	return roomTemplate.roomMapTemplate[x][y]
}

//Height returns the room height
func (roomTemplate *RoomTemplate) Height() int {
	return len(roomTemplate.roomMapTemplate)
}

//Width retuns the room widht
func (roomTemplate *RoomTemplate) Width() int {
	return len(roomTemplate.roomMapTemplate[0])
}

//RoomMapTemplate gets the room map teamplate from RoomTemplate
func (roomTemplate *RoomTemplate) RoomMapTemplate() [][]string {
	return roomTemplate.roomMapTemplate
}

//TileTemplates gets all the TileTemplates in the RoomTemplate
func (roomTemplate *RoomTemplate) TileTemplates() map[string]*TileTemplate {
	return roomTemplate.tileTemplates
}

//AddTileTemplate adds a new TileTemplate to the RoomTemplate
func (roomTemplate *RoomTemplate) AddTileTemplate(tiletTemplate *TileTemplate) {
	roomTemplate.tileTemplates[tiletTemplate.Mark()] = tiletTemplate
}
