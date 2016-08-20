// Automatic generation for PIXI.CanvasPool
// generated file CanvasPool.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// The CanvasPool is a global static object that allows Pixi and Phaser to pool canvas DOM elements.
type CanvasPool struct {
    *js.Object
}


// The pool into which the canvas dom elements are placed.
func (self *CanvasPool) GetPool() []interface{}{
	array := self.Get("pool")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The pool into which the canvas dom elements are placed.
func (self *CanvasPool) SetPool(member []interface{}) {
    self.Set("pool", member)
}



// Creates a new Canvas DOM element, or pulls one from the pool if free.
func (self *CanvasPool) CreateI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("create", args))
}

// Gets the first free canvas index from the pool.
func (self *CanvasPool) GetFirstI(args ...interface{}) float64{
    return self.Call("getFirst", args).Float()
}

// Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveI(args ...interface{}) {
    self.Call("remove", args)
}

// Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveByCanvasI(args ...interface{}) {
    self.Call("removeByCanvas", args)
}

// Gets the total number of used canvas elements in the pool.
func (self *CanvasPool) GetTotalI(args ...interface{}) float64{
    return self.Call("getTotal", args).Float()
}

// Gets the total number of free canvas elements in the pool.
func (self *CanvasPool) GetFreeI(args ...interface{}) float64{
    return self.Call("getFree", args).Float()
}
