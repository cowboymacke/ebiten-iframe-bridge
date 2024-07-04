build:
	env GOOS=js GOARCH=wasm go build -o web-server/ebit.wasm github.com/cowboymacke/ebiten-iframe-bridge
web:
	go run web-server/main.go -d web-server/
