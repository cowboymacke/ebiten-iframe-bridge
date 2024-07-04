package main

import (
	"log"
	"syscall/js"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game *Game
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func onJsMessage(_ js.Value, args []js.Value) any {
	consoleLog(args[0].String())

	messageToJs("hello world message from go")

	return nil
}

func main() {
	game = &Game{}

	js.Global().Set("sendEditorMessage", js.FuncOf(onJsMessage))
	js.Global().Call("booted")

	ebiten.SetWindowSize(800, 600)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func messageToJs(value any) {
	js.Global().Call("onEditorMessage", value)
}

func consoleLog(message string) {
	js.Global().Get("console").Call("log", message)
}
