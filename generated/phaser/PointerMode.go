// Automatic generation for Phaser.PointerMode
// generated file PointerMode.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Enumeration categorizing operational modes of pointers.
// 
// PointerType values represent valid bitmasks.
// For example, a value representing both Mouse and Touch devices
// can be expressed as `PointerMode.CURSOR | PointerMode.CONTACT`.
// 
// Values may be added for future mode categorizations.
type PointerMode struct {
    *js.Object
}


// Enumeration categorizing operational modes of pointers.
// 
// PointerType values represent valid bitmasks.
// For example, a value representing both Mouse and Touch devices
// can be expressed as `PointerMode.CURSOR | PointerMode.CONTACT`.
// 
// Values may be added for future mode categorizations.
func NewPointerMode() *PointerMode {
    return &PointerMode{js.Global.Get("Phaser").Get("PointerMode").New()}
}

// Enumeration categorizing operational modes of pointers.
// 
// PointerType values represent valid bitmasks.
// For example, a value representing both Mouse and Touch devices
// can be expressed as `PointerMode.CURSOR | PointerMode.CONTACT`.
// 
// Values may be added for future mode categorizations.
func NewPointerModeI(args ...interface{}) *PointerMode {
    return &PointerMode{js.Global.Get("Phaser").Get("PointerMode").New(args)}
}



// A 'CURSOR' is a pointer with a *passive cursor* such as a mouse, touchpad, watcom stylus, or even TV-control arrow-pad.
// 
// It has the property that a cursor is passively moved without activating the input.
// This currently corresponds with {@link Phaser.Pointer#isMouse} property.
func (self *PointerMode) CURSOR() interface{}{
    return self.Object.Get("CURSOR")
}

// A 'CURSOR' is a pointer with a *passive cursor* such as a mouse, touchpad, watcom stylus, or even TV-control arrow-pad.
// 
// It has the property that a cursor is passively moved without activating the input.
// This currently corresponds with {@link Phaser.Pointer#isMouse} property.
func (self *PointerMode) SetCURSORA(member interface{}) {
    self.Object.Set("CURSOR", member)
}

// A 'CONTACT' pointer has an *active cursor* that only tracks movement when actived; notably this is a touch-style input.
func (self *PointerMode) CONTACT() interface{}{
    return self.Object.Get("CONTACT")
}

// A 'CONTACT' pointer has an *active cursor* that only tracks movement when actived; notably this is a touch-style input.
func (self *PointerMode) SetCONTACTA(member interface{}) {
    self.Object.Set("CONTACT", member)
}


