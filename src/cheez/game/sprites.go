package game

import (
	"gmenih341/cheez/src/cheez/engine"

	// Needed so we can load the spritesheet PNG
	_ "image/png"

	"image"
	"os"

	"github.com/faiface/pixel"
)

func loadSvg(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func rectPos(x, y float64) pixel.Rect {
	return pixel.R(
		x*60,
		y*60,
		(x+1)*60,
		(y+1)*60,
	)
}

func loadCheezAssets() map[engine.Piece]*pixel.Sprite {
	pic, err := loadSvg("assets/spritesheet.png")
	if err != nil {
		panic(err)
	}

	// (0, 0) is in the bottom left corner!
	return map[engine.Piece]*pixel.Sprite{
		engine.Light | engine.Queen:  pixel.NewSprite(pic, rectPos(0, 0)),
		engine.Light | engine.King:   pixel.NewSprite(pic, rectPos(1, 0)),
		engine.Light | engine.Rook:   pixel.NewSprite(pic, rectPos(2, 0)),
		engine.Light | engine.Knight: pixel.NewSprite(pic, rectPos(3, 0)),
		engine.Light | engine.Bishop: pixel.NewSprite(pic, rectPos(4, 0)),
		engine.Light | engine.Pawn:   pixel.NewSprite(pic, rectPos(5, 0)),

		engine.Dark | engine.Queen:  pixel.NewSprite(pic, rectPos(0, 1)),
		engine.Dark | engine.King:   pixel.NewSprite(pic, rectPos(1, 1)),
		engine.Dark | engine.Rook:   pixel.NewSprite(pic, rectPos(2, 1)),
		engine.Dark | engine.Knight: pixel.NewSprite(pic, rectPos(3, 1)),
		engine.Dark | engine.Bishop: pixel.NewSprite(pic, rectPos(4, 1)),
		engine.Dark | engine.Pawn:   pixel.NewSprite(pic, rectPos(5, 1)),
	}
}
