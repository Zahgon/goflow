package goflow

import (
	"reflect"
)

// iip is the Initial Information Packet.
// IIPs are delivered to process input ports on the network start.
type iip struct {
	data any
	addr address
}

// AddIIP adds an Initial Information packet to the network.
func (n *Graph) AddIIP(processName, portName string, data any) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveIIP detaches an IIP from specific process and port.
func (n *Graph) RemoveIIP(processName, portName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove item from the slice

// sendIIPs sends Initial Information Packets upon network start.
func (n *Graph) sendIIPs() error {
	_ = "STUB: not implemented"
	// Send initial IPs
	return nil
}

// Get the receiver port channel

// Try to find a proc and attach a new channel to it

// Increase reference count for the channel

// Send data to the port

// channelByInPortAddr returns a channel by address from the network inports.
func (n *Graph) channelByInPortAddr(addr address) (channel reflect.Value, found bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// channelByConnectionAddr returns a channel by address from connections.
func (n *Graph) channelByConnectionAddr(addr address) (channel reflect.Value, found bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}
