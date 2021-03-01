package game

import (
	"fmt"
	"gmenih341/gess/src/cheez/engine"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	Engine *engine.Engine

	state *gameState

	canvas       *pixelgl.Canvas
	imd          *imdraw.IMDraw
	spriteCanvas *pixelgl.Canvas
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
		loadCheezAssets(),
		win,
	}
}

func (c *Game) handleMouse() {
	relativePosition := c.window.Bounds().Center().Sub(c.canvas.Bounds().Center())
	x, y := getTileFromMousePosition(c.window.MousePosition().Sub(relativePosition))

	if x == 255 && y == 255 {
		return
	}

	c.state.highlightTile(x, y)

	if c.window.JustPressed(pixelgl.MouseButton1) {
		c.pickUpFigure(x, y)
	}

	if c.window.JustReleased(pixelgl.MouseButton1) {
		c.dropFigure(x, y)
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
