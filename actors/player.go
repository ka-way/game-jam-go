package actors

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x float64
	y float64
	w int
	h int

	imgopt *ebiten.DrawImageOptions

	xSpeed      float64
	ySpeed      float64
	rotateSpeed float64

	imgPath string
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Init(x float64, y float64) {
	p.imgopt = &ebiten.DrawImageOptions{}
	p.imgPath = "assets/units/mag.png"
	p.rotateSpeed = 30
	p.x = x
	p.y = y
	p.xSpeed = 5
	p.ySpeed = 1
}

func (p *Player) InitAfterLoad(images map[string]*ebiten.Image) {
	img := images[p.imgPath]
	w, h := img.Size()
	p.w = w
	p.h = h

	p.imgopt.GeoM.Translate(float64(p.x), float64(p.y))
}

func (p *Player) center() (x float64, y float64) {
	return p.x + float64(p.w/2), p.y + float64(p.h/2)
}

func (p *Player) Update(tick int, vw int, vh int) {
	p.updateSpeed(tick, vw, vh)

	p.translate(tick, vw, vh)
	p.rotate(p.rotateSpeed / 360)
}

func (p *Player) updateSpeed(tick int, vw int, vh int) {
	centerX, centerY := p.center()
	if centerX > float64(vw) || centerX < 0 {
		p.xSpeed *= -1
	}
	if centerY > float64(vh) || centerY < 0 {
		p.ySpeed *= -1
	}
}

func (p *Player) translate(tick int, vw int, vh int) {
	p.x += p.xSpeed
	p.y += p.ySpeed
	p.imgopt.GeoM.Translate(p.xSpeed, p.ySpeed)
}

func (p *Player) rotate(theta float64) {
	centerX, centerY := p.center()
	p.imgopt.GeoM.Translate(-float64(centerX), -float64(centerY))
	p.imgopt.GeoM.Rotate(theta)
	p.imgopt.GeoM.Translate(float64(centerX), float64(centerY))
}

func (p *Player) Draw(screen *ebiten.Image, images map[string]*ebiten.Image) {
	screen.DrawImage(images[p.imgPath], p.imgopt)
}
