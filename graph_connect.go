package goflow

import (
	"reflect"
)

// address is a full port accessor including the index part.
type address struct {
	proc  string // Process name
	port  string // Component port name
	key   string // Port key (only for map ports)
	index int    // Port index (only for array ports)
}

func (a address) String() string { _ = "STUB: not implemented"; return "" }

// connection stores information about a connection within the net.
type connection struct {
	src     address
	tgt     address
	channel reflect.Value
	buffer  int
}

// Connect a sender to a receiver and create a channel between them using BufferSize graph configuration.
// Normally such a connection is unbuffered but you can change by setting flow.DefaultBufferSize > 0 or
// by using ConnectBuf() function instead.
// It returns true on success or panics and returns false if error occurs.
func (n *Graph) Connect(senderName, senderPort, receiverName, receiverPort string) error {
	_ = "STUB: not implemented"
	return nil
}

// ConnectBuf connects a sender to a receiver using a channel with a buffer of a given size.
// It returns true on success or panics and returns false if error occurs.
func (n *Graph) ConnectBuf(senderName, senderPort, receiverName, receiverPort string, bufferSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// tells if a new channel will need to be created for this connection
// Try to find an existing outbound channel from the same sender,
// so it can be used as fan-out FIFO

// Then try to find an existing inbound channel to the same receiver,
// so it can be used as a fan-in FIFO

// Increase the number of listeners on this already used channel

// Register the first listener on a newly created channel

// Add connection info

// getProcPort finds an assignable port field in one of the subprocesses.
func (n *Graph) getProcPort(procName, portName string, dir reflect.ChanDir) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *

	// Check if process exists
	new(reflect.Value), nil
}

// Check if process is settable

// Get the port value

// Check if sender is a sub-graph

// Sender is a proc

func attachPort(port reflect.Value, addr address, dir reflect.ChanDir, ch reflect.Value, bufSize int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func attachChanPort(port reflect.Value, dir reflect.ChanDir, ch reflect.Value, bufSize int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func attachMapPort(port reflect.Value, key string, dir reflect.ChanDir, ch reflect.Value, bufSize int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func attachArrayPort(port reflect.Value, key int, dir reflect.ChanDir, ch reflect.Value, bufSize int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func validateChanDir(portType reflect.Type, dir reflect.ChanDir) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCanSet(portVal reflect.Value) error { _ = "STUB: not implemented"; return nil }

func selectOrMakeChan(new, existing reflect.Value, t reflect.Type, bufSize int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// parseAddress unfolds a string port name into parts, including array index or hashmap key.
func parseAddress(proc, port string) address { _ = "STUB: not implemented"; return *new(address) }

// capitalizePortName converts port names defined in UPPER or lower case to Title case,
// which is more common for structs in Go.
func capitalizePortName(name string) string { _ = "STUB: not implemented"; return "" }

// findExistingChan returns a channel attached to receiver if it already exists among connections.
func (n *Graph) findExistingChan(addr address, dir reflect.ChanDir) reflect.Value {
	_ = "STUB: not implemented"
	return *

	// Find existing channel attached to the receiver
	new(reflect.Value)
}

// incChanListenersCount increments SendChanRefCount.
// The count is needed when multiple senders are connected
// to the same receiver. When the network is terminated and
// senders need to close their output port, this counter
// can help to avoid closing the same channel multiple times.
func (n *Graph) incChanListenersCount(c reflect.Value) { _ = "STUB: not implemented"; return }

// decChanListenersCount decrements SendChanRefCount
// It returns true if the RefCount has reached 0.
func (n *Graph) decChanListenersCount(c reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// yes you may try to close a nonexistent channel, see what happens...

// // Disconnect removes a connection between sender's outport and receiver's inport.
// func (n *Graph) Disconnect(senderName, senderPort, receiverName, receiverPort string) bool {
// 	var sender, receiver interface{}
// 	var ok bool
// 	sender, ok = n.procs[senderName]
// 	if !ok {
// 		return false
// 	}
// 	receiver, ok = n.procs[receiverName]
// 	if !ok {
// 		return false
// 	}
// 	res := unsetProcPort(sender, senderPort, true)
// 	res = res && unsetProcPort(receiver, receiverPort, false)
// 	return res
// }

// // Unsets an port of a given process
// func unsetProcPort(proc interface{}, portName string, isOut bool) bool {
// 	v := reflect.ValueOf(proc)
// 	var ch reflect.Value
// 	if v.Elem().FieldByName("Graph").IsValid() {
// 		if subnet, ok := v.Elem().FieldByName("Graph").Addr().Interface().(*Graph); ok {
// 			if isOut {
// 				ch = subnet.getOutPort(portName)
// 			} else {
// 				ch = subnet.getInPort(portName)
// 			}
// 		} else {
// 			return false
// 		}
// 	} else {
// 		ch = v.Elem().FieldByName(portName)
// 	}
// 	if !ch.IsValid() {
// 		return false
// 	}
// 	ch.Set(reflect.Zero(ch.Type()))
// 	return true
// }
