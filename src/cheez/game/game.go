package game

import (
	"fmt"
	"gmenih341/cheez/src/cheez/engine"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	Engine *engine.Engine

	state *gameState

	canvas       *pixelgl.Canvas
	imd          *imdraw.IMDraw
	spriteCanvas *pixelgl.Canvas
	atlas        *text.Atlas
	sprites      map[engine.Piece]*pixel.Sprite
	window       *pixelgl.Window
}

func NewGame(win *pixelgl.Window) *Game {
	size := float64(8) * tileSize
	bounds := pixel.R(0, 0, size, size)
	return &Game{
		engine.NewEngine(time.Minute * 5),
		newState(),

		pixelgl.NewCanvas(bounds),
		imdraw.New(nil),
		pixelgl.NewCanvas(bounds),
		text.NewAtlas(basicfont.Face7x13, text.ASCII),
		loadCheezAssets(),
		win,
	}
}

func (g *Game) relativeMousePos() pixel.Vec {
	relativePosition := g.window.Bounds().Center().Sub(g.canvas.Bounds().Center())
	return g.window.MousePosition().Sub(relativePosition)
}

func (g *Game) handleMouse() {
	x, y := getTileFromMousePosition(g.relativeMousePos())

	if x == 255 && y == 255 {
		return
	}

	g.state.highlightTile(x, y)

	if g.window.JustPressed(pixelgl.MouseButton1) {
		g.pickUpFigure(x, y)
	}

	if g.window.JustReleased(pixelgl.MouseButton1) {
		g.dropFigure(x, y)
	}

	if g.window.JustReleased(pixelgl.MouseButton2) {
		g.state.setValidMoves(g.Engine.GetValidMoves(engine.T(x, y)))
	}
}

func (c *Game) clear() {
	c.window.Clear(colornames.Blanchedalmond)
	c.spriteCanvas.Clear(pixel.Alpha(0))
	c.canvas.Clear(pixel.Alpha(0))
	c.imd.Clear()
}

func (c *Game) Update(dt int) {
	c.handleMouse()
}

func (c *Game) Draw() {
	c.clear()

	c.drawBoard()
	c.imd.Draw(c.canvas)

	c.spriteCanvas.Draw(c.canvas, pixel.IM.Moved(c.canvas.Bounds().Center()))
	c.canvas.Draw(c.window, pixel.IM.Moved(c.window.Bounds().Center()))
}

func (g *Game) pickUpFigure(x, y uint8) {
	fmt.Println("Picking up", x, y)
	g.state.setDragging(x, y)
}

func (g *Game) dropFigure(x, y uint8) {
	oX, oY := g.state.resetDragging()

	if oX == x && oY == y {
		return
	}

	g.Engine.MovePiece(engine.T(oX, oY), engine.T(x, y))
}
