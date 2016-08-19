package main

import "strings"

type Class struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Constructor Constructor       `json:"constructor"`
	Functions   []Function        `json:"functions"`
	Members     []Member          `json:"members"`
	Imports     map[string]string `json:"-"`
}

func (c *Class) GetNamespace() string {
	/*lastPoint := strings.LastIndex(c.Name, ".")
	if lastPoint != -1 {
		return c.Name[0:lastPoint]
	}
	return ""*/
	return "phaser"
}

func (c *Class) GetPackage() string {
	return strings.Replace(c.GetNamespace(), ".", "/", -1)
}

func (c *Class) GetNameNoNamespace() string {
	return TypeNoNamespace(c.Name)
}

func (c *Class) GetDescriptionLines() []string {
	return SplitMultilines(c.Description)
}
