// Package phaser Automatic generation for Phaser.Easing.Elastic
// generated file EasingElastic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingElastic Elastic easing.
type EasingElastic struct {
    *js.Object
}

// NewEasingElastic Elastic easing.
func NewEasingElastic() *EasingElastic {
    return &EasingElastic{js.Global.Get("Phaser").Get("Easing").Get("Elastic").New()}
}
// NewEasingElasticI Elastic easing.
func NewEasingElasticI(args ...interface{}) *EasingElastic {
    return &EasingElastic{js.Global.Get("Phaser").Get("Easing").Get("Elastic").New(args)}
}



// EasingElastic Binding conversion method to EasingElastic point 
func ToEasingElastic(jsStruct interface{}) *EasingElastic {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EasingElastic{Object: object}
	}
	return nil
}




// In Elastic ease-in.
func (self *EasingElastic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Elastic ease-in.
func (self *EasingElastic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Elastic ease-out.
func (self *EasingElastic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Elastic ease-out.
func (self *EasingElastic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Elastic ease-in/out.
func (self *EasingElastic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Elastic ease-in/out.
func (self *EasingElastic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

