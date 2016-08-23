// Automatic generation for Phaser.Easing
// generated file Easing.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A collection of easing methods defining ease-in and ease-out curves.
type Easing struct {
    *js.Object
}


// A collection of easing methods defining ease-in and ease-out curves.
func NewEasing() *Easing {
    return &Easing{js.Global.Get("Phaser").Get("Easing").New()}
}

// A collection of easing methods defining ease-in and ease-out curves.
func NewEasingI(args ...interface{}) *Easing {
    return &Easing{js.Global.Get("Phaser").Get("Easing").New(args)}
}




