package phaser

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

func WrapWebGLContext(o *js.Object) WebGLContext {
	context, _ := webgl.NewContext(o, nil)
	return *context
}

func WrapDOMElement(o *js.Object) DOMElement {
	return WrapElement(o)
}
