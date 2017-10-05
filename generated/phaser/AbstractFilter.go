// Package phaser Automatic generation for PIXI.AbstractFilter
// generated file AbstractFilter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// AbstractFilter This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// 
// If you want to make a custom filter this should be your base class.
type AbstractFilter struct {
    *js.Object
}

// NewAbstractFilter This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// 
// If you want to make a custom filter this should be your base class.
func NewAbstractFilter(fragmentSrc []interface{}, uniforms interface{}) *AbstractFilter {
    return &AbstractFilter{js.Global.Get("PIXI").Get("AbstractFilter").New(fragmentSrc, uniforms)}
}
// NewAbstractFilterI This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// 
// If you want to make a custom filter this should be your base class.
func NewAbstractFilterI(args ...interface{}) *AbstractFilter {
    return &AbstractFilter{js.Global.Get("PIXI").Get("AbstractFilter").New(args)}
}



// AbstractFilter Binding conversion method to AbstractFilter point 
func ToAbstractFilter(jsStruct interface{}) *AbstractFilter {
    if object, ok := jsStruct.(*js.Object); ok {
		return &AbstractFilter{Object: object}
	}
	return nil
}



// Dirty empty description
func (self *AbstractFilter) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA empty description
func (self *AbstractFilter) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Padding empty description
func (self *AbstractFilter) Padding() int{
    return self.Object.Get("padding").Int()
}

// SetPaddingA empty description
func (self *AbstractFilter) SetPaddingA(member int) {
    self.Object.Set("padding", member)
}


// SyncUniforms Syncs the uniforms between the class object and the shaders.
func (self *AbstractFilter) SyncUniforms() {
    self.Object.Call("syncUniforms")
}

// SyncUniformsI Syncs the uniforms between the class object and the shaders.
func (self *AbstractFilter) SyncUniformsI(args ...interface{}) {
    self.Object.Call("syncUniforms", args)
}

