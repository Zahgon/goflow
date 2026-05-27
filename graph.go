// Package goflow implements a dataflow and flow-based programming library for Go.
package goflow

import (
	"reflect"
	"sync"
)

// GraphConfig sets up properties for a graph.
type GraphConfig struct {
	BufferSize int
}

// Graph represents a graph of processes connected with packet channels.
type Graph struct {
	conf                   GraphConfig      // Graph configuration
	waitGrp                *sync.WaitGroup  // Wait group for a graceful termination
	procs                  map[string]any   // Network processes
	inPorts                map[string]port  // Map of network incoming ports to component ports
	outPorts               map[string]port  // Map of network outgoing ports to component ports
	connections            []connection     // Network graph edges (inter-process connections)
	chanListenersCount     map[uintptr]uint // Tracks how many outports use the same channel
	chanListenersCountLock sync.Locker      // Used to synchronize operations on the chanListenersCount map
	closedChans            sync.Map         // Guards against double-close: stores closed channel pointers
	iips                   []iip            // Initial Information Packets to be sent to the network on start
}

// NewGraph returns a new initialized empty graph instance.
func NewGraph(config ...GraphConfig) *Graph { _ = "STUB: not implemented"; return nil }

// NewDefaultGraph is a ComponentConstructor for the factory.
func NewDefaultGraph() any {
	_ = "STUB: not implemented"

	// // Register an empty graph component in the registry
	//
	//	func init() {
	//		Register("Graph", NewDefaultGraph)
	//		Annotate("Graph", ComponentInfo{
	//			Description: "A clear graph",
	//			Icon:        "cogs",
	//		})
	//	}
	return *new(any)
}

// Add adds a new process with a given name to the network.
func (n *Graph) Add(name string, c any) error {
	_ = "STUB: not implemented"
	// c should be either graph or a component
	return nil
}

// Add to the map of processes

// AddGraph adds a new blank graph instance to a network. That instance can
// be modified then at run-time.
func (n *Graph) AddGraph(name string) error { _ = "STUB: not implemented"; return nil }

// AddNew creates a new process instance using component factory and adds it to the network.
func (n *Graph) AddNew(processName string, componentName string, f *Factory) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove deletes a process from the graph. First it stops the process if running.
// Then it disconnects it from other processes and removes the connections from
// the graph. Then it drops the process itself.
func (n *Graph) Remove(processName string) error { _ = "STUB: not implemented"; return nil }

// // Rename changes a process name in all connections, external ports, IIPs and the
// // graph itself.
// func (n *Graph) Rename(processName, newName string) bool {
// 	if _, exists := n.procs[processName]; !exists {
// 		return false
// 	}
// 	if _, busy := n.procs[newName]; busy {
// 		// New name is already taken
// 		return false
// 	}
// 	for i, conn := range n.connections {
// 		if conn.src.proc == processName {
// 			n.connections[i].src.proc = newName
// 		}
// 		if conn.tgt.proc == processName {
// 			n.connections[i].tgt.proc = newName
// 		}
// 	}
// 	for key, port := range n.inPorts {
// 		if port.proc == processName {
// 			tmp := n.inPorts[key]
// 			tmp.proc = newName
// 			n.inPorts[key] = tmp
// 		}
// 	}
// 	for key, port := range n.outPorts {
// 		if port.proc == processName {
// 			tmp := n.outPorts[key]
// 			tmp.proc = newName
// 			n.outPorts[key] = tmp
// 		}
// 	}
// 	n.procs[newName] = n.procs[processName]
// 	delete(n.procs, processName)
// 	return true
// }

// // Get returns a node contained in the network by its name.
// func (n *Graph) Get(processName string) interface{} {
// 	if proc, ok := n.procs[processName]; ok {
// 		return proc
// 	} else {
// 		panic("Process with name '" + processName + "' was not found")
// 	}
// }

// // getWait returns net's wait group.
// func (n *Graph) getWait() *sync.WaitGroup {
// 	return n.waitGrp
// }

// Process runs the network.
func (n *Graph) Process() { _ = "STUB: not implemented"; return }

// TODO provide a nicer way to handle graph errors

// closeChan closes a channel safely, ensuring it is only closed once.
// The sync.Map guards against double-close when multiple goroutines
// share the same underlying channel (e.g. fan-in, or IIP + process output).
func (n *Graph) closeChan(c reflect.Value) { _ = "STUB: not implemented"; return }

func (n *Graph) closeProcOuts(proc any) { _ = "STUB: not implemented"; return }
