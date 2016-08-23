package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

func GetGlobalConst(key string) int {
	return js.Global.Get("Phaser").Get(key).Int()
}

func GetKeyboardKey(key string) int {
	return js.Global.Get("Phaser").Get("Keyboard").Get(key).Int()
}
