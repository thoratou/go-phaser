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


// The Keyboard class monitors keyboard input and dispatches keyboard events.
// 
// _Note_: many keyboards are unable to process certain combinations of keys due to hardware limitations known as ghosting.
// See http://www.html5gamedevs.com/topic/4876-impossible-to-use-more-than-2-keyboard-input-buttons-at-the-same-time/ for more details.
// 
// Also please be aware that certain browser extensions can disable or override Phaser keyboard handling.
// For example the Chrome extension vimium is known to disable Phaser from using the D key. And there are others.
// So please check your extensions before opening Phaser issues.
func NewKeyboard(game *Game) *Keyboard {
    return &Keyboard{js.Global.Call("Phaser.Keyboard", game)}
}

// The Keyboard class monitors keyboard input and dispatches keyboard events.
// 
// _Note_: many keyboards are unable to process certain combinations of keys due to hardware limitations known as ghosting.
// See http://www.html5gamedevs.com/topic/4876-impossible-to-use-more-than-2-keyboard-input-buttons-at-the-same-time/ for more details.
// 
// Also please be aware that certain browser extensions can disable or override Phaser keyboard handling.
// For example the Chrome extension vimium is known to disable Phaser from using the D key. And there are others.
// So please check your extensions before opening Phaser issues.
func NewKeyboardI(args ...interface{}) *Keyboard {
    return &Keyboard{js.Global.Call("Phaser.Keyboard", args)}
}



// Local reference to game.
func (self *Keyboard) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *Keyboard) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Keyboard input will only be processed if enabled.
func (self *Keyboard) GetEnabledA() bool{
    return self.Object.Get("enabled").Bool()
}

// Keyboard input will only be processed if enabled.
func (self *Keyboard) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// The most recent DOM event from keydown or keyup. This is updated every time a new key is pressed or released.
func (self *Keyboard) GetEventA() interface{}{
    return self.Object.Get("event")
}

// The most recent DOM event from keydown or keyup. This is updated every time a new key is pressed or released.
func (self *Keyboard) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// The most recent DOM event from keypress.
func (self *Keyboard) GetPressEventA() interface{}{
    return self.Object.Get("pressEvent")
}

// The most recent DOM event from keypress.
func (self *Keyboard) SetPressEventA(member interface{}) {
    self.Object.Set("pressEvent", member)
}

// The context under which the callbacks are run.
func (self *Keyboard) GetCallbackContextA() interface{}{
    return self.Object.Get("callbackContext")
}

// The context under which the callbacks are run.
func (self *Keyboard) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// This callback is invoked every time a key is pressed down, including key repeats when a key is held down.
func (self *Keyboard) SetOnDownCallbackA(member func(...interface{})) {
    self.Object.Set("onDownCallback", member)
}

// This callback is invoked every time a DOM onkeypress event is raised, which is only for printable keys.
func (self *Keyboard) SetOnPressCallbackA(member func(...interface{})) {
    self.Object.Set("onPressCallback", member)
}

// This callback is invoked every time a key is released.
func (self *Keyboard) SetOnUpCallbackA(member func(...interface{})) {
    self.Object.Set("onUpCallback", member)
}

// Returns the string value of the most recently pressed key.
func (self *Keyboard) GetLastCharA() string{
    return self.Object.Get("lastChar").String()
}

// Returns the string value of the most recently pressed key.
func (self *Keyboard) SetLastCharA(member string) {
    self.Object.Set("lastChar", member)
}

// Returns the most recently pressed Key. This is a Phaser.Key object and it changes every time a key is pressed.
func (self *Keyboard) GetLastKeyA() *Key{
    return &Key{self.Object.Get("lastKey")}
}

// Returns the most recently pressed Key. This is a Phaser.Key object and it changes every time a key is pressed.
func (self *Keyboard) SetLastKeyA(member *Key) {
    self.Object.Set("lastKey", member)
}



// Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks(context interface{}) {
    self.Object.Call("addCallbacks", context)
}

// Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks1O(context interface{}, onDown func(...interface{})) {
    self.Object.Call("addCallbacks", context, onDown)
}

// Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks2O(context interface{}, onDown func(...interface{}), onUp func(...interface{})) {
    self.Object.Call("addCallbacks", context, onDown, onUp)
}

// Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacks3O(context interface{}, onDown func(...interface{}), onUp func(...interface{}), onPress func(...interface{})) {
    self.Object.Call("addCallbacks", context, onDown, onUp, onPress)
}

// Add callbacks to the Keyboard handler so that each time a key is pressed down or released the callbacks are activated.
func (self *Keyboard) AddCallbacksI(args ...interface{}) {
    self.Object.Call("addCallbacks", args)
}

// If you need more fine-grained control over a Key you can create a new Phaser.Key object via this method.
// The Key object can then be polled, have events attached to it, etc.
func (self *Keyboard) AddKey(keycode int) *Key{
    return &Key{self.Object.Call("addKey", keycode)}
}

// If you need more fine-grained control over a Key you can create a new Phaser.Key object via this method.
// The Key object can then be polled, have events attached to it, etc.
func (self *Keyboard) AddKeyI(args ...interface{}) *Key{
    return &Key{self.Object.Call("addKey", args)}
}

// A practical way to create an object containing user selected hotkeys.
// 
// For example,
// 
//     addKeys( { 'up': Phaser.KeyCode.W, 'down': Phaser.KeyCode.S, 'left': Phaser.KeyCode.A, 'right': Phaser.KeyCode.D } );
// 
// would return an object containing properties (`up`, `down`, `left` and `right`) referring to {@link Phaser.Key} object.
func (self *Keyboard) AddKeys(keys interface{}) interface{}{
    return self.Object.Call("addKeys", keys)
}

// A practical way to create an object containing user selected hotkeys.
// 
// For example,
// 
//     addKeys( { 'up': Phaser.KeyCode.W, 'down': Phaser.KeyCode.S, 'left': Phaser.KeyCode.A, 'right': Phaser.KeyCode.D } );
// 
// would return an object containing properties (`up`, `down`, `left` and `right`) referring to {@link Phaser.Key} object.
func (self *Keyboard) AddKeysI(args ...interface{}) interface{}{
    return self.Object.Call("addKeys", args)
}

// Removes a Key object from the Keyboard manager.
func (self *Keyboard) RemoveKey(keycode int) {
    self.Object.Call("removeKey", keycode)
}

// Removes a Key object from the Keyboard manager.
func (self *Keyboard) RemoveKeyI(args ...interface{}) {
    self.Object.Call("removeKey", args)
}

// Creates and returns an object containing 4 hotkeys for Up, Down, Left and Right.
func (self *Keyboard) CreateCursorKeys() interface{}{
    return self.Object.Call("createCursorKeys")
}

// Creates and returns an object containing 4 hotkeys for Up, Down, Left and Right.
func (self *Keyboard) CreateCursorKeysI(args ...interface{}) interface{}{
    return self.Object.Call("createCursorKeys", args)
}

// Starts the Keyboard event listeners running (keydown and keyup). They are attached to the window.
// This is called automatically by Phaser.Input and should not normally be invoked directly.
func (self *Keyboard) Start() {
    self.Object.Call("start")
}

// Starts the Keyboard event listeners running (keydown and keyup). They are attached to the window.
// This is called automatically by Phaser.Input and should not normally be invoked directly.
func (self *Keyboard) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Stops the Keyboard event listeners from running (keydown, keyup and keypress). They are removed from the window.
func (self *Keyboard) Stop() {
    self.Object.Call("stop")
}

// Stops the Keyboard event listeners from running (keydown, keyup and keypress). They are removed from the window.
func (self *Keyboard) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Stops the Keyboard event listeners from running (keydown and keyup). They are removed from the window.
// Also clears all key captures and currently created Key objects.
func (self *Keyboard) Destroy() {
    self.Object.Call("destroy")
}

// Stops the Keyboard event listeners from running (keydown and keyup). They are removed from the window.
// Also clears all key captures and currently created Key objects.
func (self *Keyboard) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// By default when a key is pressed Phaser will not stop the event from propagating up to the browser.
// There are some keys this can be annoying for, like the arrow keys or space bar, which make the browser window scroll.
// 
// The `addKeyCapture` method enables consuming keyboard event for specific keys so it doesn't bubble up to the the browser
// and cause the default browser behavior.
// 
// Pass in either a single keycode or an array/hash of keycodes.
func (self *Keyboard) AddKeyCapture(keycode interface{}) {
    self.Object.Call("addKeyCapture", keycode)
}

// By default when a key is pressed Phaser will not stop the event from propagating up to the browser.
// There are some keys this can be annoying for, like the arrow keys or space bar, which make the browser window scroll.
// 
// The `addKeyCapture` method enables consuming keyboard event for specific keys so it doesn't bubble up to the the browser
// and cause the default browser behavior.
// 
// Pass in either a single keycode or an array/hash of keycodes.
func (self *Keyboard) AddKeyCaptureI(args ...interface{}) {
    self.Object.Call("addKeyCapture", args)
}

// Removes an existing key capture.
func (self *Keyboard) RemoveKeyCapture(keycode int) {
    self.Object.Call("removeKeyCapture", keycode)
}

// Removes an existing key capture.
func (self *Keyboard) RemoveKeyCaptureI(args ...interface{}) {
    self.Object.Call("removeKeyCapture", args)
}

// Clear all set key captures.
func (self *Keyboard) ClearCaptures() {
    self.Object.Call("clearCaptures")
}

// Clear all set key captures.
func (self *Keyboard) ClearCapturesI(args ...interface{}) {
    self.Object.Call("clearCaptures", args)
}

// Updates all currently defined keys.
func (self *Keyboard) Update() {
    self.Object.Call("update")
}

// Updates all currently defined keys.
func (self *Keyboard) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Process the keydown event.
func (self *Keyboard) ProcessKeyDown(event *KeyboardEvent) {
    self.Object.Call("processKeyDown", event)
}

// Process the keydown event.
func (self *Keyboard) ProcessKeyDownI(args ...interface{}) {
    self.Object.Call("processKeyDown", args)
}

// Process the keypress event.
func (self *Keyboard) ProcessKeyPress(event *KeyboardEvent) {
    self.Object.Call("processKeyPress", event)
}

// Process the keypress event.
func (self *Keyboard) ProcessKeyPressI(args ...interface{}) {
    self.Object.Call("processKeyPress", args)
}

// Process the keyup event.
func (self *Keyboard) ProcessKeyUp(event *KeyboardEvent) {
    self.Object.Call("processKeyUp", event)
}

// Process the keyup event.
func (self *Keyboard) ProcessKeyUpI(args ...interface{}) {
    self.Object.Call("processKeyUp", args)
}

// Resets all Keys.
func (self *Keyboard) Reset() {
    self.Object.Call("reset")
}

// Resets all Keys.
func (self *Keyboard) Reset1O(hard bool) {
    self.Object.Call("reset", hard)
}

// Resets all Keys.
func (self *Keyboard) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDuration(keycode int) bool{
    return self.Object.Call("downDuration", keycode).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDuration1O(keycode int, duration int) bool{
    return self.Object.Call("downDuration", keycode, duration).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) DownDurationI(args ...interface{}) bool{
    return self.Object.Call("downDuration", args).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDuration(keycode interface{}) bool{
    return self.Object.Call("upDuration", keycode).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDuration1O(keycode interface{}, duration int) bool{
    return self.Object.Call("upDuration", keycode, duration).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Keyboard) UpDurationI(args ...interface{}) bool{
    return self.Object.Call("upDuration", args).Bool()
}

// Returns true of the key is currently pressed down. Note that it can only detect key presses on the web browser.
func (self *Keyboard) IsDown(keycode int) bool{
    return self.Object.Call("isDown", keycode).Bool()
}

// Returns true of the key is currently pressed down. Note that it can only detect key presses on the web browser.
func (self *Keyboard) IsDownI(args ...interface{}) bool{
    return self.Object.Call("isDown", args).Bool()
}
