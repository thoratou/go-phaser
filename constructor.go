package main

type Constructor struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	GoFunctions []Function  `json:"-"`
}
