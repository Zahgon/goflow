//go:build ignore
// +build ignore

package goflow

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type protocolHandler func(*websocket.Conn, interface{})

// Runtime is a NoFlo-compatible runtime implementing the FBP protocol
type Runtime struct {
	// Unique runtime ID for use with Flowhub
	id string
	// Protocol command handlers
	handlers map[string]protocolHandler
	// Graphs created at runtime and exposed as components
	graphs map[string]*Graph
	// Main graph ID
	mainId string
	// Main graph
	main *Graph
	// Websocket server onReady signal
	ready chan struct{}
	// Websocket server onShutdown signal
	done chan struct{}
	// Gorilla Webscocket upgrader
	upgrader websocket.Upgrader
}

func sendJSON(ws *websocket.Conn, msg interface{}) { _ = "STUB: not implemented"; return }

// Register command handlers
func (r *Runtime) Init(name string) { _ = "STUB: not implemented"; return }

// Currently unsupported

// Currently unsupported

// Need to obtain ports annotation for the first time

// Id returns runtime's UUID v4
func (r *Runtime) Id() string {
	_ = "STUB: not implemented"

	// Ready returns a channel which is closed when the runtime is ready to work
	return ""
}

func (r *Runtime) Ready() chan struct{} {
	_ = "STUB: not implemented"

	// Stop tells the runtime to shut down
	return nil
}

func (r *Runtime) Stop() { _ = "STUB: not implemented"; return }

func (r *Runtime) Handle(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (r *Runtime) Listen(address string) { _ = "STUB: not implemented"; return }

// Wait for termination signal
