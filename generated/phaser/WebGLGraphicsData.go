// Package phaser Automatic generation for PIXI.WebGLGraphicsData
// generated file WebGLGraphicsData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLGraphicsData empty description
type WebGLGraphicsData struct {
    *js.Object
}

// NewWebGLGraphicsData empty description
func NewWebGLGraphicsData() *WebGLGraphicsData {
    return &WebGLGraphicsData{js.Global.Get("PIXI").Get("WebGLGraphicsData").New()}
}
// NewWebGLGraphicsDataI empty description
func NewWebGLGraphicsDataI(args ...interface{}) *WebGLGraphicsData {
    return &WebGLGraphicsData{js.Global.Get("PIXI").Get("WebGLGraphicsData").New(args)}
}




// Reset empty description
func (self *WebGLGraphicsData) Reset() {
    self.Object.Call("reset")
}

// ResetI empty description
func (self *WebGLGraphicsData) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Upload empty description
func (self *WebGLGraphicsData) Upload() {
    self.Object.Call("upload")
}

// UploadI empty description
func (self *WebGLGraphicsData) UploadI(args ...interface{}) {
    self.Object.Call("upload", args)
}

