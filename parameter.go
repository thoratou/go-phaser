package main

type Parameter struct {
	Name        string      `json:"name"`
	Type        interface{} `json:"type"`
	Description string      `json:"description"`
	Default     interface{} `json:"default"`
	Optional    interface{} `json:"optional"`
	Nullable    interface{} `json:"nullable"`
}
