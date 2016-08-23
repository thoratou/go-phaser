// Automatic generation for PIXI.WebGLGraphicsData
// generated file WebGLGraphicsData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLGraphicsData struct {
    *js.Object
}


// 
func NewWebGLGraphicsData() *WebGLGraphicsData {
    return &WebGLGraphicsData{js.Global.Get("PIXI").Get("WebGLGraphicsData").New()}
}

// 
func NewWebGLGraphicsDataI(args ...interface{}) *WebGLGraphicsData {
    return &WebGLGraphicsData{js.Global.Get("PIXI").Get("WebGLGraphicsData").New(args)}
}





// 
func (self *WebGLGraphicsData) Reset() {
    self.Object.Call("reset")
}

// 
func (self *WebGLGraphicsData) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// 
func (self *WebGLGraphicsData) Upload() {
    self.Object.Call("upload")
}

// 
func (self *WebGLGraphicsData) UploadI(args ...interface{}) {
    self.Object.Call("upload", args)
}
