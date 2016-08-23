// Automatic generation for quat
// generated file Quat.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Quat struct {
    *js.Object
}


// 
func NewQuat() *Quat {
    return &Quat{js.Global.Get("quat").New()}
}

// 
func NewQuatI(args ...interface{}) *Quat {
    return &Quat{js.Global.Get("quat").New(args)}
}




