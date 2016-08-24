// Package phaser Automatic generation for PIXI.AbstractFilter
// generated file AbstractFilter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// AbstractFilter This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// If you want to make a custom filter this should be your base class.
type AbstractFilter struct {
    *js.Object
}

// NewAbstractFilter This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// If you want to make a custom filter this should be your base class.
func NewAbstractFilter(fragmentSrc []interface{}, uniforms interface{}) *AbstractFilter {
    return &AbstractFilter{js.Global.Get("PIXI").Get("AbstractFilter").New(fragmentSrc, uniforms)}
}
// NewAbstractFilterI This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// If you want to make a custom filter this should be your base class.
func NewAbstractFilterI(args ...interface{}) *AbstractFilter {
    return &AbstractFilter{js.Global.Get("PIXI").Get("AbstractFilter").New(args)}
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

