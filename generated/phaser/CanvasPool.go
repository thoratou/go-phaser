// Package phaser Automatic generation for PIXI.CanvasPool
// generated file CanvasPool.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// CanvasPool The CanvasPool is a global static object that allows Pixi and Phaser to pool canvas DOM elements.
type CanvasPool struct {
    *js.Object
}

// NewCanvasPool The CanvasPool is a global static object that allows Pixi and Phaser to pool canvas DOM elements.
func NewCanvasPool() *CanvasPool {
    return &CanvasPool{js.Global.Get("PIXI").Get("CanvasPool").New()}
}
// NewCanvasPoolI The CanvasPool is a global static object that allows Pixi and Phaser to pool canvas DOM elements.
func NewCanvasPoolI(args ...interface{}) *CanvasPool {
    return &CanvasPool{js.Global.Get("PIXI").Get("CanvasPool").New(args)}
}



// CanvasPool Binding conversion method to CanvasPool point 
func ToCanvasPool(jsStruct interface{}) *CanvasPool {
    if object, ok := jsStruct.(*js.Object); ok {
		return &CanvasPool{Object: object}
	}
	return nil
}



// Pool The pool into which the canvas dom elements are placed.
func (self *CanvasPool) Pool() []interface{}{
	array00 := self.Object.Get("pool")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetPoolA The pool into which the canvas dom elements are placed.
func (self *CanvasPool) SetPoolA(member []interface{}) {
    self.Object.Set("pool", member)
}


// Create Creates a new Canvas DOM element, or pulls one from the pool if free.
func (self *CanvasPool) Create(parent interface{}, width int, height int) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent, width, height))
}

// CreateI Creates a new Canvas DOM element, or pulls one from the pool if free.
func (self *CanvasPool) CreateI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", args))
}

// GetFirst Gets the first free canvas index from the pool.
func (self *CanvasPool) GetFirst() int{
    return self.Object.Call("getFirst").Int()
}

// GetFirstI Gets the first free canvas index from the pool.
func (self *CanvasPool) GetFirstI(args ...interface{}) int{
    return self.Object.Call("getFirst", args).Int()
}

// Remove Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) Remove(parent interface{}) {
    self.Object.Call("remove", parent)
}

// RemoveI Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// RemoveByCanvas Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveByCanvas(canvas *dom.HTMLCanvasElement) {
    self.Object.Call("removeByCanvas", canvas)
}

// RemoveByCanvasI Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveByCanvasI(args ...interface{}) {
    self.Object.Call("removeByCanvas", args)
}

// GetTotal Gets the total number of used canvas elements in the pool.
func (self *CanvasPool) GetTotal() int{
    return self.Object.Call("getTotal").Int()
}

// GetTotalI Gets the total number of used canvas elements in the pool.
func (self *CanvasPool) GetTotalI(args ...interface{}) int{
    return self.Object.Call("getTotal", args).Int()
}

// GetFree Gets the total number of free canvas elements in the pool.
func (self *CanvasPool) GetFree() int{
    return self.Object.Call("getFree").Int()
}

// GetFreeI Gets the total number of free canvas elements in the pool.
func (self *CanvasPool) GetFreeI(args ...interface{}) int{
    return self.Object.Call("getFree", args).Int()
}

