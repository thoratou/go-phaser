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


// 
func (self *AbstractFilter) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// 
func (self *AbstractFilter) SetDirty(member bool) {
    self.Set("dirty", member)
}

// 
func (self *AbstractFilter) GetPadding() float64{
    return self.Get("padding").Float()
}

// 
func (self *AbstractFilter) SetPadding(member float64) {
    self.Set("padding", member)
}



// Syncs the uniforms between the class object and the shaders.
func (self *AbstractFilter) SyncUniformsI(args ...interface{}) {
    self.Call("syncUniforms", args)
}
