// Automatic generation for PIXI.AbstractFilter
// generated file AbstractFilter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// If you want to make a custom filter this should be your base class.
type AbstractFilter struct {
    *js.Object
}


// This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// If you want to make a custom filter this should be your base class.
func NewAbstractFilter(fragmentSrc []interface{}, uniforms interface{}) *AbstractFilter {
    return &AbstractFilter{js.Global.Call("PIXI.AbstractFilter", fragmentSrc, uniforms)}
}

// This is the base class for creating a PIXI filter. Currently only webGL supports filters.
// If you want to make a custom filter this should be your base class.
func NewAbstractFilterI(args ...interface{}) *AbstractFilter {
    return &AbstractFilter{js.Global.Call("PIXI.AbstractFilter", args)}
}



// 
func (self *AbstractFilter) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// 
func (self *AbstractFilter) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// 
func (self *AbstractFilter) GetPaddingA() int{
    return self.Object.Get("padding").Int()
}

// 
func (self *AbstractFilter) SetPaddingA(member int) {
    self.Object.Set("padding", member)
}



// Syncs the uniforms between the class object and the shaders.
func (self *AbstractFilter) SyncUniforms() {
    self.Object.Call("syncUniforms")
}

// Syncs the uniforms between the class object and the shaders.
func (self *AbstractFilter) SyncUniformsI(args ...interface{}) {
    self.Object.Call("syncUniforms", args)
}
