// Package phaser Automatic generation for Phaser.PointerMode
// generated file PointerMode.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PointerMode Enumeration categorizing operational modes of pointers.
// 
// PointerType values represent valid bitmasks.
// For example, a value representing both Mouse and Touch devices
// can be expressed as `PointerMode.CURSOR | PointerMode.CONTACT`.
// 
// Values may be added for future mode categorizations.
type PointerMode struct {
    *js.Object
}

// NewPointerMode Enumeration categorizing operational modes of pointers.
// 
// PointerType values represent valid bitmasks.
// For example, a value representing both Mouse and Touch devices
// can be expressed as `PointerMode.CURSOR | PointerMode.CONTACT`.
// 
// Values may be added for future mode categorizations.
func NewPointerMode() *PointerMode {
    return &PointerMode{js.Global.Get("Phaser").Get("PointerMode").New()}
}
// NewPointerModeI Enumeration categorizing operational modes of pointers.
// 
// PointerType values represent valid bitmasks.
// For example, a value representing both Mouse and Touch devices
// can be expressed as `PointerMode.CURSOR | PointerMode.CONTACT`.
// 
// Values may be added for future mode categorizations.
func NewPointerModeI(args ...interface{}) *PointerMode {
    return &PointerMode{js.Global.Get("Phaser").Get("PointerMode").New(args)}
}



// PointerMode Binding conversion method to PointerMode point 
func ToPointerMode(jsStruct interface{}) *PointerMode {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PointerMode{Object: object}
	}
	return nil
}



// CURSOR A 'CURSOR' is a pointer with a *passive cursor* such as a mouse, touchpad, watcom stylus, or even TV-control arrow-pad.
// 
// It has the property that a cursor is passively moved without activating the input.
// This currently corresponds with {@link Phaser.Pointer#isMouse} property.
func (self *PointerMode) CURSOR() interface{}{
    return self.Object.Get("CURSOR")
}

// SetCURSORA A 'CURSOR' is a pointer with a *passive cursor* such as a mouse, touchpad, watcom stylus, or even TV-control arrow-pad.
// 
// It has the property that a cursor is passively moved without activating the input.
// This currently corresponds with {@link Phaser.Pointer#isMouse} property.
func (self *PointerMode) SetCURSORA(member interface{}) {
    self.Object.Set("CURSOR", member)
}

// CONTACT A 'CONTACT' pointer has an *active cursor* that only tracks movement when actived; notably this is a touch-style input.
func (self *PointerMode) CONTACT() interface{}{
    return self.Object.Get("CONTACT")
}

// SetCONTACTA A 'CONTACT' pointer has an *active cursor* that only tracks movement when actived; notably this is a touch-style input.
func (self *PointerMode) SetCONTACTA(member interface{}) {
    self.Object.Set("CONTACT", member)
}


