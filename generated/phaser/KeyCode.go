// Package phaser Automatic generation for Phaser.KeyCode
// generated file KeyCode.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// KeyCode A key code represents a physical key on a keyboard.
// 
// The KeyCode class contains commonly supported keyboard key codes which can be used
// as keycode`-parameters in several {@link Phaser.Keyboard} and {@link Phaser.Key} methods.
// 
// _Note_: These values should only be used indirectly, eg. as `Phaser.KeyCode.KEY`.
// Future versions may replace the actual values, such that they remain compatible with `keycode`-parameters.
// The current implementation maps to the {@link https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode KeyboardEvent.keyCode} property.
// 
// _Note_: Use `Phaser.KeyCode.KEY` instead of `Phaser.Keyboard.KEY` to refer to a key code;
// the latter approach is supported for compatibility.
type KeyCode struct {
    *js.Object
}

// NewKeyCode A key code represents a physical key on a keyboard.
// 
// The KeyCode class contains commonly supported keyboard key codes which can be used
// as keycode`-parameters in several {@link Phaser.Keyboard} and {@link Phaser.Key} methods.
// 
// _Note_: These values should only be used indirectly, eg. as `Phaser.KeyCode.KEY`.
// Future versions may replace the actual values, such that they remain compatible with `keycode`-parameters.
// The current implementation maps to the {@link https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode KeyboardEvent.keyCode} property.
// 
// _Note_: Use `Phaser.KeyCode.KEY` instead of `Phaser.Keyboard.KEY` to refer to a key code;
// the latter approach is supported for compatibility.
func NewKeyCode() *KeyCode {
    return &KeyCode{js.Global.Get("Phaser").Get("KeyCode").New()}
}
// NewKeyCodeI A key code represents a physical key on a keyboard.
// 
// The KeyCode class contains commonly supported keyboard key codes which can be used
// as keycode`-parameters in several {@link Phaser.Keyboard} and {@link Phaser.Key} methods.
// 
// _Note_: These values should only be used indirectly, eg. as `Phaser.KeyCode.KEY`.
// Future versions may replace the actual values, such that they remain compatible with `keycode`-parameters.
// The current implementation maps to the {@link https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode KeyboardEvent.keyCode} property.
// 
// _Note_: Use `Phaser.KeyCode.KEY` instead of `Phaser.Keyboard.KEY` to refer to a key code;
// the latter approach is supported for compatibility.
func NewKeyCodeI(args ...interface{}) *KeyCode {
    return &KeyCode{js.Global.Get("Phaser").Get("KeyCode").New(args)}
}



// KeyCode Binding conversion method to KeyCode point 
func ToKeyCode(jsStruct interface{}) *KeyCode {
    if object, ok := jsStruct.(*js.Object); ok {
		return &KeyCode{Object: object}
	}
	return nil
}



// A empty description
func (self *KeyCode) A() interface{}{
    return self.Object.Get("A")
}

// SetAA empty description
func (self *KeyCode) SetAA(member interface{}) {
    self.Object.Set("A", member)
}

// B empty description
func (self *KeyCode) B() interface{}{
    return self.Object.Get("B")
}

// SetBA empty description
func (self *KeyCode) SetBA(member interface{}) {
    self.Object.Set("B", member)
}

// C empty description
func (self *KeyCode) C() interface{}{
    return self.Object.Get("C")
}

// SetCA empty description
func (self *KeyCode) SetCA(member interface{}) {
    self.Object.Set("C", member)
}

// D empty description
func (self *KeyCode) D() interface{}{
    return self.Object.Get("D")
}

// SetDA empty description
func (self *KeyCode) SetDA(member interface{}) {
    self.Object.Set("D", member)
}

// E empty description
func (self *KeyCode) E() interface{}{
    return self.Object.Get("E")
}

// SetEA empty description
func (self *KeyCode) SetEA(member interface{}) {
    self.Object.Set("E", member)
}

// F empty description
func (self *KeyCode) F() interface{}{
    return self.Object.Get("F")
}

// SetFA empty description
func (self *KeyCode) SetFA(member interface{}) {
    self.Object.Set("F", member)
}

// G empty description
func (self *KeyCode) G() interface{}{
    return self.Object.Get("G")
}

// SetGA empty description
func (self *KeyCode) SetGA(member interface{}) {
    self.Object.Set("G", member)
}

// H empty description
func (self *KeyCode) H() interface{}{
    return self.Object.Get("H")
}

// SetHA empty description
func (self *KeyCode) SetHA(member interface{}) {
    self.Object.Set("H", member)
}

// I empty description
func (self *KeyCode) I() interface{}{
    return self.Object.Get("I")
}

// SetIA empty description
func (self *KeyCode) SetIA(member interface{}) {
    self.Object.Set("I", member)
}

// J empty description
func (self *KeyCode) J() interface{}{
    return self.Object.Get("J")
}

// SetJA empty description
func (self *KeyCode) SetJA(member interface{}) {
    self.Object.Set("J", member)
}

// K empty description
func (self *KeyCode) K() interface{}{
    return self.Object.Get("K")
}

// SetKA empty description
func (self *KeyCode) SetKA(member interface{}) {
    self.Object.Set("K", member)
}

// L empty description
func (self *KeyCode) L() interface{}{
    return self.Object.Get("L")
}

// SetLA empty description
func (self *KeyCode) SetLA(member interface{}) {
    self.Object.Set("L", member)
}

// M empty description
func (self *KeyCode) M() interface{}{
    return self.Object.Get("M")
}

// SetMA empty description
func (self *KeyCode) SetMA(member interface{}) {
    self.Object.Set("M", member)
}

// N empty description
func (self *KeyCode) N() interface{}{
    return self.Object.Get("N")
}

// SetNA empty description
func (self *KeyCode) SetNA(member interface{}) {
    self.Object.Set("N", member)
}

// O empty description
func (self *KeyCode) O() interface{}{
    return self.Object.Get("O")
}

// SetOA empty description
func (self *KeyCode) SetOA(member interface{}) {
    self.Object.Set("O", member)
}

// P empty description
func (self *KeyCode) P() interface{}{
    return self.Object.Get("P")
}

// SetPA empty description
func (self *KeyCode) SetPA(member interface{}) {
    self.Object.Set("P", member)
}

// Q empty description
func (self *KeyCode) Q() interface{}{
    return self.Object.Get("Q")
}

// SetQA empty description
func (self *KeyCode) SetQA(member interface{}) {
    self.Object.Set("Q", member)
}

// R empty description
func (self *KeyCode) R() interface{}{
    return self.Object.Get("R")
}

// SetRA empty description
func (self *KeyCode) SetRA(member interface{}) {
    self.Object.Set("R", member)
}

// S empty description
func (self *KeyCode) S() interface{}{
    return self.Object.Get("S")
}

// SetSA empty description
func (self *KeyCode) SetSA(member interface{}) {
    self.Object.Set("S", member)
}

// T empty description
func (self *KeyCode) T() interface{}{
    return self.Object.Get("T")
}

// SetTA empty description
func (self *KeyCode) SetTA(member interface{}) {
    self.Object.Set("T", member)
}

// U empty description
func (self *KeyCode) U() interface{}{
    return self.Object.Get("U")
}

// SetUA empty description
func (self *KeyCode) SetUA(member interface{}) {
    self.Object.Set("U", member)
}

// V empty description
func (self *KeyCode) V() interface{}{
    return self.Object.Get("V")
}

// SetVA empty description
func (self *KeyCode) SetVA(member interface{}) {
    self.Object.Set("V", member)
}

// W empty description
func (self *KeyCode) W() interface{}{
    return self.Object.Get("W")
}

// SetWA empty description
func (self *KeyCode) SetWA(member interface{}) {
    self.Object.Set("W", member)
}

// X empty description
func (self *KeyCode) X() interface{}{
    return self.Object.Get("X")
}

// SetXA empty description
func (self *KeyCode) SetXA(member interface{}) {
    self.Object.Set("X", member)
}

// Y empty description
func (self *KeyCode) Y() interface{}{
    return self.Object.Get("Y")
}

// SetYA empty description
func (self *KeyCode) SetYA(member interface{}) {
    self.Object.Set("Y", member)
}

// Z empty description
func (self *KeyCode) Z() interface{}{
    return self.Object.Get("Z")
}

// SetZA empty description
func (self *KeyCode) SetZA(member interface{}) {
    self.Object.Set("Z", member)
}

// ZERO empty description
func (self *KeyCode) ZERO() interface{}{
    return self.Object.Get("ZERO")
}

// SetZEROA empty description
func (self *KeyCode) SetZEROA(member interface{}) {
    self.Object.Set("ZERO", member)
}

// ONE empty description
func (self *KeyCode) ONE() interface{}{
    return self.Object.Get("ONE")
}

// SetONEA empty description
func (self *KeyCode) SetONEA(member interface{}) {
    self.Object.Set("ONE", member)
}

// TWO empty description
func (self *KeyCode) TWO() interface{}{
    return self.Object.Get("TWO")
}

// SetTWOA empty description
func (self *KeyCode) SetTWOA(member interface{}) {
    self.Object.Set("TWO", member)
}

// THREE empty description
func (self *KeyCode) THREE() interface{}{
    return self.Object.Get("THREE")
}

// SetTHREEA empty description
func (self *KeyCode) SetTHREEA(member interface{}) {
    self.Object.Set("THREE", member)
}

// FOUR empty description
func (self *KeyCode) FOUR() interface{}{
    return self.Object.Get("FOUR")
}

// SetFOURA empty description
func (self *KeyCode) SetFOURA(member interface{}) {
    self.Object.Set("FOUR", member)
}

// FIVE empty description
func (self *KeyCode) FIVE() interface{}{
    return self.Object.Get("FIVE")
}

// SetFIVEA empty description
func (self *KeyCode) SetFIVEA(member interface{}) {
    self.Object.Set("FIVE", member)
}

// SIX empty description
func (self *KeyCode) SIX() interface{}{
    return self.Object.Get("SIX")
}

// SetSIXA empty description
func (self *KeyCode) SetSIXA(member interface{}) {
    self.Object.Set("SIX", member)
}

// SEVEN empty description
func (self *KeyCode) SEVEN() interface{}{
    return self.Object.Get("SEVEN")
}

// SetSEVENA empty description
func (self *KeyCode) SetSEVENA(member interface{}) {
    self.Object.Set("SEVEN", member)
}

// EIGHT empty description
func (self *KeyCode) EIGHT() interface{}{
    return self.Object.Get("EIGHT")
}

// SetEIGHTA empty description
func (self *KeyCode) SetEIGHTA(member interface{}) {
    self.Object.Set("EIGHT", member)
}

// NINE empty description
func (self *KeyCode) NINE() interface{}{
    return self.Object.Get("NINE")
}

// SetNINEA empty description
func (self *KeyCode) SetNINEA(member interface{}) {
    self.Object.Set("NINE", member)
}

// NUMPAD_0 empty description
func (self *KeyCode) NUMPAD_0() interface{}{
    return self.Object.Get("NUMPAD_0")
}

// SetNUMPAD_0A empty description
func (self *KeyCode) SetNUMPAD_0A(member interface{}) {
    self.Object.Set("NUMPAD_0", member)
}

// NUMPAD_1 empty description
func (self *KeyCode) NUMPAD_1() interface{}{
    return self.Object.Get("NUMPAD_1")
}

// SetNUMPAD_1A empty description
func (self *KeyCode) SetNUMPAD_1A(member interface{}) {
    self.Object.Set("NUMPAD_1", member)
}

// NUMPAD_2 empty description
func (self *KeyCode) NUMPAD_2() interface{}{
    return self.Object.Get("NUMPAD_2")
}

// SetNUMPAD_2A empty description
func (self *KeyCode) SetNUMPAD_2A(member interface{}) {
    self.Object.Set("NUMPAD_2", member)
}

// NUMPAD_3 empty description
func (self *KeyCode) NUMPAD_3() interface{}{
    return self.Object.Get("NUMPAD_3")
}

// SetNUMPAD_3A empty description
func (self *KeyCode) SetNUMPAD_3A(member interface{}) {
    self.Object.Set("NUMPAD_3", member)
}

// NUMPAD_4 empty description
func (self *KeyCode) NUMPAD_4() interface{}{
    return self.Object.Get("NUMPAD_4")
}

// SetNUMPAD_4A empty description
func (self *KeyCode) SetNUMPAD_4A(member interface{}) {
    self.Object.Set("NUMPAD_4", member)
}

// NUMPAD_5 empty description
func (self *KeyCode) NUMPAD_5() interface{}{
    return self.Object.Get("NUMPAD_5")
}

// SetNUMPAD_5A empty description
func (self *KeyCode) SetNUMPAD_5A(member interface{}) {
    self.Object.Set("NUMPAD_5", member)
}

// NUMPAD_6 empty description
func (self *KeyCode) NUMPAD_6() interface{}{
    return self.Object.Get("NUMPAD_6")
}

// SetNUMPAD_6A empty description
func (self *KeyCode) SetNUMPAD_6A(member interface{}) {
    self.Object.Set("NUMPAD_6", member)
}

// NUMPAD_7 empty description
func (self *KeyCode) NUMPAD_7() interface{}{
    return self.Object.Get("NUMPAD_7")
}

// SetNUMPAD_7A empty description
func (self *KeyCode) SetNUMPAD_7A(member interface{}) {
    self.Object.Set("NUMPAD_7", member)
}

// NUMPAD_8 empty description
func (self *KeyCode) NUMPAD_8() interface{}{
    return self.Object.Get("NUMPAD_8")
}

// SetNUMPAD_8A empty description
func (self *KeyCode) SetNUMPAD_8A(member interface{}) {
    self.Object.Set("NUMPAD_8", member)
}

// NUMPAD_9 empty description
func (self *KeyCode) NUMPAD_9() interface{}{
    return self.Object.Get("NUMPAD_9")
}

// SetNUMPAD_9A empty description
func (self *KeyCode) SetNUMPAD_9A(member interface{}) {
    self.Object.Set("NUMPAD_9", member)
}

// NUMPAD_MULTIPLY empty description
func (self *KeyCode) NUMPAD_MULTIPLY() interface{}{
    return self.Object.Get("NUMPAD_MULTIPLY")
}

// SetNUMPAD_MULTIPLYA empty description
func (self *KeyCode) SetNUMPAD_MULTIPLYA(member interface{}) {
    self.Object.Set("NUMPAD_MULTIPLY", member)
}

// NUMPAD_ADD empty description
func (self *KeyCode) NUMPAD_ADD() interface{}{
    return self.Object.Get("NUMPAD_ADD")
}

// SetNUMPAD_ADDA empty description
func (self *KeyCode) SetNUMPAD_ADDA(member interface{}) {
    self.Object.Set("NUMPAD_ADD", member)
}

// NUMPAD_ENTER empty description
func (self *KeyCode) NUMPAD_ENTER() interface{}{
    return self.Object.Get("NUMPAD_ENTER")
}

// SetNUMPAD_ENTERA empty description
func (self *KeyCode) SetNUMPAD_ENTERA(member interface{}) {
    self.Object.Set("NUMPAD_ENTER", member)
}

// NUMPAD_SUBTRACT empty description
func (self *KeyCode) NUMPAD_SUBTRACT() interface{}{
    return self.Object.Get("NUMPAD_SUBTRACT")
}

// SetNUMPAD_SUBTRACTA empty description
func (self *KeyCode) SetNUMPAD_SUBTRACTA(member interface{}) {
    self.Object.Set("NUMPAD_SUBTRACT", member)
}

// NUMPAD_DECIMAL empty description
func (self *KeyCode) NUMPAD_DECIMAL() interface{}{
    return self.Object.Get("NUMPAD_DECIMAL")
}

// SetNUMPAD_DECIMALA empty description
func (self *KeyCode) SetNUMPAD_DECIMALA(member interface{}) {
    self.Object.Set("NUMPAD_DECIMAL", member)
}

// NUMPAD_DIVIDE empty description
func (self *KeyCode) NUMPAD_DIVIDE() interface{}{
    return self.Object.Get("NUMPAD_DIVIDE")
}

// SetNUMPAD_DIVIDEA empty description
func (self *KeyCode) SetNUMPAD_DIVIDEA(member interface{}) {
    self.Object.Set("NUMPAD_DIVIDE", member)
}

// F1 empty description
func (self *KeyCode) F1() interface{}{
    return self.Object.Get("F1")
}

// SetF1A empty description
func (self *KeyCode) SetF1A(member interface{}) {
    self.Object.Set("F1", member)
}

// F2 empty description
func (self *KeyCode) F2() interface{}{
    return self.Object.Get("F2")
}

// SetF2A empty description
func (self *KeyCode) SetF2A(member interface{}) {
    self.Object.Set("F2", member)
}

// F3 empty description
func (self *KeyCode) F3() interface{}{
    return self.Object.Get("F3")
}

// SetF3A empty description
func (self *KeyCode) SetF3A(member interface{}) {
    self.Object.Set("F3", member)
}

// F4 empty description
func (self *KeyCode) F4() interface{}{
    return self.Object.Get("F4")
}

// SetF4A empty description
func (self *KeyCode) SetF4A(member interface{}) {
    self.Object.Set("F4", member)
}

// F5 empty description
func (self *KeyCode) F5() interface{}{
    return self.Object.Get("F5")
}

// SetF5A empty description
func (self *KeyCode) SetF5A(member interface{}) {
    self.Object.Set("F5", member)
}

// F6 empty description
func (self *KeyCode) F6() interface{}{
    return self.Object.Get("F6")
}

// SetF6A empty description
func (self *KeyCode) SetF6A(member interface{}) {
    self.Object.Set("F6", member)
}

// F7 empty description
func (self *KeyCode) F7() interface{}{
    return self.Object.Get("F7")
}

// SetF7A empty description
func (self *KeyCode) SetF7A(member interface{}) {
    self.Object.Set("F7", member)
}

// F8 empty description
func (self *KeyCode) F8() interface{}{
    return self.Object.Get("F8")
}

// SetF8A empty description
func (self *KeyCode) SetF8A(member interface{}) {
    self.Object.Set("F8", member)
}

// F9 empty description
func (self *KeyCode) F9() interface{}{
    return self.Object.Get("F9")
}

// SetF9A empty description
func (self *KeyCode) SetF9A(member interface{}) {
    self.Object.Set("F9", member)
}

// F10 empty description
func (self *KeyCode) F10() interface{}{
    return self.Object.Get("F10")
}

// SetF10A empty description
func (self *KeyCode) SetF10A(member interface{}) {
    self.Object.Set("F10", member)
}

// F11 empty description
func (self *KeyCode) F11() interface{}{
    return self.Object.Get("F11")
}

// SetF11A empty description
func (self *KeyCode) SetF11A(member interface{}) {
    self.Object.Set("F11", member)
}

// F12 empty description
func (self *KeyCode) F12() interface{}{
    return self.Object.Get("F12")
}

// SetF12A empty description
func (self *KeyCode) SetF12A(member interface{}) {
    self.Object.Set("F12", member)
}

// F13 empty description
func (self *KeyCode) F13() interface{}{
    return self.Object.Get("F13")
}

// SetF13A empty description
func (self *KeyCode) SetF13A(member interface{}) {
    self.Object.Set("F13", member)
}

// F14 empty description
func (self *KeyCode) F14() interface{}{
    return self.Object.Get("F14")
}

// SetF14A empty description
func (self *KeyCode) SetF14A(member interface{}) {
    self.Object.Set("F14", member)
}

// F15 empty description
func (self *KeyCode) F15() interface{}{
    return self.Object.Get("F15")
}

// SetF15A empty description
func (self *KeyCode) SetF15A(member interface{}) {
    self.Object.Set("F15", member)
}

// COLON empty description
func (self *KeyCode) COLON() interface{}{
    return self.Object.Get("COLON")
}

// SetCOLONA empty description
func (self *KeyCode) SetCOLONA(member interface{}) {
    self.Object.Set("COLON", member)
}

// EQUALS empty description
func (self *KeyCode) EQUALS() interface{}{
    return self.Object.Get("EQUALS")
}

// SetEQUALSA empty description
func (self *KeyCode) SetEQUALSA(member interface{}) {
    self.Object.Set("EQUALS", member)
}

// COMMA empty description
func (self *KeyCode) COMMA() interface{}{
    return self.Object.Get("COMMA")
}

// SetCOMMAA empty description
func (self *KeyCode) SetCOMMAA(member interface{}) {
    self.Object.Set("COMMA", member)
}

// UNDERSCORE empty description
func (self *KeyCode) UNDERSCORE() interface{}{
    return self.Object.Get("UNDERSCORE")
}

// SetUNDERSCOREA empty description
func (self *KeyCode) SetUNDERSCOREA(member interface{}) {
    self.Object.Set("UNDERSCORE", member)
}

// PERIOD empty description
func (self *KeyCode) PERIOD() interface{}{
    return self.Object.Get("PERIOD")
}

// SetPERIODA empty description
func (self *KeyCode) SetPERIODA(member interface{}) {
    self.Object.Set("PERIOD", member)
}

// QUESTION_MARK empty description
func (self *KeyCode) QUESTION_MARK() interface{}{
    return self.Object.Get("QUESTION_MARK")
}

// SetQUESTION_MARKA empty description
func (self *KeyCode) SetQUESTION_MARKA(member interface{}) {
    self.Object.Set("QUESTION_MARK", member)
}

// TILDE empty description
func (self *KeyCode) TILDE() interface{}{
    return self.Object.Get("TILDE")
}

// SetTILDEA empty description
func (self *KeyCode) SetTILDEA(member interface{}) {
    self.Object.Set("TILDE", member)
}

// OPEN_BRACKET empty description
func (self *KeyCode) OPEN_BRACKET() interface{}{
    return self.Object.Get("OPEN_BRACKET")
}

// SetOPEN_BRACKETA empty description
func (self *KeyCode) SetOPEN_BRACKETA(member interface{}) {
    self.Object.Set("OPEN_BRACKET", member)
}

// BACKWARD_SLASH empty description
func (self *KeyCode) BACKWARD_SLASH() interface{}{
    return self.Object.Get("BACKWARD_SLASH")
}

// SetBACKWARD_SLASHA empty description
func (self *KeyCode) SetBACKWARD_SLASHA(member interface{}) {
    self.Object.Set("BACKWARD_SLASH", member)
}

// CLOSED_BRACKET empty description
func (self *KeyCode) CLOSED_BRACKET() interface{}{
    return self.Object.Get("CLOSED_BRACKET")
}

// SetCLOSED_BRACKETA empty description
func (self *KeyCode) SetCLOSED_BRACKETA(member interface{}) {
    self.Object.Set("CLOSED_BRACKET", member)
}

// QUOTES empty description
func (self *KeyCode) QUOTES() interface{}{
    return self.Object.Get("QUOTES")
}

// SetQUOTESA empty description
func (self *KeyCode) SetQUOTESA(member interface{}) {
    self.Object.Set("QUOTES", member)
}

// BACKSPACE empty description
func (self *KeyCode) BACKSPACE() interface{}{
    return self.Object.Get("BACKSPACE")
}

// SetBACKSPACEA empty description
func (self *KeyCode) SetBACKSPACEA(member interface{}) {
    self.Object.Set("BACKSPACE", member)
}

// TAB empty description
func (self *KeyCode) TAB() interface{}{
    return self.Object.Get("TAB")
}

// SetTABA empty description
func (self *KeyCode) SetTABA(member interface{}) {
    self.Object.Set("TAB", member)
}

// CLEAR empty description
func (self *KeyCode) CLEAR() interface{}{
    return self.Object.Get("CLEAR")
}

// SetCLEARA empty description
func (self *KeyCode) SetCLEARA(member interface{}) {
    self.Object.Set("CLEAR", member)
}

// ENTER empty description
func (self *KeyCode) ENTER() interface{}{
    return self.Object.Get("ENTER")
}

// SetENTERA empty description
func (self *KeyCode) SetENTERA(member interface{}) {
    self.Object.Set("ENTER", member)
}

// SHIFT empty description
func (self *KeyCode) SHIFT() interface{}{
    return self.Object.Get("SHIFT")
}

// SetSHIFTA empty description
func (self *KeyCode) SetSHIFTA(member interface{}) {
    self.Object.Set("SHIFT", member)
}

// CONTROL empty description
func (self *KeyCode) CONTROL() interface{}{
    return self.Object.Get("CONTROL")
}

// SetCONTROLA empty description
func (self *KeyCode) SetCONTROLA(member interface{}) {
    self.Object.Set("CONTROL", member)
}

// ALT empty description
func (self *KeyCode) ALT() interface{}{
    return self.Object.Get("ALT")
}

// SetALTA empty description
func (self *KeyCode) SetALTA(member interface{}) {
    self.Object.Set("ALT", member)
}

// CAPS_LOCK empty description
func (self *KeyCode) CAPS_LOCK() interface{}{
    return self.Object.Get("CAPS_LOCK")
}

// SetCAPS_LOCKA empty description
func (self *KeyCode) SetCAPS_LOCKA(member interface{}) {
    self.Object.Set("CAPS_LOCK", member)
}

// ESC empty description
func (self *KeyCode) ESC() interface{}{
    return self.Object.Get("ESC")
}

// SetESCA empty description
func (self *KeyCode) SetESCA(member interface{}) {
    self.Object.Set("ESC", member)
}

// SPACEBAR empty description
func (self *KeyCode) SPACEBAR() interface{}{
    return self.Object.Get("SPACEBAR")
}

// SetSPACEBARA empty description
func (self *KeyCode) SetSPACEBARA(member interface{}) {
    self.Object.Set("SPACEBAR", member)
}

// PAGE_UP empty description
func (self *KeyCode) PAGE_UP() interface{}{
    return self.Object.Get("PAGE_UP")
}

// SetPAGE_UPA empty description
func (self *KeyCode) SetPAGE_UPA(member interface{}) {
    self.Object.Set("PAGE_UP", member)
}

// PAGE_DOWN empty description
func (self *KeyCode) PAGE_DOWN() interface{}{
    return self.Object.Get("PAGE_DOWN")
}

// SetPAGE_DOWNA empty description
func (self *KeyCode) SetPAGE_DOWNA(member interface{}) {
    self.Object.Set("PAGE_DOWN", member)
}

// END empty description
func (self *KeyCode) END() interface{}{
    return self.Object.Get("END")
}

// SetENDA empty description
func (self *KeyCode) SetENDA(member interface{}) {
    self.Object.Set("END", member)
}

// HOME empty description
func (self *KeyCode) HOME() interface{}{
    return self.Object.Get("HOME")
}

// SetHOMEA empty description
func (self *KeyCode) SetHOMEA(member interface{}) {
    self.Object.Set("HOME", member)
}

// LEFT empty description
func (self *KeyCode) LEFT() interface{}{
    return self.Object.Get("LEFT")
}

// SetLEFTA empty description
func (self *KeyCode) SetLEFTA(member interface{}) {
    self.Object.Set("LEFT", member)
}

// UP empty description
func (self *KeyCode) UP() interface{}{
    return self.Object.Get("UP")
}

// SetUPA empty description
func (self *KeyCode) SetUPA(member interface{}) {
    self.Object.Set("UP", member)
}

// RIGHT empty description
func (self *KeyCode) RIGHT() interface{}{
    return self.Object.Get("RIGHT")
}

// SetRIGHTA empty description
func (self *KeyCode) SetRIGHTA(member interface{}) {
    self.Object.Set("RIGHT", member)
}

// DOWN empty description
func (self *KeyCode) DOWN() interface{}{
    return self.Object.Get("DOWN")
}

// SetDOWNA empty description
func (self *KeyCode) SetDOWNA(member interface{}) {
    self.Object.Set("DOWN", member)
}

// PLUS empty description
func (self *KeyCode) PLUS() interface{}{
    return self.Object.Get("PLUS")
}

// SetPLUSA empty description
func (self *KeyCode) SetPLUSA(member interface{}) {
    self.Object.Set("PLUS", member)
}

// MINUS empty description
func (self *KeyCode) MINUS() interface{}{
    return self.Object.Get("MINUS")
}

// SetMINUSA empty description
func (self *KeyCode) SetMINUSA(member interface{}) {
    self.Object.Set("MINUS", member)
}

// INSERT empty description
func (self *KeyCode) INSERT() interface{}{
    return self.Object.Get("INSERT")
}

// SetINSERTA empty description
func (self *KeyCode) SetINSERTA(member interface{}) {
    self.Object.Set("INSERT", member)
}

// DELETE empty description
func (self *KeyCode) DELETE() interface{}{
    return self.Object.Get("DELETE")
}

// SetDELETEA empty description
func (self *KeyCode) SetDELETEA(member interface{}) {
    self.Object.Set("DELETE", member)
}

// HELP empty description
func (self *KeyCode) HELP() interface{}{
    return self.Object.Get("HELP")
}

// SetHELPA empty description
func (self *KeyCode) SetHELPA(member interface{}) {
    self.Object.Set("HELP", member)
}

// NUM_LOCK empty description
func (self *KeyCode) NUM_LOCK() interface{}{
    return self.Object.Get("NUM_LOCK")
}

// SetNUM_LOCKA empty description
func (self *KeyCode) SetNUM_LOCKA(member interface{}) {
    self.Object.Set("NUM_LOCK", member)
}


