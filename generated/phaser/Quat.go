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



// Quat Binding conversion method to Quat point 
func ToQuat(jsStruct interface{}) *Quat {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Quat{Object: object}
	}
	return nil
}




