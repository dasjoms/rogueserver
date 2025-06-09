package defs

// TrainingAction represents a command sent from a neural network client.
type TrainingAction struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args,omitempty"`
}
