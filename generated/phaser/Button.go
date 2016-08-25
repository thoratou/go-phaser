// Package phaser Automatic generation for Phaser.Button
// generated file Button.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Button Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
type Button struct {
    *js.Object
}

// NewButton Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton(game *Game) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game)}
}
// NewButton1O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton1O(game *Game, x int) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x)}
}
// NewButton2O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton2O(game *Game, x int, y int) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y)}
}
// NewButton3O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton3O(game *Game, x int, y int, key string) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key)}
}
// NewButton4O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton4O(game *Game, x int, y int, key string, callback interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key, callback)}
}
// NewButton5O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton5O(game *Game, x int, y int, key string, callback interface{}, callbackContext interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key, callback, callbackContext)}
}
// NewButton6O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton6O(game *Game, x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key, callback, callbackContext, overFrame)}
}
// NewButton7O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton7O(game *Game, x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}, outFrame interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key, callback, callbackContext, overFrame, outFrame)}
}
// NewButton8O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton8O(game *Game, x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key, callback, callbackContext, overFrame, outFrame, downFrame)}
}
// NewButton9O Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButton9O(game *Game, x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}, upFrame interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(game, x, y, key, callback, callbackContext, overFrame, outFrame, downFrame, upFrame)}
}
// NewButtonI Create a new `Button` object. A Button is a special type of Sprite that is set-up to handle Pointer events automatically.
// 
// The four states a Button responds to are:
// 
// * 'Over' - when the Pointer moves over the Button. This is also commonly known as 'hover'.
// * 'Out' - when the Pointer that was previously over the Button moves out of it.
// * 'Down' - when the Pointer is pressed down on the Button. I.e. touched on a touch enabled device or clicked with the mouse.
// * 'Up' - when the Pointer that was pressed down on the Button is released again.
// 
// A different texture/frame and activation sound can be specified for any of the states.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); the same values that can be used with a Sprite constructor.
func NewButtonI(args ...interface{}) *Button {
    return &Button{js.Global.Get("Phaser").Get("Button").New(args)}
}



// Button Binding conversion method to Button point 
func ToButton(jsStruct interface{}) *Button {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Button{Object: object}
	}
	return nil
}



// Type The Phaser Object Type.
func (self *Button) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The Phaser Object Type.
func (self *Button) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// PhysicsType The const physics body type of this object.
func (self *Button) PhysicsType() int{
    return self.Object.Get("physicsType").Int()
}

// SetPhysicsTypeA The const physics body type of this object.
func (self *Button) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// OnOverSound The Sound to be played when this Buttons Over state is activated.
func (self *Button) OnOverSound() interface{}{
    return self.Object.Get("onOverSound")
}

// SetOnOverSoundA The Sound to be played when this Buttons Over state is activated.
func (self *Button) SetOnOverSoundA(member interface{}) {
    self.Object.Set("onOverSound", member)
}

// OnOutSound The Sound to be played when this Buttons Out state is activated.
func (self *Button) OnOutSound() interface{}{
    return self.Object.Get("onOutSound")
}

// SetOnOutSoundA The Sound to be played when this Buttons Out state is activated.
func (self *Button) SetOnOutSoundA(member interface{}) {
    self.Object.Set("onOutSound", member)
}

// OnDownSound The Sound to be played when this Buttons Down state is activated.
func (self *Button) OnDownSound() interface{}{
    return self.Object.Get("onDownSound")
}

// SetOnDownSoundA The Sound to be played when this Buttons Down state is activated.
func (self *Button) SetOnDownSoundA(member interface{}) {
    self.Object.Set("onDownSound", member)
}

// OnUpSound The Sound to be played when this Buttons Up state is activated.
func (self *Button) OnUpSound() interface{}{
    return self.Object.Get("onUpSound")
}

// SetOnUpSoundA The Sound to be played when this Buttons Up state is activated.
func (self *Button) SetOnUpSoundA(member interface{}) {
    self.Object.Set("onUpSound", member)
}

// OnOverSoundMarker The Sound Marker used in conjunction with the onOverSound.
func (self *Button) OnOverSoundMarker() string{
    return self.Object.Get("onOverSoundMarker").String()
}

// SetOnOverSoundMarkerA The Sound Marker used in conjunction with the onOverSound.
func (self *Button) SetOnOverSoundMarkerA(member string) {
    self.Object.Set("onOverSoundMarker", member)
}

// OnOutSoundMarker The Sound Marker used in conjunction with the onOutSound.
func (self *Button) OnOutSoundMarker() string{
    return self.Object.Get("onOutSoundMarker").String()
}

// SetOnOutSoundMarkerA The Sound Marker used in conjunction with the onOutSound.
func (self *Button) SetOnOutSoundMarkerA(member string) {
    self.Object.Set("onOutSoundMarker", member)
}

// OnDownSoundMarker The Sound Marker used in conjunction with the onDownSound.
func (self *Button) OnDownSoundMarker() string{
    return self.Object.Get("onDownSoundMarker").String()
}

// SetOnDownSoundMarkerA The Sound Marker used in conjunction with the onDownSound.
func (self *Button) SetOnDownSoundMarkerA(member string) {
    self.Object.Set("onDownSoundMarker", member)
}

// OnUpSoundMarker The Sound Marker used in conjunction with the onUpSound.
func (self *Button) OnUpSoundMarker() string{
    return self.Object.Get("onUpSoundMarker").String()
}

// SetOnUpSoundMarkerA The Sound Marker used in conjunction with the onUpSound.
func (self *Button) SetOnUpSoundMarkerA(member string) {
    self.Object.Set("onUpSoundMarker", member)
}

// OnInputOver The Signal (or event) dispatched when this Button is in an Over state.
func (self *Button) OnInputOver() *Signal{
    return &Signal{self.Object.Get("onInputOver")}
}

// SetOnInputOverA The Signal (or event) dispatched when this Button is in an Over state.
func (self *Button) SetOnInputOverA(member *Signal) {
    self.Object.Set("onInputOver", member)
}

// OnInputOut The Signal (or event) dispatched when this Button is in an Out state.
func (self *Button) OnInputOut() *Signal{
    return &Signal{self.Object.Get("onInputOut")}
}

// SetOnInputOutA The Signal (or event) dispatched when this Button is in an Out state.
func (self *Button) SetOnInputOutA(member *Signal) {
    self.Object.Set("onInputOut", member)
}

// OnInputDown The Signal (or event) dispatched when this Button is in an Down state.
func (self *Button) OnInputDown() *Signal{
    return &Signal{self.Object.Get("onInputDown")}
}

// SetOnInputDownA The Signal (or event) dispatched when this Button is in an Down state.
func (self *Button) SetOnInputDownA(member *Signal) {
    self.Object.Set("onInputDown", member)
}

// OnInputUp The Signal (or event) dispatched when this Button is in an Up state.
func (self *Button) OnInputUp() *Signal{
    return &Signal{self.Object.Get("onInputUp")}
}

// SetOnInputUpA The Signal (or event) dispatched when this Button is in an Up state.
func (self *Button) SetOnInputUpA(member *Signal) {
    self.Object.Set("onInputUp", member)
}

// OnOverMouseOnly If true then onOver events (such as onOverSound) will only be triggered if the Pointer object causing them was the Mouse Pointer.
// The frame will still be changed as applicable.
func (self *Button) OnOverMouseOnly() bool{
    return self.Object.Get("onOverMouseOnly").Bool()
}

// SetOnOverMouseOnlyA If true then onOver events (such as onOverSound) will only be triggered if the Pointer object causing them was the Mouse Pointer.
// The frame will still be changed as applicable.
func (self *Button) SetOnOverMouseOnlyA(member bool) {
    self.Object.Set("onOverMouseOnly", member)
}

// JustReleasedPreventsOver Suppresse the over event if a pointer was just released and it matches the given {@link Phaser.PointerModer pointer mode bitmask}.
// 
// This behavior was introduced in Phaser 2.3.1; this property is a soft-revert of the change.
func (self *Button) JustReleasedPreventsOver() *PointerMode{
    return &PointerMode{self.Object.Get("justReleasedPreventsOver")}
}

// SetJustReleasedPreventsOverA Suppresse the over event if a pointer was just released and it matches the given {@link Phaser.PointerModer pointer mode bitmask}.
// 
// This behavior was introduced in Phaser 2.3.1; this property is a soft-revert of the change.
func (self *Button) SetJustReleasedPreventsOverA(member *PointerMode) {
    self.Object.Set("justReleasedPreventsOver", member)
}

// FreezeFrames When true the the texture frame will not be automatically switched on up/down/over/out events.
func (self *Button) FreezeFrames() bool{
    return self.Object.Get("freezeFrames").Bool()
}

// SetFreezeFramesA When true the the texture frame will not be automatically switched on up/down/over/out events.
func (self *Button) SetFreezeFramesA(member bool) {
    self.Object.Set("freezeFrames", member)
}

// ForceOut When the Button is touched / clicked and then released you can force it to enter a state of "out" instead of "up".
// 
// This can also accept a {@link Phaser.PointerModer pointer mode bitmask} for more refined control.
func (self *Button) ForceOut() interface{}{
    return self.Object.Get("forceOut")
}

// SetForceOutA When the Button is touched / clicked and then released you can force it to enter a state of "out" instead of "up".
// 
// This can also accept a {@link Phaser.PointerModer pointer mode bitmask} for more refined control.
func (self *Button) SetForceOutA(member interface{}) {
    self.Object.Set("forceOut", member)
}

// Anchor The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Button) Anchor() *Point{
    return &Point{self.Object.Get("anchor")}
}

// SetAnchorA The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Button) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// Texture The texture that the sprite is using
func (self *Button) Texture() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// SetTextureA The texture that the sprite is using
func (self *Button) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// Tint The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Button) Tint() int{
    return self.Object.Get("tint").Int()
}

// SetTintA The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Button) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// TintedTexture A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Button) TintedTexture() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// SetTintedTextureA A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Button) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// BlendMode The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Button) BlendMode() int{
    return self.Object.Get("blendMode").Int()
}

// SetBlendModeA The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Button) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// Shader The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Button) Shader() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// SetShaderA The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Button) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Exists Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Button) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Button) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// Width The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Button) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Button) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Button) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Button) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Children [read-only] The array of children of this container.
func (self *Button) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *Button) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Button) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Button) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// Game A reference to the currently running Game.
func (self *Button) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Button) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Button) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Button) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Data An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Button) Data() interface{}{
    return self.Object.Get("data")
}

// SetDataA An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Button) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// Components The components this Game Object has installed.
func (self *Button) Components() interface{}{
    return self.Object.Get("components")
}

// SetComponentsA The components this Game Object has installed.
func (self *Button) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// Z The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Button) Z() int{
    return self.Object.Get("z").Int()
}

// SetZA The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Button) SetZA(member int) {
    self.Object.Set("z", member)
}

// Events All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Button) Events() *Events{
    return &Events{self.Object.Get("events")}
}

// SetEventsA All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Button) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// Animations If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Button) Animations() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// SetAnimationsA If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Button) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// Key The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Button) Key() interface{}{
    return self.Object.Get("key")
}

// SetKeyA The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Button) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// World The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Button) World() *Point{
    return &Point{self.Object.Get("world")}
}

// SetWorldA The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Button) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// Debug A debug flag designed for use with `Game.enableStep`.
func (self *Button) Debug() bool{
    return self.Object.Get("debug").Bool()
}

// SetDebugA A debug flag designed for use with `Game.enableStep`.
func (self *Button) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// PreviousPosition The position the Game Object was located in the previous frame.
func (self *Button) PreviousPosition() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// SetPreviousPositionA The position the Game Object was located in the previous frame.
func (self *Button) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// PreviousRotation The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Button) PreviousRotation() int{
    return self.Object.Get("previousRotation").Int()
}

// SetPreviousRotationA The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Button) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// RenderOrderID The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Button) RenderOrderID() int{
    return self.Object.Get("renderOrderID").Int()
}

// SetRenderOrderIDA The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Button) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// Fresh A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Button) Fresh() bool{
    return self.Object.Get("fresh").Bool()
}

// SetFreshA A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Button) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// PendingDestroy A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Button) PendingDestroy() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Button) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// Angle The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *Button) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *Button) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// AutoCull A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Button) AutoCull() bool{
    return self.Object.Get("autoCull").Bool()
}

// SetAutoCullA A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Button) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// InCamera Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Button) InCamera() bool{
    return self.Object.Get("inCamera").Bool()
}

// SetInCameraA Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Button) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// OffsetX The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Button) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// SetOffsetXA The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Button) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// OffsetY The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Button) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// SetOffsetYA The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Button) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// CenterX The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Button) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// SetCenterXA The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Button) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// CenterY The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Button) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// SetCenterYA The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Button) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// Left The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Button) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Button) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Button) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Button) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Button) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Button) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Button) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Button) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// CropRect The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Button) CropRect() *Rectangle{
    return &Rectangle{self.Object.Get("cropRect")}
}

// SetCropRectA The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Button) SetCropRectA(member *Rectangle) {
    self.Object.Set("cropRect", member)
}

// DestroyPhase As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Button) DestroyPhase() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// SetDestroyPhaseA As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Button) SetDestroyPhaseA(member bool) {
    self.Object.Set("destroyPhase", member)
}

// FixedToCamera A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *Button) FixedToCamera() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// SetFixedToCameraA A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *Button) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// CameraOffset The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Button) CameraOffset() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Button) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// Input The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Button) Input() interface{}{
    return self.Object.Get("input")
}

// SetInputA The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Button) SetInputA(member interface{}) {
    self.Object.Set("input", member)
}

// InputEnabled By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *Button) InputEnabled() bool{
    return self.Object.Get("inputEnabled").Bool()
}

// SetInputEnabledA By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *Button) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}

// Alive A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Button) Alive() bool{
    return self.Object.Get("alive").Bool()
}

// SetAliveA A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Button) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// Lifespan The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Button) Lifespan() int{
    return self.Object.Get("lifespan").Int()
}

// SetLifespanA The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Button) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}

// Frame Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Button) Frame() int{
    return self.Object.Get("frame").Int()
}

// SetFrameA Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Button) SetFrameA(member int) {
    self.Object.Set("frame", member)
}

// FrameName Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Button) FrameName() string{
    return self.Object.Get("frameName").String()
}

// SetFrameNameA Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Button) SetFrameNameA(member string) {
    self.Object.Set("frameName", member)
}

// Smoothed Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Button) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Button) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}


// ClearFrames Clears all of the frames set on this Button.
func (self *Button) ClearFrames() {
    self.Object.Call("clearFrames")
}

// ClearFramesI Clears all of the frames set on this Button.
func (self *Button) ClearFramesI(args ...interface{}) {
    self.Object.Call("clearFrames", args)
}

// RemovedFromWorld Called when this Button is removed from the World.
func (self *Button) RemovedFromWorld() {
    self.Object.Call("removedFromWorld")
}

// RemovedFromWorldI Called when this Button is removed from the World.
func (self *Button) RemovedFromWorldI(args ...interface{}) {
    self.Object.Call("removedFromWorld", args)
}

// SetStateFrame Set the frame name/ID for the given state.
func (self *Button) SetStateFrame(state interface{}, frame interface{}, switchImmediately bool) {
    self.Object.Call("setStateFrame", state, frame, switchImmediately)
}

// SetStateFrameI Set the frame name/ID for the given state.
func (self *Button) SetStateFrameI(args ...interface{}) {
    self.Object.Call("setStateFrame", args)
}

// ChangeStateFrame Change the frame to that of the given state, _if_ the state has a frame assigned _and_ if the frames are not currently "frozen".
func (self *Button) ChangeStateFrame(state interface{}) bool{
    return self.Object.Call("changeStateFrame", state).Bool()
}

// ChangeStateFrameI Change the frame to that of the given state, _if_ the state has a frame assigned _and_ if the frames are not currently "frozen".
func (self *Button) ChangeStateFrameI(args ...interface{}) bool{
    return self.Object.Call("changeStateFrame", args).Bool()
}

// SetFrames Used to manually set the frames that will be used for the different states of the Button.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); these are the same values that can be used with a Sprite constructor.
func (self *Button) SetFrames() {
    self.Object.Call("setFrames")
}

// SetFrames1O Used to manually set the frames that will be used for the different states of the Button.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); these are the same values that can be used with a Sprite constructor.
func (self *Button) SetFrames1O(overFrame interface{}) {
    self.Object.Call("setFrames", overFrame)
}

// SetFrames2O Used to manually set the frames that will be used for the different states of the Button.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); these are the same values that can be used with a Sprite constructor.
func (self *Button) SetFrames2O(overFrame interface{}, outFrame interface{}) {
    self.Object.Call("setFrames", overFrame, outFrame)
}

// SetFrames3O Used to manually set the frames that will be used for the different states of the Button.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); these are the same values that can be used with a Sprite constructor.
func (self *Button) SetFrames3O(overFrame interface{}, outFrame interface{}, downFrame interface{}) {
    self.Object.Call("setFrames", overFrame, outFrame, downFrame)
}

// SetFrames4O Used to manually set the frames that will be used for the different states of the Button.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); these are the same values that can be used with a Sprite constructor.
func (self *Button) SetFrames4O(overFrame interface{}, outFrame interface{}, downFrame interface{}, upFrame interface{}) {
    self.Object.Call("setFrames", overFrame, outFrame, downFrame, upFrame)
}

// SetFramesI Used to manually set the frames that will be used for the different states of the Button.
// 
// Frames can be specified as either an integer (the frame ID) or a string (the frame name); these are the same values that can be used with a Sprite constructor.
func (self *Button) SetFramesI(args ...interface{}) {
    self.Object.Call("setFrames", args)
}

// SetStateSound Set the sound/marker for the given state.
func (self *Button) SetStateSound(state interface{}) {
    self.Object.Call("setStateSound", state)
}

// SetStateSound1O Set the sound/marker for the given state.
func (self *Button) SetStateSound1O(state interface{}, sound interface{}) {
    self.Object.Call("setStateSound", state, sound)
}

// SetStateSound2O Set the sound/marker for the given state.
func (self *Button) SetStateSound2O(state interface{}, sound interface{}, marker string) {
    self.Object.Call("setStateSound", state, sound, marker)
}

// SetStateSoundI Set the sound/marker for the given state.
func (self *Button) SetStateSoundI(args ...interface{}) {
    self.Object.Call("setStateSound", args)
}

// PlayStateSound Play the sound for the given state, _if_ the state has a sound assigned.
func (self *Button) PlayStateSound(state interface{}) bool{
    return self.Object.Call("playStateSound", state).Bool()
}

// PlayStateSoundI Play the sound for the given state, _if_ the state has a sound assigned.
func (self *Button) PlayStateSoundI(args ...interface{}) bool{
    return self.Object.Call("playStateSound", args).Bool()
}

// SetSounds Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds() {
    self.Object.Call("setSounds")
}

// SetSounds1O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds1O(overSound interface{}) {
    self.Object.Call("setSounds", overSound)
}

// SetSounds2O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds2O(overSound interface{}, overMarker string) {
    self.Object.Call("setSounds", overSound, overMarker)
}

// SetSounds3O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds3O(overSound interface{}, overMarker string, downSound interface{}) {
    self.Object.Call("setSounds", overSound, overMarker, downSound)
}

// SetSounds4O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds4O(overSound interface{}, overMarker string, downSound interface{}, downMarker string) {
    self.Object.Call("setSounds", overSound, overMarker, downSound, downMarker)
}

// SetSounds5O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds5O(overSound interface{}, overMarker string, downSound interface{}, downMarker string, outSound interface{}) {
    self.Object.Call("setSounds", overSound, overMarker, downSound, downMarker, outSound)
}

// SetSounds6O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds6O(overSound interface{}, overMarker string, downSound interface{}, downMarker string, outSound interface{}, outMarker string) {
    self.Object.Call("setSounds", overSound, overMarker, downSound, downMarker, outSound, outMarker)
}

// SetSounds7O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds7O(overSound interface{}, overMarker string, downSound interface{}, downMarker string, outSound interface{}, outMarker string, upSound interface{}) {
    self.Object.Call("setSounds", overSound, overMarker, downSound, downMarker, outSound, outMarker, upSound)
}

// SetSounds8O Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSounds8O(overSound interface{}, overMarker string, downSound interface{}, downMarker string, outSound interface{}, outMarker string, upSound interface{}, upMarker string) {
    self.Object.Call("setSounds", overSound, overMarker, downSound, downMarker, outSound, outMarker, upSound, upMarker)
}

// SetSoundsI Sets the sounds to be played whenever this Button is interacted with. Sounds can be either full Sound objects, or markers pointing to a section of a Sound object.
// The most common forms of sounds are 'hover' effects and 'click' effects, which is why the order of the parameters is overSound then downSound.
// 
// Call this function with no parameters to reset all sounds on this Button.
func (self *Button) SetSoundsI(args ...interface{}) {
    self.Object.Call("setSounds", args)
}

// SetOverSound The Sound to be played when a Pointer moves over this Button.
func (self *Button) SetOverSound(sound interface{}) {
    self.Object.Call("setOverSound", sound)
}

// SetOverSound1O The Sound to be played when a Pointer moves over this Button.
func (self *Button) SetOverSound1O(sound interface{}, marker string) {
    self.Object.Call("setOverSound", sound, marker)
}

// SetOverSoundI The Sound to be played when a Pointer moves over this Button.
func (self *Button) SetOverSoundI(args ...interface{}) {
    self.Object.Call("setOverSound", args)
}

// SetOutSound The Sound to be played when a Pointer moves out of this Button.
func (self *Button) SetOutSound(sound interface{}) {
    self.Object.Call("setOutSound", sound)
}

// SetOutSound1O The Sound to be played when a Pointer moves out of this Button.
func (self *Button) SetOutSound1O(sound interface{}, marker string) {
    self.Object.Call("setOutSound", sound, marker)
}

// SetOutSoundI The Sound to be played when a Pointer moves out of this Button.
func (self *Button) SetOutSoundI(args ...interface{}) {
    self.Object.Call("setOutSound", args)
}

// SetDownSound The Sound to be played when a Pointer presses down on this Button.
func (self *Button) SetDownSound(sound interface{}) {
    self.Object.Call("setDownSound", sound)
}

// SetDownSound1O The Sound to be played when a Pointer presses down on this Button.
func (self *Button) SetDownSound1O(sound interface{}, marker string) {
    self.Object.Call("setDownSound", sound, marker)
}

// SetDownSoundI The Sound to be played when a Pointer presses down on this Button.
func (self *Button) SetDownSoundI(args ...interface{}) {
    self.Object.Call("setDownSound", args)
}

// SetUpSound The Sound to be played when a Pointer has pressed down and is released from this Button.
func (self *Button) SetUpSound(sound interface{}) {
    self.Object.Call("setUpSound", sound)
}

// SetUpSound1O The Sound to be played when a Pointer has pressed down and is released from this Button.
func (self *Button) SetUpSound1O(sound interface{}, marker string) {
    self.Object.Call("setUpSound", sound, marker)
}

// SetUpSoundI The Sound to be played when a Pointer has pressed down and is released from this Button.
func (self *Button) SetUpSoundI(args ...interface{}) {
    self.Object.Call("setUpSound", args)
}

// OnInputOverHandler Internal function that handles input events.
func (self *Button) OnInputOverHandler(sprite *Button, pointer *Pointer) {
    self.Object.Call("onInputOverHandler", sprite, pointer)
}

// OnInputOverHandlerI Internal function that handles input events.
func (self *Button) OnInputOverHandlerI(args ...interface{}) {
    self.Object.Call("onInputOverHandler", args)
}

// OnInputOutHandler Internal function that handles input events.
func (self *Button) OnInputOutHandler(sprite *Button, pointer *Pointer) {
    self.Object.Call("onInputOutHandler", sprite, pointer)
}

// OnInputOutHandlerI Internal function that handles input events.
func (self *Button) OnInputOutHandlerI(args ...interface{}) {
    self.Object.Call("onInputOutHandler", args)
}

// OnInputDownHandler Internal function that handles input events.
func (self *Button) OnInputDownHandler(sprite *Button, pointer *Pointer) {
    self.Object.Call("onInputDownHandler", sprite, pointer)
}

// OnInputDownHandlerI Internal function that handles input events.
func (self *Button) OnInputDownHandlerI(args ...interface{}) {
    self.Object.Call("onInputDownHandler", args)
}

// OnInputUpHandler Internal function that handles input events.
func (self *Button) OnInputUpHandler(sprite *Button, pointer *Pointer) {
    self.Object.Call("onInputUpHandler", sprite, pointer)
}

// OnInputUpHandlerI Internal function that handles input events.
func (self *Button) OnInputUpHandlerI(args ...interface{}) {
    self.Object.Call("onInputUpHandler", args)
}

// PreUpdate Automatically called by World.preUpdate.
func (self *Button) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Automatically called by World.preUpdate.
func (self *Button) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// SetTexture Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Button) SetTexture(texture *Texture) {
    self.Object.Call("setTexture", texture)
}

// SetTexture1O Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Button) SetTexture1O(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// SetTextureI Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Button) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// OnTextureUpdate When the texture is updated, this event will fire to update the scale and frame
func (self *Button) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// OnTextureUpdateI When the texture is updated, this event will fire to update the scale and frame
func (self *Button) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// GetBounds Returns the bounds of the Sprite as a rectangle.
// The bounds calculation takes the worldTransform into account.
// 
// It is important to note that the transform is not updated when you call this method.
// So if this Sprite is the child of a Display Object which has had its transform
// updated since the last render pass, those changes will not yet have been applied
// to this Sprites worldTransform. If you need to ensure that all parent transforms
// are factored into this getBounds operation then you should call `updateTransform`
// on the root most object in this Sprites display list first.
func (self *Button) GetBounds(matrix *Matrix) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", matrix)}
}

// GetBoundsI Returns the bounds of the Sprite as a rectangle.
// The bounds calculation takes the worldTransform into account.
// 
// It is important to note that the transform is not updated when you call this method.
// So if this Sprite is the child of a Display Object which has had its transform
// updated since the last render pass, those changes will not yet have been applied
// to this Sprites worldTransform. If you need to ensure that all parent transforms
// are factored into this getBounds operation then you should call `updateTransform`
// on the root most object in this Sprites display list first.
func (self *Button) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *Button) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGL1O Renders the object using the WebGL renderer
func (self *Button) _renderWebGL1O(renderSession *RenderSession, matrix *Matrix) {
    self.Object.Call("_renderWebGL", renderSession, matrix)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *Button) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *Button) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvas1O Renders the object using the Canvas renderer
func (self *Button) _renderCanvas1O(renderSession *RenderSession, matrix *Matrix) {
    self.Object.Call("_renderCanvas", renderSession, matrix)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *Button) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// AddChild Adds a child to the container.
func (self *Button) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *Button) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Button) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Button) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *Button) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *Button) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *Button) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *Button) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *Button) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *Button) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *Button) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *Button) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *Button) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *Button) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *Button) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *Button) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *Button) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *Button) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Button) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Button) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// SetStageReference Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Button) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// SetStageReferenceI Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Button) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// RemoveStageReference Removes the current stage reference from the container and all of its children.
func (self *Button) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// RemoveStageReferenceI Removes the current stage reference from the container and all of its children.
func (self *Button) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Update Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Button) Update() {
    self.Object.Call("update")
}

// UpdateI Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Button) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PostUpdate Internal method called by the World postUpdate cycle.
func (self *Button) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Internal method called by the World postUpdate cycle.
func (self *Button) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Play Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Button) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Play1O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Button) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Play2O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Button) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Play3O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Button) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// PlayI Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Button) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}

// AlignIn Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignIn(container interface{}) interface{}{
    return self.Object.Call("alignIn", container)
}

// AlignIn1O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignIn1O(container interface{}, position int) interface{}{
    return self.Object.Call("alignIn", container, position)
}

// AlignIn2O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignIn2O(container interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX)
}

// AlignIn3O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX, offsetY)
}

// AlignInI Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignInI(args ...interface{}) interface{}{
    return self.Object.Call("alignIn", args)
}

// AlignTo Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignTo(parent interface{}) interface{}{
    return self.Object.Call("alignTo", parent)
}

// AlignTo1O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignTo1O(parent interface{}, position int) interface{}{
    return self.Object.Call("alignTo", parent, position)
}

// AlignTo2O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignTo2O(parent interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX)
}

// AlignTo3O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX, offsetY)
}

// AlignToI Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *Button) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// BringToTop Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// BringToTopI Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// SendToBack Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// SendToBackI Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// MoveUp Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// MoveUpI Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// MoveDown Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// MoveDownI Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Button) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown", args)}
}

// Crop Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *Button) Crop(rect *Rectangle) {
    self.Object.Call("crop", rect)
}

// Crop1O Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *Button) Crop1O(rect *Rectangle, copy bool) {
    self.Object.Call("crop", rect, copy)
}

// CropI Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *Button) CropI(args ...interface{}) {
    self.Object.Call("crop", args)
}

// UpdateCrop If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Button) UpdateCrop() {
    self.Object.Call("updateCrop")
}

// UpdateCropI If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Button) UpdateCropI(args ...interface{}) {
    self.Object.Call("updateCrop", args)
}

// Destroy Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Button) Destroy() {
    self.Object.Call("destroy")
}

// Destroy1O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Button) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroy2O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Button) Destroy2O(destroyChildren bool, destroyTexture bool) {
    self.Object.Call("destroy", destroyChildren, destroyTexture)
}

// DestroyI Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Button) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Revive Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Button) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Revive1O Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Button) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// ReviveI Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Button) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", args)}
}

// Kill Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Button) Kill() *DisplayObject{
    return &DisplayObject{self.Object.Call("kill")}
}

// KillI Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Button) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}

// LoadTexture Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Button) LoadTexture(key interface{}) {
    self.Object.Call("loadTexture", key)
}

// LoadTexture1O Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Button) LoadTexture1O(key interface{}, frame interface{}) {
    self.Object.Call("loadTexture", key, frame)
}

// LoadTexture2O Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Button) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
    self.Object.Call("loadTexture", key, frame, stopAnimation)
}

// LoadTextureI Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Button) LoadTextureI(args ...interface{}) {
    self.Object.Call("loadTexture", args)
}

// SetFrame Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Button) SetFrame(frame *Frame) {
    self.Object.Call("setFrame", frame)
}

// SetFrameI Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Button) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// ResizeFrame Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Button) ResizeFrame(parent interface{}, width int, height int) {
    self.Object.Call("resizeFrame", parent, width, height)
}

// ResizeFrameI Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Button) ResizeFrameI(args ...interface{}) {
    self.Object.Call("resizeFrame", args)
}

// ResetFrame Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Button) ResetFrame() {
    self.Object.Call("resetFrame")
}

// ResetFrameI Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Button) ResetFrameI(args ...interface{}) {
    self.Object.Call("resetFrame", args)
}

// Overlap Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Button) Overlap(displayObject interface{}) bool{
    return self.Object.Call("overlap", displayObject).Bool()
}

// OverlapI Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Button) OverlapI(args ...interface{}) bool{
    return self.Object.Call("overlap", args).Bool()
}

// Reset Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Button) Reset(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y)}
}

// Reset1O Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Button) Reset1O(x int, y int, health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y, health)}
}

// ResetI Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Button) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}

