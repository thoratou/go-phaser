package main

import "strings"

type Class struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Constructor Constructor `json:"constructor"`
	Functions   []Function  `json:"functions"`
	Members     []Member    `json:"members"`
}

func (c *Class) GetNamespace() string {
	lastPoint := strings.LastIndex(c.Name, ".")
	if lastPoint != -1 {
		return c.Name[0:lastPoint]
	}
	return ""
}

func (c *Class) GetPackage() string {
	return strings.Replace(c.GetNamespace(), ".", "/", -1)
}

func (c *Class) GetNameNoNamespace() string {
	lastPoint := strings.LastIndex(c.Name, ".")
	if lastPoint != -1 {
		return c.Name[lastPoint+1 : len(c.Name)]
	}
	return c.Name
}

func (c *Class) GetDescriptionLines() []string {
	lines := strings.Replace(c.Description, "\r", "\n", -1)
	return strings.Split(lines, "\n")
}
