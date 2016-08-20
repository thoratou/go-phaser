package phaser

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
	"honnef.co/go/js/dom"
)

func WrapWebGLContext(o *js.Object) WebGLContext {
	context, _ := webgl.NewContext(o, nil)
	return WebGLContext(*context)
}

func WrapDOMElement(o *js.Object) DOMElement {
	return dom.WrapElement(o)
}

func WrapHTMLCanvasElement(o *js.Object) dom.HTMLCanvasElement {
	return dom.WrapHTMLElement(o).(dom.HTMLCanvasElement)
}

func WrapHTMLVideoElement(o *js.Object) dom.HTMLVideoElement {
	return dom.WrapHTMLElement(o).(dom.HTMLVideoElement)
}

func WrapCanvasRenderingContext2D(o *js.Object) dom.CanvasRenderingContext2D {
	return dom.CanvasRenderingContext2D{
		Object: o,
	}
}
