//go:build ignore
// +build ignore

package goflow

// Internal representation of NoFlo JSON format
type graphDescription struct {
	Properties struct {
		Name string
	}
	Processes map[string]struct {
		Component string
		Metadata  struct {
			Sync     bool  `json:",omitempty"`
			PoolSize int64 `json:",omitempty"`
		} `json:",omitempty"`
	}
	Connections []struct {
		Data interface{} `json:",omitempty"`
		Src  struct {
			Process string
			Port    string
		} `json:",omitempty"`
		Tgt struct {
			Process string
			Port    string
		}
		Metadata struct {
			Buffer int `json:",omitempty"`
		} `json:",omitempty"`
	}
	Exports []struct {
		Private string
		Public  string
	}
}

// ParseJSON converts a JSON network definition string into
// a flow.Graph object that can be run or used in other networks
func ParseJSON(js []byte) *Graph {
	_ = "STUB: not implemented"
	// Parse JSON into Go struct
	return nil
}

// fmt.Printf("%+v\n", descr)

// Create a new Graph

// Add processes to the network

// Process mode detection

// Add connections

// Check if it is an IIP or actual connection

// Add a connection

// Add an IIP

// Add port exports

// Split private into proc.port

// Try to detect port direction using reflection

// It's an inport

// It's an outport

// It's not a proper port

// TODO add support for subgraphs

// Register a component to be reused

// LoadJSON loads a JSON graph definition file into
// a flow.Graph object that can be run or used in other networks
func LoadJSON(filename string) *Graph { _ = "STUB: not implemented"; return nil }

// RegisterJSON registers an external JSON graph definition as a component
// that can be instantiated at run-time using component Factory.
// It returns true on success or false if component name is already taken.
func RegisterJSON(componentName, filePath string) bool { _ = "STUB: not implemented"; return false }
