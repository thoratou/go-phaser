// Automatic generation for Phaser.Color
// generated file Color.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Phaser.Color class is a set of static methods that assist in color manipulation and conversion.
type Color struct {
    *js.Object
}




// Packs the r, g, b, a components into a single integer, for use with Int32Array.
// If device is little endian then ABGR order is used. Otherwise RGBA order is used.
func (self *Color) PackPixelI(args ...interface{}) float64{
    return self.Call("packPixel", args).Float()
}

// Unpacks the r, g, b, a components into the specified color object, or a new
// object, for use with Int32Array. If little endian, then ABGR order is used when
// unpacking, otherwise, RGBA order is used. The resulting color object has the
// `r, g, b, a` properties which are unrelated to endianness.
// 
// Note that the integer is assumed to be packed in the correct endianness. On little-endian
// the format is 0xAABBGGRR and on big-endian the format is 0xRRGGBBAA. If you want a
// endian-independent method, use fromRGBA(rgba) and toRGBA(r, g, b, a).
func (self *Color) UnpackPixelI(args ...interface{}) interface{}{
    return self.Call("unpackPixel", args)
}

// A utility to convert an integer in 0xRRGGBBAA format to a color object.
// This does not rely on endianness.
func (self *Color) FromRGBAI(args ...interface{}) interface{}{
    return self.Call("fromRGBA", args)
}

// A utility to convert RGBA components to a 32 bit integer in RRGGBBAA format.
func (self *Color) ToRGBAI(args ...interface{}) float64{
    return self.Call("toRGBA", args).Float()
}

// Converts RGBA components to a 32 bit integer in AABBGGRR format.
func (self *Color) ToABGRI(args ...interface{}) float64{
    return self.Call("toABGR", args).Float()
}

// Converts an RGB color value to HSL (hue, saturation and lightness).
// Conversion forumla from http://en.wikipedia.org/wiki/HSL_color_space.
// Assumes RGB values are contained in the set [0, 255] and returns h, s and l in the set [0, 1].
// Based on code by Michael Jackson (https://github.com/mjijackson)
func (self *Color) RGBtoHSLI(args ...interface{}) interface{}{
    return self.Call("RGBtoHSL", args)
}

// Converts an HSL (hue, saturation and lightness) color value to RGB.
// Conversion forumla from http://en.wikipedia.org/wiki/HSL_color_space.
// Assumes HSL values are contained in the set [0, 1] and returns r, g and b values in the set [0, 255].
// Based on code by Michael Jackson (https://github.com/mjijackson)
func (self *Color) HSLtoRGBI(args ...interface{}) interface{}{
    return self.Call("HSLtoRGB", args)
}

// Converts an RGB color value to HSV (hue, saturation and value).
// Conversion forumla from http://en.wikipedia.org/wiki/HSL_color_space.
// Assumes RGB values are contained in the set [0, 255] and returns h, s and v in the set [0, 1].
// Based on code by Michael Jackson (https://github.com/mjijackson)
func (self *Color) RGBtoHSVI(args ...interface{}) interface{}{
    return self.Call("RGBtoHSV", args)
}

// Converts an HSV (hue, saturation and value) color value to RGB.
// Conversion forumla from http://en.wikipedia.org/wiki/HSL_color_space.
// Assumes HSV values are contained in the set [0, 1] and returns r, g and b values in the set [0, 255].
// Based on code by Michael Jackson (https://github.com/mjijackson)
func (self *Color) HSVtoRGBI(args ...interface{}) interface{}{
    return self.Call("HSVtoRGB", args)
}

// Converts a hue to an RGB color.
// Based on code by Michael Jackson (https://github.com/mjijackson)
func (self *Color) HueToColorI(args ...interface{}) float64{
    return self.Call("hueToColor", args).Float()
}

// A utility function to create a lightweight 'color' object with the default components.
// Any components that are not specified will default to zero.
// 
// This is useful when you want to use a shared color object for the getPixel and getPixelAt methods.
func (self *Color) CreateColorI(args ...interface{}) interface{}{
    return self.Call("createColor", args)
}

// Takes a color object and updates the rgba, color and color32 properties.
func (self *Color) UpdateColorI(args ...interface{}) float64{
    return self.Call("updateColor", args).Float()
}

// Given an alpha and 3 color values this will return an integer representation of it.
func (self *Color) GetColor32I(args ...interface{}) float64{
    return self.Call("getColor32", args).Float()
}

// Given 3 color values this will return an integer representation of it.
func (self *Color) GetColorI(args ...interface{}) float64{
    return self.Call("getColor", args).Float()
}

// Converts the given color values into a string.
// If prefix was '#' it will be in the format `#RRGGBB` otherwise `0xAARRGGBB`.
func (self *Color) RGBtoStringI(args ...interface{}) string{
    return self.Call("RGBtoString", args).String()
}

// Converts a hex string into an integer color value.
func (self *Color) HexToRGBI(args ...interface{}) float64{
    return self.Call("hexToRGB", args).Float()
}

// Converts a hex string into a Phaser Color object.
// 
// The hex string can supplied as `'#0033ff'` or the short-hand format of `'#03f'`; it can begin with an optional "#" or "0x", or be unprefixed.    
// 
// An alpha channel is _not_ supported.
func (self *Color) HexToColorI(args ...interface{}) interface{}{
    return self.Call("hexToColor", args)
}

// Converts a CSS 'web' string into a Phaser Color object.
// 
// The web string can be in the format `'rgb(r,g,b)'` or `'rgba(r,g,b,a)'` where r/g/b are in the range [0..255] and a is in the range [0..1].
func (self *Color) WebToColorI(args ...interface{}) interface{}{
    return self.Call("webToColor", args)
}

// Converts a value - a "hex" string, a "CSS 'web' string", or a number - into red, green, blue, and alpha components.
// 
// The value can be a string (see `hexToColor` and `webToColor` for the supported formats) or a packed integer (see `getRGB`).
// 
// An alpha channel is _not_ supported when specifying a hex string.
func (self *Color) ValueToColorI(args ...interface{}) interface{}{
    return self.Call("valueToColor", args)
}

// Return a string containing a hex representation of the given color component.
func (self *Color) ComponentToHexI(args ...interface{}) string{
    return self.Call("componentToHex", args).String()
}

// Get HSV color wheel values in an array which will be 360 elements in size.
func (self *Color) HSVColorWheelI(args ...interface{}) []interface{}{
	array := self.Call("HSVColorWheel", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Get HSL color wheel values in an array which will be 360 elements in size.
func (self *Color) HSLColorWheelI(args ...interface{}) []interface{}{
	array := self.Call("HSLColorWheel", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Interpolates the two given colours based on the supplied step and currentStep properties.
func (self *Color) InterpolateColorI(args ...interface{}) float64{
    return self.Call("interpolateColor", args).Float()
}

// Interpolates the two given colours based on the supplied step and currentStep properties.
func (self *Color) InterpolateColorWithRGBI(args ...interface{}) float64{
    return self.Call("interpolateColorWithRGB", args).Float()
}

// Interpolates the two given colours based on the supplied step and currentStep properties.
func (self *Color) InterpolateRGBI(args ...interface{}) float64{
    return self.Call("interpolateRGB", args).Float()
}

// Returns a random color value between black and white
// Set the min value to start each channel from the given offset.
// Set the max value to restrict the maximum color used per channel.
func (self *Color) GetRandomColorI(args ...interface{}) float64{
    return self.Call("getRandomColor", args).Float()
}

// Return the component parts of a color as an Object with the properties alpha, red, green, blue.
// 
// Alpha will only be set if it exist in the given color (0xAARRGGBB)
func (self *Color) GetRGBI(args ...interface{}) interface{}{
    return self.Call("getRGB", args)
}

// Returns a CSS friendly string value from the given color.
func (self *Color) GetWebRGBI(args ...interface{}) string{
    return self.Call("getWebRGB", args).String()
}

// Given a native color value (in the format 0xAARRGGBB) this will return the Alpha component, as a value between 0 and 255.
func (self *Color) GetAlphaI(args ...interface{}) float64{
    return self.Call("getAlpha", args).Float()
}

// Given a native color value (in the format 0xAARRGGBB) this will return the Alpha component as a value between 0 and 1.
func (self *Color) GetAlphaFloatI(args ...interface{}) float64{
    return self.Call("getAlphaFloat", args).Float()
}

// Given a native color value (in the format 0xAARRGGBB) this will return the Red component, as a value between 0 and 255.
func (self *Color) GetRedI(args ...interface{}) float64{
    return self.Call("getRed", args).Float()
}

// Given a native color value (in the format 0xAARRGGBB) this will return the Green component, as a value between 0 and 255.
func (self *Color) GetGreenI(args ...interface{}) float64{
    return self.Call("getGreen", args).Float()
}

// Given a native color value (in the format 0xAARRGGBB) this will return the Blue component, as a value between 0 and 255.
func (self *Color) GetBlueI(args ...interface{}) float64{
    return self.Call("getBlue", args).Float()
}

// Blends the source color, ignoring the backdrop.
func (self *Color) BlendNormalI(args ...interface{}) int{
    return self.Call("blendNormal", args).Int()
}

// Selects the lighter of the backdrop and source colors.
func (self *Color) BlendLightenI(args ...interface{}) int{
    return self.Call("blendLighten", args).Int()
}

// Selects the darker of the backdrop and source colors.
func (self *Color) BlendDarkenI(args ...interface{}) int{
    return self.Call("blendDarken", args).Int()
}

// Multiplies the backdrop and source color values.
// The result color is always at least as dark as either of the two constituent
// colors. Multiplying any color with black produces black;
// multiplying with white leaves the original color unchanged.
func (self *Color) BlendMultiplyI(args ...interface{}) int{
    return self.Call("blendMultiply", args).Int()
}

// Takes the average of the source and backdrop colors.
func (self *Color) BlendAverageI(args ...interface{}) int{
    return self.Call("blendAverage", args).Int()
}

// Adds the source and backdrop colors together and returns the value, up to a maximum of 255.
func (self *Color) BlendAddI(args ...interface{}) int{
    return self.Call("blendAdd", args).Int()
}

// Combines the source and backdrop colors and returns their value minus 255.
func (self *Color) BlendSubtractI(args ...interface{}) int{
    return self.Call("blendSubtract", args).Int()
}

// Subtracts the darker of the two constituent colors from the lighter.
// 
// Painting with white inverts the backdrop color; painting with black produces no change.
func (self *Color) BlendDifferenceI(args ...interface{}) int{
    return self.Call("blendDifference", args).Int()
}

// Negation blend mode.
func (self *Color) BlendNegationI(args ...interface{}) int{
    return self.Call("blendNegation", args).Int()
}

// Multiplies the complements of the backdrop and source color values, then complements the result.
// The result color is always at least as light as either of the two constituent colors. 
// Screening any color with white produces white; screening with black leaves the original color unchanged.
func (self *Color) BlendScreenI(args ...interface{}) int{
    return self.Call("blendScreen", args).Int()
}

// Produces an effect similar to that of the Difference mode, but lower in contrast. 
// Painting with white inverts the backdrop color; painting with black produces no change.
func (self *Color) BlendExclusionI(args ...interface{}) int{
    return self.Call("blendExclusion", args).Int()
}

// Multiplies or screens the colors, depending on the backdrop color.
// Source colors overlay the backdrop while preserving its highlights and shadows. 
// The backdrop color is not replaced, but is mixed with the source color to reflect the lightness or darkness of the backdrop.
func (self *Color) BlendOverlayI(args ...interface{}) int{
    return self.Call("blendOverlay", args).Int()
}

// Darkens or lightens the colors, depending on the source color value. 
// 
// If the source color is lighter than 0.5, the backdrop is lightened, as if it were dodged; 
// this is useful for adding highlights to a scene. 
// 
// If the source color is darker than 0.5, the backdrop is darkened, as if it were burned in. 
// The degree of lightening or darkening is proportional to the difference between the source color and 0.5; 
// if it is equal to 0.5, the backdrop is unchanged.
// 
// Painting with pure black or white produces a distinctly darker or lighter area, but does not result in pure black or white. 
// The effect is similar to shining a diffused spotlight on the backdrop.
func (self *Color) BlendSoftLightI(args ...interface{}) int{
    return self.Call("blendSoftLight", args).Int()
}

// Multiplies or screens the colors, depending on the source color value. 
// 
// If the source color is lighter than 0.5, the backdrop is lightened, as if it were screened; 
// this is useful for adding highlights to a scene. 
// 
// If the source color is darker than 0.5, the backdrop is darkened, as if it were multiplied; 
// this is useful for adding shadows to a scene. 
// 
// The degree of lightening or darkening is proportional to the difference between the source color and 0.5; 
// if it is equal to 0.5, the backdrop is unchanged.
// 
// Painting with pure black or white produces pure black or white. The effect is similar to shining a harsh spotlight on the backdrop.
func (self *Color) BlendHardLightI(args ...interface{}) int{
    return self.Call("blendHardLight", args).Int()
}

// Brightens the backdrop color to reflect the source color. 
// Painting with black produces no change.
func (self *Color) BlendColorDodgeI(args ...interface{}) int{
    return self.Call("blendColorDodge", args).Int()
}

// Darkens the backdrop color to reflect the source color.
// Painting with white produces no change.
func (self *Color) BlendColorBurnI(args ...interface{}) int{
    return self.Call("blendColorBurn", args).Int()
}

// An alias for blendAdd, it simply sums the values of the two colors.
func (self *Color) BlendLinearDodgeI(args ...interface{}) int{
    return self.Call("blendLinearDodge", args).Int()
}

// An alias for blendSubtract, it simply sums the values of the two colors and subtracts 255.
func (self *Color) BlendLinearBurnI(args ...interface{}) int{
    return self.Call("blendLinearBurn", args).Int()
}

// This blend mode combines Linear Dodge and Linear Burn (rescaled so that neutral colors become middle gray).
// Dodge applies to values of top layer lighter than middle gray, and burn to darker values.
// The calculation simplifies to the sum of bottom layer and twice the top layer, subtract 128. The contrast decreases.
func (self *Color) BlendLinearLightI(args ...interface{}) int{
    return self.Call("blendLinearLight", args).Int()
}

// This blend mode combines Color Dodge and Color Burn (rescaled so that neutral colors become middle gray).
// Dodge applies when values in the top layer are lighter than middle gray, and burn to darker values.
// The middle gray is the neutral color. When color is lighter than this, this effectively moves the white point of the bottom 
// layer down by twice the difference; when it is darker, the black point is moved up by twice the difference. The perceived contrast increases.
func (self *Color) BlendVividLightI(args ...interface{}) int{
    return self.Call("blendVividLight", args).Int()
}

// If the backdrop color (light source) is lighter than 50%, the blendDarken mode is used, and colors lighter than the backdrop color do not change.
// If the backdrop color is darker than 50% gray, colors lighter than the blend color are replaced, and colors darker than the blend color do not change.
func (self *Color) BlendPinLightI(args ...interface{}) int{
    return self.Call("blendPinLight", args).Int()
}

// Runs blendVividLight on the source and backdrop colors.
// If the resulting color is 128 or more, it receives a value of 255; if less than 128, a value of 0.
// Therefore, all blended pixels have red, green, and blue channel values of either 0 or 255.
// This changes all pixels to primary additive colors (red, green, or blue), white, or black.
func (self *Color) BlendHardMixI(args ...interface{}) int{
    return self.Call("blendHardMix", args).Int()
}

// Reflect blend mode. This mode is useful when adding shining objects or light zones to images.
func (self *Color) BlendReflectI(args ...interface{}) int{
    return self.Call("blendReflect", args).Int()
}

// Glow blend mode. This mode is a variation of reflect mode with the source and backdrop colors swapped.
func (self *Color) BlendGlowI(args ...interface{}) int{
    return self.Call("blendGlow", args).Int()
}

// Phoenix blend mode. This subtracts the lighter color from the darker color, and adds 255, giving a bright result.
func (self *Color) BlendPhoenixI(args ...interface{}) int{
    return self.Call("blendPhoenix", args).Int()
}
