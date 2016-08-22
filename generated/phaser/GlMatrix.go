// Automatic generation for glMatrix
// generated file GlMatrix.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type GlMatrix struct {
    *js.Object
}


// 
func NewGlMatrix() *GlMatrix {
    return &GlMatrix{js.Global.Call("glMatrix")}
}

// 
func NewGlMatrixI(args ...interface{}) *GlMatrix {
    return &GlMatrix{js.Global.Call("glMatrix", args)}
}




