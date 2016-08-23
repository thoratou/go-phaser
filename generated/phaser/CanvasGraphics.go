// Automatic generation for PIXI.CanvasGraphics
// generated file CanvasGraphics.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A set of functions used by the canvas renderer to draw the primitive graphics data.
type CanvasGraphics struct {
    *js.Object
}


// A set of functions used by the canvas renderer to draw the primitive graphics data.
func NewCanvasGraphics() *CanvasGraphics {
    return &CanvasGraphics{js.Global.Get("PIXI").Get("CanvasGraphics").New()}
}

// A set of functions used by the canvas renderer to draw the primitive graphics data.
func NewCanvasGraphicsI(args ...interface{}) *CanvasGraphics {
    return &CanvasGraphics{js.Global.Get("PIXI").Get("CanvasGraphics").New(args)}
}




