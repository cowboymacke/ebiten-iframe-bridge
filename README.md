# Ebiten-iframe-bridge

A template for setting up communication between the Ebiten game engine and its host page via an iframe.

## Features

- Bidirectional communication between Ebiten (Go) and JavaScript
- Easy-to-use message passing system
- Template structure for quick project setup

## Prerequisites

- Go
- Make

## Installation

Clone the repository:

```sh
git clone https://github.com/yourusername/ebiten-iframe-bridge.git
cd ebiten-iframe-bridge
```

## Build and Run

### Build

```sh
make build
```

### Run

```sh
make web
```

### Build and Run in One Command

```sh
make build && make web
```

After running, navigate to http://localhost:8100 in your web browser.

## Usage

### Sending Messages

#### From Ebiten (Go) to JavaScript

In `main_wasm.go`, use the `messageToJs` function:

```go
messageToJs("hello world message from go")
```

#### From JavaScript to Ebiten (Go)

In your JavaScript code (e.g., `index.html`), use:

```js
window.sendEditorMessage('hello world from js')
```

### Receiving Messages

#### In Ebiten (Go)

In `main_wasm.go`, implement the `onJsMessage` function:

```go
func onJsMessage(_ js.Value, args []js.Value) any {
    // Handle args from JavaScript here
    return nil
}
```

#### In JavaScript

In your JavaScript code (e.g., `index.html`), implement:

```js
window.onEditorMessage = function (message) {
    // Handle messages from Go here
};
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT
