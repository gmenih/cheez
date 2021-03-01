package game

import (
	"math"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

var tileSize float64 = 90
var halfSize = tileSize / 2

func getOffset(x, y uint8) (float64, float64) {
	return (float64(x) * tileSize), (float64(y) * tileSize)
}

func getTileFromMousePosition(v pixel.Vec) (uint8, uint8) {
	xTile := math.Floor(v.X / tileSize)
	yTile := math.Floor(v.Y / tileSize)

	if yTile < 8 && xTile < 8 && xTile >= 0 && yTile >= 0 {
		return uint8(xTile), uint8(yTile)
	}

	return 255, 255
}

func getTileColor(x, y uint8) pixel.RGBA {
	isBlack := ((x + y) % 2) >= 1
	if isBlack {
		return pixel.RGB(0.0470588235294118, 0.180392156862745, 0.388235294117647)
	}

	return pixel.RGB(0.568627450980392, 0.631372549019608, 0.741176470588235)
}

func (c *Game) drawTileBackground(x, y uint8) {
	offsetX, offsetY := getOffset(x, y)

	c.imd.Color = getTileColor(x, y)
	c.imd.Push(pixel.V(offsetX, offsetY))
	c.imd.Push(pixel.V(offsetX+tileSize, offsetY+tileSize))
	c.imd.Rectangle(0)
}

func (g *Game) drawFigure(x, y uint8) {
	offsetX, offsetY := getOffset(x, y)
	figure := g.Engine.GetPiece(x, y)

	if sprite, ok := g.sprites[figure]; ok {
		sprite.Draw(g.spriteCanvas, pixel.IM.Moved(pixel.V(offsetX+halfSize, offsetY+halfSize)))
	}
}

func (c *Game) highlightTile() {
	x, y := reverseTileInt(c.state.hoveredTile)
	if x != 255 && y != 255 {
		offsetX, offsetY := getOffset(x, y)

		c.imd.Color = colornames.Darkviolet
		c.imd.Push(pixel.V(offsetX, offsetY))
		c.imd.Push(pixel.V(offsetX+tileSize, offsetY+tileSize))
		c.imd.Rectangle(3)
	}
}

func (c *Game) drawBoard() {
	for y := uint8(0); y < 8; y++ {
		for x := uint8(0); x < 8; x++ {

			c.drawTileBackground(x, y)
			if c.Engine.GetPiece(x, y) != 0 && !c.state.isDragging(x, y) {
				c.drawFigure(x, y)
			}
		}
	}

	c.highlightTile()
}
