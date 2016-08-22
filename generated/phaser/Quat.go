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
    return &Quat{js.Global.Call("quat")}
}

// 
func NewQuatI(args ...interface{}) *Quat {
    return &Quat{js.Global.Call("quat", args)}
}




