// Package phaser Automatic generation for PIXI.EarCut
// generated file EarCut.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EarCut empty description
type EarCut struct {
    *js.Object
}

// NewEarCut empty description
func NewEarCut() *EarCut {
    return &EarCut{js.Global.Get("PIXI").Get("EarCut").New()}
}
// NewEarCutI empty description
func NewEarCutI(args ...interface{}) *EarCut {
    return &EarCut{js.Global.Get("PIXI").Get("EarCut").New(args)}
}




