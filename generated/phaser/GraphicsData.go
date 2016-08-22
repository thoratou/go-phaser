// Automatic generation for PIXI.GraphicsData
// generated file GraphicsData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A GraphicsData object.
type GraphicsData struct {
    *js.Object
}


// A GraphicsData object.
func NewGraphicsData() *GraphicsData {
    return &GraphicsData{js.Global.Call("PIXI.GraphicsData")}
}

// A GraphicsData object.
func NewGraphicsDataI(args ...interface{}) *GraphicsData {
    return &GraphicsData{js.Global.Call("PIXI.GraphicsData", args)}
}




