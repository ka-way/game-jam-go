package main

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/ka-way/game-jam-go/actors"

	"image/color"
	_ "image/png"
)

const (
	fontSize = 24
)

var (
	// 画面サイズです。どこでも使えるようにグローバル変数にすると便利です。
	vw, vh = 640, 480
	// Updateが呼ばれることを、Ebitenでは tick と呼びます。
	tick = 0
	// Draw が呼ばれることを、Ebitenでは frame と呼びます。
	frame = 0
	// 画像パス
	img_mag = "assets/units/mag.png"

	// フォント
	arcadeFont font.Face
	// 操作プレイヤー
	player  actors.Player
	player2 actors.Player
)

// ebiten.Game interface を満たす型がEbitenには必要なので、この game 構造体に
// Update, Draw, Layout 関数を持たせます。
type game struct{}

// ゲームを初期化します。
func newGame() (*game, error) {
	g := &game{}

	// Load Assets.
	for _, name := range []string{
		img_mag,
	} {
		if err := loadImage(name); err != nil {
			return nil, err
		}
	}

	// 暫定初期化処理
	const dpi = 72
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	player = *actors.NewPlayer()
	player.Init(100, 200)
	player.InitAfterLoad(images)

	player2 = *actors.NewPlayer()
	player2.Init(200, 300)
	player2.InitAfterLoad(images)

	return g, nil
}

// Update関数は、画面のリフレッシュレートに関わらず
// 常に毎秒60回呼ばれます（既定値）。
// 描画ではなく更新処理を行うことが推奨されます。
func (g *game) Update() error {
	tick++

	player.Update(tick, vw, vh)
	player2.Update(tick, vw, vh)

	return nil
}

// Draw関数は、画面のリフレッシュレートと同期して呼ばれます（既定値）。
// 描画処理のみを行うことが推奨されます。ここで状態の変更を行うといろいろ事故ります。
func (g *game) Draw(screen *ebiten.Image) {
	frame++

	player.Draw(screen, images)
	player2.Draw(screen, images)

	text.Draw(screen, "testes", arcadeFont, 100, fontSize, color.White)
}

// Layout関数は、ウィンドウのリサイズの挙動を決定します。とりあえず常に画面サイズを返せば無難です。
func (g *game) Layout(ow, oh int) (int, int) {
	return vw, vh
}

// すべてのGoプログラムはmain packageのmain関数から始まります。
func main() {
	if err := _main(); err != nil {
		panic(err)
	}
}

func _main() error {
	g, err := newGame()
	if err != nil {
		return err
	}
	// ウィンドウタイトルを変更します。
	ebiten.SetWindowTitle("Ebitengine Game Jam")
	// ウィンドウサイズを決定します。
	ebiten.SetWindowSize(vw, vh)
	// ゲームスタート！
	return ebiten.RunGame(g)
}
