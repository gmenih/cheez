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
	isBlack := ((x + y) % 2) == 0
	if isBlack {
		return pixel.RGB(0.0470588235294118, 0.180392156862745, 0.388235294117647)
	}

	return pixel.RGB(0.568627450980392, 0.631372549019608, 0.741176470588235)
}

func (g *Game) drawTileBackground(x, y uint8) {
	offsetX, offsetY := getOffset(x, y)

	g.imd.Color = getTileColor(x, y)
	g.imd.Push(pixel.V(offsetX, offsetY))
	g.imd.Push(pixel.V(offsetX+tileSize, offsetY+tileSize))
	g.imd.Rectangle(0)
}

func (g *Game) drawFigure(x, y uint8) {
	offsetX, offsetY := getOffset(x, y)
	figure := g.Engine.GetPiece(x, y)

	if sprite, ok := g.sprites[figure]; ok {
		p := pixel.V(offsetX+halfSize, offsetY+halfSize)
		if g.state.isDragging(x, y) {
			p = g.relativeMousePos()
		}

		sprite.Draw(g.spriteCanvas, pixel.IM.Moved(p))
	}
}

func (g *Game) highlightTile() {
	x, y := reverseTileInt(g.state.hoveredTile)
	if x != 255 && y != 255 {
		offsetX, offsetY := getOffset(x, y)

		g.imd.Color = colornames.Darkviolet
		g.imd.Push(pixel.V(offsetX, offsetY))
		g.imd.Push(pixel.V(offsetX+tileSize, offsetY+tileSize))
		g.imd.Rectangle(3)
	}
}

func (g *Game) drawDraggingPiece() {
	if g.state.draggingTile != 255 {
		x, y := reverseTileInt(g.state.draggingTile)
		g.drawFigure(x, y)
	}
}

func (g *Game) drawBoard() {
	for y := uint8(0); y < 8; y++ {
		for x := uint8(0); x < 8; x++ {
			g.drawTileBackground(x, y)
			if g.Engine.GetPiece(x, y) != 0 && !g.state.isDragging(x, y) {
				g.drawFigure(x, y)
			}
		}
	}

	g.highlightTile()

	for _, t := range g.state.validMoves {
		offsetX, offsetY := getOffset(t.X, t.Y)

		g.imd.Color = pixel.RGB(0, 0, 0).Add(pixel.Alpha(0.5))
		g.imd.Push(pixel.V(offsetX+halfSize, offsetY+halfSize))
		g.imd.Circle(10, 0)
	}

	// always draw it last so it stays on top
	g.drawDraggingPiece()
}
