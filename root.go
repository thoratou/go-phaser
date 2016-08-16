package main

import (
	"encoding/json"
	"io"
)

type Root struct {
	Classes []Class `json:"classes"`
}

func Decode(r io.Reader) (root *Root, err error) {
	root = new(Root)
	err = json.NewDecoder(r).Decode(root)
	return
}
