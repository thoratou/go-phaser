// Package phaser Automatic generation for quat
// generated file Quat.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Quat empty description
type Quat struct {
    *js.Object
}

// NewQuat empty description
func NewQuat() *Quat {
    return &Quat{js.Global.Get("quat").New()}
}
// NewQuatI empty description
func NewQuatI(args ...interface{}) *Quat {
    return &Quat{js.Global.Get("quat").New(args)}
}




