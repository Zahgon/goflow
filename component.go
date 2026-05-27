package goflow

// Component is a unit that can start a process.
type Component interface {
	Process()
}

// Done notifies that the process is finished.
type Done struct{}

// Wait is a channel signaling of a completion.
type Wait chan Done

// Run the component process.
func Run(c Component) Wait { _ = "STUB: not implemented"; return *new(Wait) }
