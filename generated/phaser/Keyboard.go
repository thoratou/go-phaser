// Package phaser Automatic generation for Phaser.Keyboard
// generated file Keyboard.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Keyboard The Keyboard class monitors keyboard input and dispatches keyboard events.
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

// NewKeyboard The Keyboard class monitors keyboard input and dispatches keyboard events.
// 
// _Note_: many keyboards are unable to process certain combinations of keys due to hardware limitations known as ghosting.
// See http://www.html5gamedevs.com/topic/4876-impossible-to-use-more-than-2-keyboard-input-buttons-at-the-same-time/ for more details.
// 
// Also please be aware that certain browser extensions can disable or override Phaser keyboard handling.
// For example the Chrome extension vimium is known to disable Phaser from using the D key. And there are others.
// So please check your extensions before opening Phaser issues.
func NewKeyboard(game *Game) *Keyboard {
    return &Keyboard{js.Global.Get("Phaser").Get("Keyboard").New(game)}
}
// NewKeyboardI The Keyboard class monitors keyboard input and dispatches keyboard events.
// 
// _Note_: many keyboards are unable to process certain combinations of keys due to hardware limitations known as ghosting.
// See http://www.html5gamedevs.com/topic/4876-impossible-to-use-more-than-2-keyboard-input-buttons-at-the-same-time/ for more details.
// 
// Also please be aware that certain browser extensions can disable or override Phaser keyboard handling.
// For example the Chrome extension vimium is known to disable Phaser from using the D key. And there are others.
// So please check your extensions before opening Phaser issues.
func NewKeyboardI(args ...interface{}) *Keyboard {
    return &Keyboard{js.Global.Get("Phaser").Get("Keyboard").New(args)}
}



// Game Local reference to game.
func (self *Keyboard) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *Keyboard) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Enabled Keyboard input will only be processed if enabled.
func (self *Keyboard) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA Keyboard input will only be processed if enabled.
func (self *Keyboard) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// Event The most recent DOM event from keydown or keyup. This is updated every time a new key is pressed or released.
func (self *Keyboard) Event() interface{}{
    return self.Object.Get("event")
}

// SetEventA The most recent DOM event from keydown or keyup. This is updated every time a new key is pressed or released.
func (self *Keyboard) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// PressEvent The most recent DOM event from keypress.
func (self *Keyboard) PressEvent() interface{}{
    return self.Object.Get("pressEvent")
}

// SetPressEventA The most recent DOM event from keypress.
func (self *Keyboard) SetPressEventA(member interface{}) {
    self.Object.Set("pressEvent", member)
}

// CallbackContext The context under which the callbacks are run.
func (self *Keyboard) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which the callbacks are run.
func (self *Keyboard) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// OnDownCallback This callback is invoked every time a key is pressed down, including key repeats when a key is held down.
func (self *Keyboard) OnDownCallback() interface{}{
    return self.Object.Get("onDownCallback")
}

// SetOnDownCallbackA This callback is invoked every time a key is pressed down, including key repeats when a key is held down.
func (self *Keyboard) SetOnDownCallbackA(member interface{}) {
    self.Object.Set("onDownCallback", member)
}

// OnPressCallback This callback is invoked every time a DOM onkeypress event is raised, which is only for printable keys.
func (self *Keyboard) OnPressCallback() interface{}{
    return self.Object.Get("onPressCallback")
}

// SetOnPressCallbackA This callback is invoked every time a DOM onkeypress event is raised, which is only for printable keys.
func (self *Keyboard) SetOnPressCallbackA(member interface{}) {
    self.Object.Set("onPressCallback", member)
}

// OnUpCallback This callback is invoked every time a key is released.
func (self *Keyboard) OnUpCallback() interface{}{
    return self.Object.Get("onUpCallback")
}

// SetOnUpCallbackA This callback is invoked every time a key is released.
func (self *Keyboard) SetOnUpCallbackA(member interface{}) {
    self.Object.Set("onUpCallback", member)
}

// LastChar Returns the string value of the most recently pressed key.
func (self *Keyboard) LastChar() string{
    return self.Object.Get("lastChar").String()
}

// SetLastCharA Returns the string value of the most recently pressed key.
func (self *Keyboard) SetLastCharA(member string) {
    self.Object.Set("lastChar", member)
}

// LastKey Returns the most recently pressed Key. This is a Phaser.Key object and it changes every time a key is pressed.
func (self *Keyboard) LastKey() *Key{
    return &Key{self.Object.Get("lastKey")}
}

// SetLastKeyA Returns the most recently pressed Key. This is a Phaser.Key object and it changes every time a key is pressed.
func (self *Keyboard) SetLastKeyA(member *Key) {
    self.Object.Set("lastKey", member)
}


// AddCallbacks Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks(context interface{}) {
    self.Object.Call("addCallbacks", context)
}

// AddCallbacks1O Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks1O(context interface{}, onDown interface{}) {
    self.Object.Call("addCallbacks", context, onDown)
}

// AddCallbacks2O Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks2O(context interface{}, onDown interface{}, onUp interface{}) {
    self.Object.Call("addCallbacks", context, onDown, onUp)
}

// AddCallbacks3O Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks3O(context interface{}, onDown interface{}, onUp interface{}, onPress interface{}) {
    self.Object.Call("addCallbacks", context, onDown, onUp, onPress)
}

// AddCallbacksI Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacksI(args ...interface{}) {
    self.Object.Call("addCallbacks", args)
}

// AddKey If you need more fine-grained control over a Key you can create a new Phaser.Key object via this method.
// The Key object can then be polled, have events attached to it, etc.
func (self *Keyboard) AddKey(keycode int) *Key{
    return &Key{self.Object.Call("addKey", keycode)}
}

// AddKeyI If you need more fine-grained control over a Key you can create a new Phaser.Key object via this method.
// The Key object can then be polled, have events attached to it, etc.
func (self *Keyboard) AddKeyI(args ...interface{}) *Key{
    return &Key{self.Object.Call("addKey", args)}
}

// AddKeys A practical way to create an object containing user selected hotkeys.
// 
// For example,
// 
//     addKeys( { 'up': Phaser.KeyCode.W, 'down': Phaser.KeyCode.S, 'left': Phaser.KeyCode.A, 'right': Phaser.KeyCode.D } );
// 
// would return an object containing properties (`up`, `down`, `left` and `right`) referring to {@link Phaser.Key} object.
func (self *Keyboard) AddKeys(keys interface{}) interface{}{
    return self.Object.Call("addKeys", keys)
}

// AddKeysI A practical way to create an object containing user selected hotkeys.
// 
// For example,
// 
//     addKeys( { 'up': Phaser.KeyCode.W, 'down': Phaser.KeyCode.S, 'left': Phaser.KeyCode.A, 'right': Phaser.KeyCode.D } );
// 
// would return an object containing properties (`up`, `down`, `left` and `right`) referring to {@link Phaser.Key} object.
func (self *Keyboard) AddKeysI(args ...interface{}) interface{}{
    return self.Object.Call("addKeys", args)
}

// RemoveKey Removes a Key object from the Keyboard manager.
func (self *Keyboard) RemoveKey(keycode int) {
    self.Object.Call("removeKey", keycode)
}

// RemoveKeyI Removes a Key object from the Keyboard manager.
func (self *Keyboard) RemoveKeyI(args ...interface{}) {
    self.Object.Call("removeKey", args)
}

// CreateCursorKeys Creates and returns an object containing 4 hotkeys for Up, Down, Left and Right.
func (self *Keyboard) CreateCursorKeys() interface{}{
    return self.Object.Call("createCursorKeys")
}

// CreateCursorKeysI Creates and returns an object containing 4 hotkeys for Up, Down, Left and Right.
func (self *Keyboard) CreateCursorKeysI(args ...interface{}) interface{}{
    return self.Object.Call("createCursorKeys", args)
}

// Start Starts the Keyboard event listeners running (keydown and keyup). They are attached to the window.
// This is called automatically by Phaser.Input and should not normally be invoked directly.
func (self *Keyboard) Start() {
    self.Object.Call("start")
}

// StartI Starts the Keyboard event listeners running (keydown and keyup). They are attached to the window.
// This is called automatically by Phaser.Input and should not normally be invoked directly.
func (self *Keyboard) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Stop Stops the Keyboard event listeners from running (keydown, keyup and keypress). They are removed from the window.
func (self *Keyboard) Stop() {
    self.Object.Call("stop")
}

// StopI Stops the Keyboard event listeners from running (keydown, keyup and keypress). They are removed from the window.
func (self *Keyboard) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Destroy Stops the Keyboard event listeners from running (keydown and keyup). They are removed from the window.
// Also clears all key captures and currently created Key objects.
func (self *Keyboard) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Stops the Keyboard event listeners from running (keydown and keyup). They are removed from the window.
// Also clears all key captures and currently created Key objects.
func (self *Keyboard) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// AddKeyCapture By default when a key is pressed Phaser will not stop the event from propagating up to the browser.
// There are some keys this can be annoying for, like the arrow keys or space bar, which make the browser window scroll.
// 
// The `addKeyCapture` method enables consuming keyboard event for specific keys so it doesn't bubble up to the the browser
// and cause the default browser behavior.
// 
// Pass in either a single keycode or an array/hash of keycodes.
func (self *Keyboard) AddKeyCapture(keycode interface{}) {
    self.Object.Call("addKeyCapture", keycode)
}

// AddKeyCaptureI By default when a key is pressed Phaser will not stop the event from propagating up to the browser.
// There are some keys this can be annoying for, like the arrow keys or space bar, which make the browser window scroll.
// 
// The `addKeyCapture` method enables consuming keyboard event for specific keys so it doesn't bubble up to the the browser
// and cause the default browser behavior.
// 
// Pass in either a single keycode or an array/hash of keycodes.
func (self *Keyboard) AddKeyCaptureI(args ...interface{}) {
    self.Object.Call("addKeyCapture", args)
}

// RemoveKeyCapture Removes an existing key capture.
func (self *Keyboard) RemoveKeyCapture(keycode int) {
    self.Object.Call("removeKeyCapture", keycode)
}

// RemoveKeyCaptureI Removes an existing key capture.
func (self *Keyboard) RemoveKeyCaptureI(args ...interface{}) {
    self.Object.Call("removeKeyCapture", args)
}

// ClearCaptures Clear all set key captures.
func (self *Keyboard) ClearCaptures() {
    self.Object.Call("clearCaptures")
}

// ClearCapturesI Clear all set key captures.
func (self *Keyboard) ClearCapturesI(args ...interface{}) {
    self.Object.Call("clearCaptures", args)
}

// Update Updates all currently defined keys.
func (self *Keyboard) Update() {
    self.Object.Call("update")
}

// UpdateI Updates all currently defined keys.
func (self *Keyboard) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// ProcessKeyDown Process the keydown event.
func (self *Keyboard) ProcessKeyDown(event *KeyboardEvent) {
    self.Object.Call("processKeyDown", event)
}

// ProcessKeyDownI Process the keydown event.
func (self *Keyboard) ProcessKeyDownI(args ...interface{}) {
    self.Object.Call("processKeyDown", args)
}

// ProcessKeyPress Process the keypress event.
func (self *Keyboard) ProcessKeyPress(event *KeyboardEvent) {
    self.Object.Call("processKeyPress", event)
}

// ProcessKeyPressI Process the keypress event.
func (self *Keyboard) ProcessKeyPressI(args ...interface{}) {
    self.Object.Call("processKeyPress", args)
}

// ProcessKeyUp Process the keyup event.
func (self *Keyboard) ProcessKeyUp(event *KeyboardEvent) {
    self.Object.Call("processKeyUp", event)
}

// ProcessKeyUpI Process the keyup event.
func (self *Keyboard) ProcessKeyUpI(args ...interface{}) {
    self.Object.Call("processKeyUp", args)
}

// Reset Resets all Keys.
func (self *Keyboard) Reset() {
    self.Object.Call("reset")
}

// Reset1O Resets all Keys.
func (self *Keyboard) Reset1O(hard bool) {
    self.Object.Call("reset", hard)
}

// ResetI Resets all Keys.
func (self *Keyboard) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// DownDuration Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDuration(keycode int) bool{
    return self.Object.Call("downDuration", keycode).Bool()
}

// DownDuration1O Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDuration1O(keycode int, duration int) bool{
    return self.Object.Call("downDuration", keycode, duration).Bool()
}

// DownDurationI Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDurationI(args ...interface{}) bool{
    return self.Object.Call("downDuration", args).Bool()
}

// UpDuration Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDuration(keycode interface{}) bool{
    return self.Object.Call("upDuration", keycode).Bool()
}

// UpDuration1O Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDuration1O(keycode interface{}, duration int) bool{
    return self.Object.Call("upDuration", keycode, duration).Bool()
}

// UpDurationI Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDurationI(args ...interface{}) bool{
    return self.Object.Call("upDuration", args).Bool()
}

// IsDown Returns true of the key is currently pressed down. Note that it can only detect key presses on the web browser.
func (self *Keyboard) IsDown(keycode int) bool{
    return self.Object.Call("isDown", keycode).Bool()
}

// IsDownI Returns true of the key is currently pressed down. Note that it can only detect key presses on the web browser.
func (self *Keyboard) IsDownI(args ...interface{}) bool{
    return self.Object.Call("isDown", args).Bool()
}

