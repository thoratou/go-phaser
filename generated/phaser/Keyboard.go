// Automatic generation for Phaser.Keyboard
// generated file Keyboard.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Keyboard class monitors keyboard input and dispatches keyboard events.
// 
// _Note_: many keyboards are unable to process certain combinations of keys due to hardware limitations known as ghosting.
// See http://www.html5gamedevs.com/topic/4876-impossible-to-use-more-than-2-keyboard-input-buttons-at-the-same-time/ for more details.
// 
// Also please be aware that certain browser extensions can disable or override Phaser keyboard handling.
// For example the Chrome extension vimium is known to disable Phaser from using the D key. And there are others.
// So please check your extensions before opening Phaser issues.
type Keyboard struct {
    *js.Object
}


// Local reference to game.
func (self *Keyboard) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *Keyboard) SetGame(member Game) {
    self.Set("game", member)
}

// Keyboard input will only be processed if enabled.
func (self *Keyboard) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// Keyboard input will only be processed if enabled.
func (self *Keyboard) SetEnabled(member bool) {
    self.Set("enabled", member)
}

// The most recent DOM event from keydown or keyup. This is updated every time a new key is pressed or released.
func (self *Keyboard) GetEvent() interface{}{
    return self.Get("event")
}

// The most recent DOM event from keydown or keyup. This is updated every time a new key is pressed or released.
func (self *Keyboard) SetEvent(member interface{}) {
    self.Set("event", member)
}

// The most recent DOM event from keypress.
func (self *Keyboard) GetPressEvent() interface{}{
    return self.Get("pressEvent")
}

// The most recent DOM event from keypress.
func (self *Keyboard) SetPressEvent(member interface{}) {
    self.Set("pressEvent", member)
}

// The context under which the callbacks are run.
func (self *Keyboard) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which the callbacks are run.
func (self *Keyboard) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// This callback is invoked every time a key is pressed down, including key repeats when a key is held down.
func (self *Keyboard) SetOnDownCallback(member func(...interface{})) {
    self.Set("onDownCallback", member)
}

// This callback is invoked every time a DOM onkeypress event is raised, which is only for printable keys.
func (self *Keyboard) SetOnPressCallback(member func(...interface{})) {
    self.Set("onPressCallback", member)
}

// This callback is invoked every time a key is released.
func (self *Keyboard) SetOnUpCallback(member func(...interface{})) {
    self.Set("onUpCallback", member)
}

// Returns the string value of the most recently pressed key.
func (self *Keyboard) GetLastChar() string{
    return self.Get("lastChar").String()
}

// Returns the string value of the most recently pressed key.
func (self *Keyboard) SetLastChar(member string) {
    self.Set("lastChar", member)
}

// Returns the most recently pressed Key. This is a Phaser.Key object and it changes every time a key is pressed.
func (self *Keyboard) GetLastKey() Key{
    return Key{self.Get("lastKey")}
}

// Returns the most recently pressed Key. This is a Phaser.Key object and it changes every time a key is pressed.
func (self *Keyboard) SetLastKey(member Key) {
    self.Set("lastKey", member)
}



// Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacksI(args ...interface{}) {
    self.Call("addCallbacks", args)
}

// If you need more fine-grained control over a Key you can create a new Phaser.Key object via this method.
// The Key object can then be polled, have events attached to it, etc.
func (self *Keyboard) AddKeyI(args ...interface{}) Key{
    return Key{self.Call("addKey", args)}
}

// A practical way to create an object containing user selected hotkeys.
// 
// For example,
// 
//     addKeys( { 'up': Phaser.KeyCode.W, 'down': Phaser.KeyCode.S, 'left': Phaser.KeyCode.A, 'right': Phaser.KeyCode.D } );
// 
// would return an object containing properties (`up`, `down`, `left` and `right`) referring to {@link Phaser.Key} object.
func (self *Keyboard) AddKeysI(args ...interface{}) interface{}{
    return self.Call("addKeys", args)
}

// Removes a Key object from the Keyboard manager.
func (self *Keyboard) RemoveKeyI(args ...interface{}) {
    self.Call("removeKey", args)
}

// Creates and returns an object containing 4 hotkeys for Up, Down, Left and Right.
func (self *Keyboard) CreateCursorKeysI(args ...interface{}) interface{}{
    return self.Call("createCursorKeys", args)
}

// Starts the Keyboard event listeners running (keydown and keyup). They are attached to the window.
// This is called automatically by Phaser.Input and should not normally be invoked directly.
func (self *Keyboard) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Stops the Keyboard event listeners from running (keydown, keyup and keypress). They are removed from the window.
func (self *Keyboard) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Stops the Keyboard event listeners from running (keydown and keyup). They are removed from the window.
// Also clears all key captures and currently created Key objects.
func (self *Keyboard) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// By default when a key is pressed Phaser will not stop the event from propagating up to the browser.
// There are some keys this can be annoying for, like the arrow keys or space bar, which make the browser window scroll.
// 
// The `addKeyCapture` method enables consuming keyboard event for specific keys so it doesn't bubble up to the the browser
// and cause the default browser behavior.
// 
// Pass in either a single keycode or an array/hash of keycodes.
func (self *Keyboard) AddKeyCaptureI(args ...interface{}) {
    self.Call("addKeyCapture", args)
}

// Removes an existing key capture.
func (self *Keyboard) RemoveKeyCaptureI(args ...interface{}) {
    self.Call("removeKeyCapture", args)
}

// Clear all set key captures.
func (self *Keyboard) ClearCapturesI(args ...interface{}) {
    self.Call("clearCaptures", args)
}

// Updates all currently defined keys.
func (self *Keyboard) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Process the keydown event.
func (self *Keyboard) ProcessKeyDownI(args ...interface{}) {
    self.Call("processKeyDown", args)
}

// Process the keypress event.
func (self *Keyboard) ProcessKeyPressI(args ...interface{}) {
    self.Call("processKeyPress", args)
}

// Process the keyup event.
func (self *Keyboard) ProcessKeyUpI(args ...interface{}) {
    self.Call("processKeyUp", args)
}

// Resets all Keys.
func (self *Keyboard) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDurationI(args ...interface{}) bool{
    return self.Call("downDuration", args).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDurationI(args ...interface{}) bool{
    return self.Call("upDuration", args).Bool()
}

// Returns true of the key is currently pressed down. Note that it can only detect key presses on the web browser.
func (self *Keyboard) IsDownI(args ...interface{}) bool{
    return self.Call("isDown", args).Bool()
}
