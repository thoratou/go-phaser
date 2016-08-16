package main

type Function struct {
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	Parameters        []Parameter `json:"parameters"`
	ReturnType        string      `json:"returns/type"`
	ReturnDescription string      `json:"returns/description"`
}
