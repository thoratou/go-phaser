// Automatic generation for PIXI.EarCut
// generated file EarCut.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type EarCut struct {
    *js.Object
}


// 
func NewEarCut() *EarCut {
    return &EarCut{js.Global.Get("PIXI").Get("EarCut").New()}
}

// 
func NewEarCutI(args ...interface{}) *EarCut {
    return &EarCut{js.Global.Get("PIXI").Get("EarCut").New(args)}
}




